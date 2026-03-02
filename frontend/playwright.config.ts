/// <reference types="node" />
import { defineConfig } from '@playwright/test';

export default defineConfig({
	webServer: {
		command: 'bun run build && bun run preview -- --strictPort --port 4173',
		url: 'http://127.0.0.1:4173',
		reuseExistingServer: !process.env.CI
	},
	testDir: 'e2e'
});
