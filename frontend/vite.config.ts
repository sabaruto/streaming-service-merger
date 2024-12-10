/// <reference types="vitest/config" />
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vite.dev/config/
export default defineConfig({
    plugins: [react(), tsconfigPaths()],
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
});
