import { NextResponse } from "next/server";
import { Client } from "pg";
import { createClient } from "@/lib/supabase/server";
import { ingestBuffer } from "@/scripts/ingest/ingest";

export const runtime = "nodejs";
export const maxDuration = 60;

/**
 * Upload + ingest a health data file. Auth via Supabase session; the file is
 * processed server-side and written with the user's owner_id. Uses a direct
 * Postgres connection (DATABASE_URL) for the normalizer's transactional work.
 */
export async function POST(request: Request) {
  const supabase = await createClient();
  const {
    data: { user },
  } = await supabase.auth.getUser();
  if (!user) return NextResponse.json({ error: "unauthorized" }, { status: 401 });

  const form = await request.formData();
  const file = form.get("file");
  if (!(file instanceof File)) {
    return NextResponse.json({ error: "no file" }, { status: 400 });
  }
  const buf = Buffer.from(await file.arrayBuffer());

  const db = new Client({ connectionString: process.env.DATABASE_URL });
  await db.connect();
  try {
    const result = await ingestBuffer(db, user.id, file.name, buf);
    return NextResponse.json(result);
  } catch (e) {
    return NextResponse.json({ error: (e as Error).message }, { status: 500 });
  } finally {
    await db.end();
  }
}
