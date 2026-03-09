import { expect, test } from "@playwright/test";

test("home route loads", async ({ page }) => {
  await page.goto("/");
  await expect(page.getByRole("heading", { level: 1, name: "Barn Log Frontend" })).toBeVisible();
});

test("required field indicator is visible and uses danger color", async ({ page }) => {
  await page.goto("/dev/components");

  const requiredIndicator = page.locator("#bchoose .b-field-required-indicator");
  await expect(requiredIndicator).toBeVisible();
  await expect(requiredIndicator).toHaveText("*");
  await expect(requiredIndicator).toHaveCSS("color", "rgb(143, 59, 47)");
});
