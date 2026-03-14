/// <reference types="node" />
import { defineConfig } from "@playwright/test";

export default defineConfig({
  webServer: {
    command: "vp build && vp preview -- --strictPort --port 4173",
    url: "http://127.0.0.1:4173",
    reuseExistingServer: !process.env.CI,
  },
  use: {
    baseURL: "http://127.0.0.1:4173",
  },
  testDir: "e2e",
});
