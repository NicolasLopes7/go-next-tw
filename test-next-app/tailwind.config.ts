import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx,json}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx,json}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx,json}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
      },
    },
  },
  plugins: [],
} satisfies Config;