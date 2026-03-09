import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-choose.fixture.svelte";

describe("BChoose", () => {
  it("renders options and keeps the seeded selection", async () => {
    render(Fixture);

    const goat = page.getByRole("radio", { name: /^Goat$/ });
    const pig = page.getByRole("radio", { name: /^Pig$/ });

    const goatId = goat.element().getAttribute("id") ?? "";
    expect(goatId).toMatch(/^.+-fixture-species-goat$/);
    await expect.element(goat).toBeVisible();
    await expect.element(pig).toBeChecked();
    await expect.element(page.getByText("Selected species: pig")).toBeVisible();
    await expect.element(page.getByText(/^Species\s*\*/)).toBeVisible();
  });

  it("uses explicit option ids when duplicate values exist", async () => {
    render(Fixture);

    const goatSheep = page.getByLabelText("Goat / Sheep");
    const firstGoatSheepId = goatSheep.nth(0).element().getAttribute("id") ?? "";
    const secondGoatSheepId = goatSheep.nth(1).element().getAttribute("id") ?? "";
    expect(firstGoatSheepId).toMatch(/^.+-fixture-species-goat-sheep-first$/);
    expect(firstGoatSheepId).not.toBe(secondGoatSheepId);

    await goatSheep.nth(1).click();

    await expect.element(goatSheep.nth(0)).not.toBeChecked();
    await expect.element(goatSheep.nth(1)).toBeChecked();
    await expect.element(page.getByText("Selected species: goat / sheep")).toBeVisible();
  });
});
