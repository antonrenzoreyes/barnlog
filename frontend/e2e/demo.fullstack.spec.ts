import { expect, test } from "@playwright/test";

test("fullstack: backend health and frontend home are reachable", async ({ page, request }) => {
  const backendURL = process.env.PLAYWRIGHT_BACKEND_URL ?? "http://127.0.0.1:8081";
  const backendHealth = await request.get(`${backendURL}/healthz`);
  expect(backendHealth.ok()).toBe(true);

  await page.goto("/");
  await expect(page.locator("h1")).toBeVisible();
});
