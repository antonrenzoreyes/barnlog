import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-badge.fixture.svelte";

describe("BBadge", () => {
  it("renders badge content with base class and class passthrough", async () => {
    render(Fixture);

    const badge = page.getByText("Badge Content");
    await expect.element(badge).toBeVisible();
    await expect.element(badge).toHaveClass(/b-badge/);
    await expect.element(badge).toHaveClass(/fixture-badge/);
  });
});
