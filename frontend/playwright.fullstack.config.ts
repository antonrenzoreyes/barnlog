/// <reference types="node" />
import { defineConfig } from "@playwright/test";

const backendURL = "http://127.0.0.1:8081";
const frontendURL = "http://localhost:4173";
const randomBase = 16;
const randomSliceStart = 2;
const runID =
  process.env.GITHUB_RUN_ID ??
  `${Date.now()}-${Math.random().toString(randomBase).slice(randomSliceStart)}`;
const tempRoot = process.env.RUNNER_TEMP ?? "/tmp";
const dbPath = `${tempRoot}/barnlog-e2e-${runID}.sqlite3`;
const fileDir = `${tempRoot}/barnlog-e2e-files-${runID}`;

export default defineConfig({
  testDir: "e2e",
  testMatch: "**/*.fullstack.spec.ts",
  use: {
    baseURL: frontendURL,
  },
  webServer: [
    {
      command: `rm -rf "${dbPath}" "${fileDir}" && go run ./backend/cmd/server`,
      cwd: "..",
      env: {
        ...process.env,
        BARNLOG_AUTO_MIGRATE: "true",
        BARNLOG_DB_PATH: dbPath,
        BARNLOG_ENV: "test",
        BARNLOG_FILE_DIR: fileDir,
        BARNLOG_HTTP_ADDR: "127.0.0.1:8081",
        BARNLOG_LOG_LEVEL: "error",
      },
      name: "Backend",
      reuseExistingServer: false,
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
      reuseExistingServer: false,
      timeout: 120_000,
      url: frontendURL,
    },
  ],
});
