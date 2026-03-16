import path from "node:path";
import { includeIgnoreFile } from "@eslint/compat";
import css from "@eslint/css";
import { defineConfig } from "eslint/config";
import eslintPluginBetterTailwindcss from "eslint-plugin-better-tailwindcss";
import svelte from "eslint-plugin-svelte";
import globals from "globals";
import { tailwind4 } from "tailwind-csstree";
import ts from "typescript-eslint";
import svelteConfig from "./svelte.config.js";

const gitignorePath = path.resolve(import.meta.dirname, ".gitignore");
const svelteRecommended = [];
for (const config of svelte.configs.recommended) {
  if (config.name === "svelte:recommended:rules") {
    svelteRecommended.push({
      ...config,
      files: [
        "*.svelte",
        "**/*.svelte",
        "*.svelte.js",
        "*.svelte.ts",
        "**/*.svelte.js",
        "**/*.svelte.ts",
      ],
    });
  } else {
    svelteRecommended.push(config);
  }
}
const betterTailwindSettings = {
  "better-tailwindcss": {
    entryPoint: "src/routes/layout.css",
  },
};

export default defineConfig(
  includeIgnoreFile(gitignorePath),
  ...svelteRecommended,
  {
    files: ["**/*.svelte", "**/*.svelte.ts", "**/*.svelte.js"],
    extends: [eslintPluginBetterTailwindcss.configs.recommended],
    settings: betterTailwindSettings,
    languageOptions: {
      globals: { ...globals.browser, ...globals.node },
      parserOptions: {
        projectService: true,
        extraFileExtensions: [".svelte"],
        parser: ts.parser,
        svelteConfig,
      },
    },
    rules: {
      "svelte/no-inline-styles": "error",
      "svelte/no-at-html-tags": "error",
      "svelte/no-at-debug-tags": "error",
      "svelte/no-unused-svelte-ignore": "error",
      "svelte/button-has-type": "error",
      "svelte/require-each-key": "error",
      "svelte/valid-each-key": "error",
      "svelte/no-reactive-functions": "error",
      "svelte/no-ignored-unsubscribe": "error",
      "svelte/no-add-event-listener": "error",
      "svelte/no-unused-props": "error",
      "svelte/require-event-dispatcher-types": "error",
      "svelte/no-target-blank": "error",
      "svelte/no-navigation-without-resolve": "error",
      "svelte/no-navigation-without-base": "error",
      "svelte/no-goto-without-base": "error",
      "svelte/no-export-load-in-svelte-module-in-kit-pages": "error",
      "svelte/valid-prop-names-in-kit-pages": "error",
      "svelte/no-store-async": "error",
      "svelte/no-reactive-reassign": "error",
      "svelte/no-immutable-reactive-statements": "error",
      "svelte/no-dupe-on-directives": "error",
      "svelte/no-dupe-use-directives": "error",
      "svelte/no-dupe-else-if-blocks": "error",
      "svelte/no-dupe-style-properties": "error",
      "svelte/no-unknown-style-directive-property": "error",
      "svelte/no-not-function-handler": "error",
      "svelte/no-object-in-text-mustaches": "error",
      "svelte/valid-compile": "error",
      "svelte/valid-style-parse": "error",
      "svelte/block-lang": ["error", { script: "ts" }],
    },
  },
  {
    files: ["**/*.css"],
    extends: [eslintPluginBetterTailwindcss.configs.recommended],
    settings: betterTailwindSettings,
    language: "css/css",
    languageOptions: {
      customSyntax: tailwind4,
      tolerant: true,
    },
    plugins: {
      css,
    },
  },
);
