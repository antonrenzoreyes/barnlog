import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-field.fixture.svelte";

describe("BField", () => {
  it("links label and hint semantics to the first control", async () => {
    render(Fixture);

    const animalLabel = page.getByText("Animal name");
    await expect.element(animalLabel).toBeVisible();
    const animalName = page.getByLabelText("Animal name");
    await expect.element(animalName).toBeVisible();
    await expect.element(animalName).toHaveAttribute("id", "animal-name");
    await expect.element(animalName).toHaveAttribute("aria-describedby", "animal-name-message");
    await expect.element(page.getByText("Use the barn tag name")).toBeVisible();
    await expect
      .element(page.getByText("Use the barn tag name"))
      .toHaveAttribute("id", "animal-name-message");
  });

  it("applies required and error semantics to the first control", async () => {
    render(Fixture);

    const weight = page.getByLabelText("Weight");
    await expect.element(weight).toHaveAttribute("required", "");
    await expect.element(weight).toHaveAttribute("aria-describedby", "weight-message");
    await expect.element(page.getByText("Weight is required")).toBeVisible();
    await expect
      .element(page.getByText("Weight is required"))
      .toHaveAttribute("id", "weight-message");
    await expect.element(page.getByText("Use kg")).not.toBeInTheDocument();
  });
});
