/// <reference types="node" />
import { defineConfig } from "@playwright/test";

const backendURL = "http://127.0.0.1:8081";
const frontendURL = "http://localhost:4173";

export default defineConfig({
  testDir: "e2e",
  testMatch: "**/*.fullstack.spec.ts",
  use: {
    baseURL: frontendURL,
  },
  webServer: [
    {
      command: "go run ./backend/cmd/server",
      cwd: "..",
      env: {
        ...process.env,
        BARNLOG_AUTO_MIGRATE: "true",
        BARNLOG_DB_PATH: "/tmp/barnlog-e2e.sqlite3",
        BARNLOG_ENV: "test",
        BARNLOG_FILE_DIR: "/tmp/barnlog-e2e-files",
        BARNLOG_HTTP_ADDR: "127.0.0.1:8081",
        BARNLOG_LOG_LEVEL: "error",
      },
      name: "Backend",
      reuseExistingServer: !process.env.CI,
      timeout: 120_000,
      url: `${backendURL}/healthz`,
    },
    {
      command: "bun run build && bun run preview -- --host localhost --strictPort --port 4173",
      cwd: ".",
      env: {
        ...process.env,
        PUBLIC_API_BASE_URL: backendURL,
      },
      name: "Frontend",
      reuseExistingServer: !process.env.CI,
      timeout: 120_000,
      url: frontendURL,
    },
  ],
});
