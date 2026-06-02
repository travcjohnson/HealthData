import { redirect } from "next/navigation";
import { createClient } from "@/lib/supabase/server";
import { AppNav } from "@/components/AppNav";
import { Disclaimer } from "@/components/Disclaimer";

export default async function AppLayout({ children }: { children: React.ReactNode }) {
  const supabase = await createClient();
  const {
    data: { user },
  } = await supabase.auth.getUser();
  if (!user) redirect("/login");

  return (
    <div className="flex min-h-screen">
      <aside className="hidden w-60 flex-shrink-0 border-r border-slate-200 bg-white p-4 md:block">
        <AppNav name={user.email ?? "Signed in"} />
        <form action="/auth/signout" method="post" className="mt-4 px-2">
          <button className="text-xs text-slate-400 hover:text-slate-600">Sign out</button>
        </form>
      </aside>
      <div className="flex min-w-0 flex-1 flex-col">
        <main className="flex-1 p-6">{children}</main>
        <Disclaimer />
      </div>
    </div>
  );
}
