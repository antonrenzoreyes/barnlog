import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-input.fixture.svelte";

describe("BInput", () => {
  it("renders text-like inputs with forwarded attrs and shared field-control class", async () => {
    render(Fixture);

    const animalName = page.getByRole("textbox", { name: "Animal name" });
    await expect.element(animalName).toBeVisible();
    await expect.element(animalName).toHaveAttribute("name", "animalName");
    await expect.element(animalName).toHaveAttribute("placeholder", "Bella");
    await expect.element(animalName).toHaveClass(/b-input/);

    const weight = page.getByRole("spinbutton", { name: "Weight" });
    await expect.element(weight).toBeVisible();
    await expect.element(weight).toHaveClass(/b-field-control-input/);
    await expect.element(page.getByText("Weight is required")).toBeVisible();
  });
});
