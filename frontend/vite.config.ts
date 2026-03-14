import { sveltekit } from "@sveltejs/kit/vite";
import { codecovSvelteKitPlugin } from "@codecov/sveltekit-plugin";
import tailwindcss from "@tailwindcss/vite";
import { playwright } from "vite-plus/test/browser-playwright";
import devtoolsJson from "vite-plugin-devtools-json";
import { defineConfig } from "vite-plus";

export default defineConfig({
  plugins: [
    tailwindcss(),
    sveltekit(),
    devtoolsJson(),
    // Put the Codecov SvelteKit plugin after all other plugins.
    codecovSvelteKitPlugin({
      enableBundleAnalysis: true,
      bundleName: "barnlog-frontend",
      uploadToken: process.env.CODECOV_TOKEN,
    }),
  ],
  fmt: {
    ignorePatterns: [
      "node_modules/**",
      ".svelte-kit/**",
      "build/**",
      "dist/**",
      "coverage/**",
      "src/lib/api/generated/**",
    ],
  },
  lint: {
    categories: {
      suspicious: "deny",
      perf: "deny",
      style: "deny",
      pedantic: "deny",
    },
    rules: {},
    env: {
      builtin: true,
    },
    globals: {},
    ignorePatterns: [
      "node_modules/**",
      ".svelte-kit/**",
      "build/**",
      "dist/**",
      "coverage/**",
      "src/lib/api/generated/**",
      "src/lib/index.ts",
    ],
    overrides: [
      {
        files: ["src/app.d.ts"],
        rules: {
          "eslint/capitalized-comments": "off",
          "unicorn/require-module-specifiers": "off",
        },
      },
      {
        files: ["src/**/*.svelte"],
        rules: {
          "eslint/sort-imports": "off",
        },
      },
      {
        files: ["src/lib/ui/components/*.svelte"],
        rules: {
          "oxc/no-rest-spread-properties": "off",
        },
      },
      {
        files: ["src/**/*.spec.ts", "src/**/*.test.ts", "e2e/**/*.ts"],
        rules: {
          "eslint/no-magic-numbers": "off",
          "eslint/sort-imports": "off",
          "oxc/no-async-await": "off",
        },
      },
      {
        files: [
          "*.config.{js,ts,mjs,cjs}",
          "eslint.config.js",
          "vite.config.ts",
          "playwright.config.ts",
        ],
        rules: {
          "eslint/sort-imports": "off",
          "eslint/sort-keys": "off",
          "oxc/no-rest-spread-properties": "off",
        },
      },
    ],
  },
  test: {
    expect: { requireAssertions: true },
    coverage: {
      provider: "v8",
      reporter: ["text", "lcov"],
      reportsDirectory: "coverage",
      thresholds: {
        lines: 70,
        functions: 70,
        statements: 70,
        branches: 60,
      },
    },
    projects: [
      {
        extends: true,
        test: {
          name: "client",
          browser: {
            enabled: true,
            provider: playwright(),
            instances: [{ browser: "chromium", headless: true }],
          },
          include: ["src/**/*.svelte.{test,spec}.{js,ts}"],
          exclude: ["src/lib/server/**"],
        },
      },

      {
        extends: true,
        test: {
          name: "unit",
          environment: "node",
          include: ["src/**/*.{test,spec}.{js,ts}"],
          exclude: ["src/**/*.svelte.{test,spec}.{js,ts}"],
        },
      },
    ],
  },
});
