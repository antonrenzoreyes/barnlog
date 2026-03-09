import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-date.fixture.svelte";

describe("BDate", () => {
  it("renders date input with forwarded attributes and seeded value", async () => {
    render(Fixture);

    const birthDate = page.getByLabelText("Birth date");
    await expect.element(birthDate).toBeVisible();
    await expect.element(birthDate).toHaveAttribute("type", "date");
    await expect.element(birthDate).toHaveAttribute("name", "birthDate");
    await expect.element(birthDate).toHaveValue("2026-03-01");
    await expect.element(birthDate).toHaveClass(/b-date/);
  });

  it("renders error field-control styling for check date", async () => {
    render(Fixture);

    const checkDate = page.getByLabelText("Check date");
    await expect.element(checkDate).toBeVisible();
    await expect.element(checkDate).toHaveClass(/b-field-control-input/);
    await expect.element(page.getByText("Date is required")).toBeVisible();
  });

  it("updates bound value when birth date changes", async () => {
    render(Fixture);

    const birthDate = page.getByLabelText("Birth date");
    await birthDate.fill("2026-04-02");
    await expect.element(birthDate).toHaveValue("2026-04-02");
    await expect.element(page.getByText("Birth date value: 2026-04-02")).toBeVisible();
  });
});
