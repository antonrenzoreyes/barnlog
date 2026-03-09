import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-icon.fixture.svelte";

describe("BIcon", () => {
  it("renders mapped icon names with expected classes", async () => {
    render(Fixture);

    const inlineIcon = page.getByRole("img", { name: "Search icon" });
    await expect.element(inlineIcon).toBeInTheDocument();
    await expect.element(inlineIcon).toHaveClass(/icon-inline/);

    const displayIcon = page.getByRole("img", { name: "App icon" });
    await expect.element(displayIcon).toBeInTheDocument();
    await expect.element(displayIcon).toHaveClass(/icon-display/);
  });
});
