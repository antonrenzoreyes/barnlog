import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-card.fixture.svelte";

describe("BCard", () => {
  it("renders snippet sections and falls back to children when body is absent", async () => {
    render(Fixture);

    await expect.element(page.getByText("Card Header")).toBeVisible();
    await expect.element(page.getByText("Body Snippet")).toBeVisible();
    await expect.element(page.getByRole("button", { name: "Footer Action" })).toBeVisible();
    await expect.element(page.getByText("Should not render")).not.toBeInTheDocument();
    await expect.element(page.getByText("Children Content")).toBeVisible();
  });
});
