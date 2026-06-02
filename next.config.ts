import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // node-only libs used in route handlers / scripts — keep out of the bundle
  serverExternalPackages: ["pdf-parse", "pg"],
  eslint: {
    // Lint runs separately in CI; don't block production builds on it.
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
