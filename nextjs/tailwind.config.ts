import type { Config } from "tailwindcss";

const config: Config = {
    content: [
        "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
        "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    ],
    theme: {
        extend: {
            backgroundImage: {
                "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
                "gradient-conic":
                    "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
            },
            colors: {
                primary: 'rgba(25, 30, 36, 1)',
                secondary: '#2A323C'
            },
            textColor: {
                default: '#a6adbb'
            },
            gridTemplateColumns: {
                "auto-fit-cards": "repeat(auto-fit, minmax(277px, 1fr))",
            }
        },

    },
    plugins: [],
};
export default config;
