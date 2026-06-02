import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // pdf-parse and other node-only libs are used in route handlers / scripts
  serverExternalPackages: ["pdf-parse"],
  eslint: {
    // Lint runs separately in CI; don't block production builds on it.
    ignoreDuringBuilds: true,
  },
};

export default nextConfig;
