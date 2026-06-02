"use client";

import { useActionState } from "react";
import { signIn, signUp, type AuthState } from "./actions";
import { Activity } from "lucide-react";

const initial: AuthState = {};

export default function LoginPage() {
  const [signInState, signInAction, signingIn] = useActionState(signIn, initial);
  const [signUpState, signUpAction, signingUp] = useActionState(signUp, initial);
  const state = { ...signUpState, ...signInState };

  return (
    <main className="flex min-h-screen items-center justify-center bg-slate-50 p-4">
      <div className="w-full max-w-sm">
        <div className="mb-6 flex items-center gap-2">
          <Activity className="h-7 w-7 text-emerald-600" />
          <h1 className="text-xl font-semibold">Health Command Center</h1>
        </div>
        <form className="card space-y-4 p-6">
          <div>
            <label className="mb-1 block text-sm font-medium text-slate-700">Email</label>
            <input
              name="email"
              type="email"
              required
              autoComplete="email"
              className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:border-emerald-500 focus:outline-none"
            />
          </div>
          <div>
            <label className="mb-1 block text-sm font-medium text-slate-700">Password</label>
            <input
              name="password"
              type="password"
              required
              autoComplete="current-password"
              minLength={6}
              className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:border-emerald-500 focus:outline-none"
            />
          </div>

          {state.error && <p className="text-sm text-red-600">{state.error}</p>}
          {state.message && <p className="text-sm text-emerald-700">{state.message}</p>}

          <div className="flex gap-2">
            <button
              formAction={signInAction}
              disabled={signingIn}
              className="flex-1 rounded-lg bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-700 disabled:opacity-50"
            >
              {signingIn ? "Signing in…" : "Sign in"}
            </button>
            <button
              formAction={signUpAction}
              disabled={signingUp}
              className="flex-1 rounded-lg border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-50 disabled:opacity-50"
            >
              {signingUp ? "Creating…" : "Sign up"}
            </button>
          </div>
        </form>
        <p className="mt-4 text-center text-xs text-slate-400">
          Informational tool only — not medical advice.
        </p>
      </div>
    </main>
  );
}
