import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./page.fixture.svelte";

describe("/+page.svelte", () => {
  it("should render h1", async () => {
    render(Fixture);

    const heading = page.getByRole("heading", { level: 1, name: "Barn Log Frontend" });
    await expect.element(heading).toBeInTheDocument();
  });
});
