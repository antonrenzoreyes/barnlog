import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-button.fixture.svelte";

describe("BButton", () => {
  it("renders variants and disabled state", async () => {
    render(Fixture);

    const primary = page.getByRole("button", { name: "Primary Action" });
    await expect.element(primary).toBeVisible();
    await expect.element(primary).toHaveClass(/btn-primary/);

    const disabledSecondary = page.getByRole("button", { name: "Disabled Secondary" });
    await expect.element(disabledSecondary).toBeVisible();
    await expect.element(disabledSecondary).toBeDisabled();
    await expect.element(disabledSecondary).toHaveClass(/btn-disabled/);
  });
});
