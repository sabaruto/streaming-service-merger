/// <reference types="vitest/config" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";
import eslintPlugin from "vite-plugin-eslint";

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        react(),
        tsconfigPaths(),
        eslintPlugin({
            cache: false,
            include: ["./src/**/*.tsx"],
            exclude: [],
        }),
    ],
    test: {
        environment: "jsdom",
        include: ["./src/**/*.test.tsx"],
        alias: {
            "@/": new URL("./src/", import.meta.url).pathname,
        },
        browser: {
            provider: "playwright",
            enabled: true,
            name: "chromium",
            headless: true,
        },
    },
    server: {
        open: true,
        hmr: true,
    },
});
