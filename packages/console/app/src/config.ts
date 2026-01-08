/**
 * Application-wide constants and configuration
 */
export const config = {
  // Base URL
  baseUrl: "https://strukcode.ai",

  // GitHub
  github: {
    repoUrl: "https://github.com/harutk7/strukcode",
    starsFormatted: {
      compact: "50K",
      full: "50,000",
    },
  },

  // Social links
  social: {
    twitter: "https://x.com/strukcode",
    discord: "https://discord.gg/strukcode",
  },

  // Static stats (used on landing page)
  stats: {
    contributors: "500",
    commits: "6,500",
    monthlyUsers: "650,000",
  },
} as const
