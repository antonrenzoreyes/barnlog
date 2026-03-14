import { sveltekit } from "@sveltejs/kit/vite";
import { codecovSvelteKitPlugin } from "@codecov/sveltekit-plugin";
import tailwindcss from "@tailwindcss/vite";
import { playwright } from "vite-plus/test/browser-playwright";
import devtoolsJson from "vite-plugin-devtools-json";
import { defineConfig } from "vite-plus";

export default defineConfig({
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
      restriction: "deny",
    },
    rules: {
      "svelte/comment-directive": "error",
      "svelte/infinite-reactive-loop": "error",
      "svelte/no-at-debug-tags": "warn",
      "svelte/no-at-html-tags": "error",
      "svelte/no-dom-manipulating": "error",
      "svelte/no-dupe-else-if-blocks": "error",
      "svelte/no-dupe-on-directives": "error",
      "svelte/no-dupe-style-properties": "error",
      "svelte/no-dupe-use-directives": "error",
      "svelte/no-export-load-in-svelte-module-in-kit-pages": "error",
      "svelte/no-immutable-reactive-statements": "error",
      "svelte/no-inner-declarations": "error",
      "svelte/no-inspect": "warn",
      "svelte/no-navigation-without-resolve": "error",
      "svelte/no-not-function-handler": "error",
      "svelte/no-object-in-text-mustaches": "error",
      "svelte/no-raw-special-elements": "error",
      "svelte/no-reactive-functions": "error",
      "svelte/no-reactive-literals": "error",
      "svelte/no-reactive-reassign": "error",
      "svelte/no-shorthand-style-property-overrides": "error",
      "svelte/no-store-async": "error",
      "svelte/no-svelte-internal": "error",
      "svelte/no-unknown-style-directive-property": "error",
      "svelte/no-unnecessary-state-wrap": "error",
      "svelte/no-unused-props": "error",
      "svelte/no-unused-svelte-ignore": "error",
      "svelte/no-useless-children-snippet": "error",
      "svelte/no-useless-mustaches": "error",
      "svelte/prefer-svelte-reactivity": "error",
      "svelte/prefer-writable-derived": "error",
      "svelte/require-each-key": "error",
      "svelte/require-event-dispatcher-types": "error",
      "svelte/require-store-reactive-access": "error",
      "svelte/system": "error",
      "svelte/valid-each-key": "error",
      "svelte/valid-prop-names-in-kit-pages": "error",
    },
    env: {
      builtin: true,
    },
    ignorePatterns: [
      "node_modules/**",
      ".svelte-kit/**",
      "build/**",
      "dist/**",
      "coverage/**",
      "src/lib/api/generated/**",
      "src/lib/index.ts",
      "**/node_modules",
      "**/.output",
      "**/.vercel",
      "**/.netlify",
      "**/.wrangler",
      ".svelte-kit",
      "build",
      "**/.DS_Store",
      "**/Thumbs.db",
      "**/.env",
      "**/.env.*",
      "!**/.env.example",
      "!**/.env.test.example",
      "**/vite.config.js.timestamp-*",
      "**/vite.config.ts.timestamp-*",
      "**/test-results",
      "**/playwright-report/",
      "**/coverage/",
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
      {
        files: ["*.svelte", "**/*.svelte"],
        rules: {
          "no-inner-declarations": "off",
          "no-self-assign": "off",
        },
        jsPlugins: ["eslint-plugin-svelte"],
      },
      {
        files: ["**/*.svelte", "**/*.svelte.ts", "**/*.svelte.js"],
        rules: {
          "svelte/no-inline-styles": "error",
          "svelte/no-at-debug-tags": "error",
          "svelte/button-has-type": "error",
          "svelte/no-ignored-unsubscribe": "error",
          "svelte/no-add-event-listener": "error",
          "svelte/no-target-blank": "error",
          "svelte/no-navigation-without-base": "error",
          "svelte/no-goto-without-base": "error",
          "svelte/valid-compile": "error",
          "svelte/valid-style-parse": "error",
          "svelte/block-lang": [
            "error",
            {
              script: "ts",
            },
          ],
        },
        jsPlugins: ["eslint-plugin-svelte"],
        env: {
          browser: true,
          node: true,
        },
      },
    ],
    jsPlugins: ["eslint-plugin-svelte"],
    options: {
      typeAware: true,
      typeCheck: true,
    },
  },
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
