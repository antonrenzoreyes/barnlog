import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Page from "./+page.svelte";

const openBirthDatePicker = async (): Promise<void> => {
  const birthDateButton = page.getByRole("button", { name: "Birth date" });
  await birthDateButton.click();
};

const getTodayLabel = (): string =>
  new Intl.DateTimeFormat(undefined, {
    day: "numeric",
    month: "short",
    year: "numeric",
  }).format(new Date());

describe("/dev/components/+page.svelte rendering", () => {
  it("renders BField, BInput, and BDate examples", async () => {
    render(Page);

    const nameInput = page.getByPlaceholder("e.g., Clover");
    await expect.element(nameInput).toBeInTheDocument();

    const tagInput = page.getByPlaceholder("e.g., G-12");
    await expect.element(tagInput).toBeInTheDocument();

    const tagError = page.getByText("Tag must be at least 2 characters.");
    await expect.element(tagError).toBeVisible();

    const dateTrigger = page.getByText("Mar 10, 2026");
    await expect.element(dateTrigger).toBeVisible();
  });

  it("renders BBadge examples", async () => {
    render(Page);
    const badgeRegion = page.getByRole("region", { name: "BBadge examples" });

    await expect.element(badgeRegion.getByText("Goat")).toBeVisible();
    await expect.element(badgeRegion.getByText("Checked In")).toBeVisible();
    await expect.element(badgeRegion.getByText("Needs Review")).toBeVisible();
  });

  it("renders BButton actions and FABs", async () => {
    render(Page);
    const actionRegion = page.getByRole("region", { name: "BButton examples" });

    await expect.element(actionRegion.getByRole("button", { name: "Save Animal" })).toBeVisible();
    await expect.element(actionRegion.getByRole("button", { name: "Cancel" })).toBeVisible();
    await expect.element(actionRegion.getByRole("button", { name: "Archive" })).toBeVisible();
    await expect.element(actionRegion.getByRole("button", { name: "View Notes" })).toBeVisible();
    await expect.element(actionRegion.getByRole("button", { name: "Add animal" })).toBeVisible();
    await expect.element(actionRegion.getByRole("button", { name: "Edit animal" })).toBeVisible();
  });
});

describe("/dev/components/+page.svelte layout semantics", () => {
  it("renders semantic card regions", async () => {
    render(Page);

    await expect
      .element(page.getByRole("region", { name: "BField and BInput examples" }))
      .toBeVisible();
    await expect.element(page.getByRole("region", { name: "BBadge examples" })).toBeVisible();
    await expect.element(page.getByRole("region", { name: "BButton examples" })).toBeVisible();
    await expect.element(page.getByRole("region", { name: "BCard footer sample" })).toBeVisible();
    await expect.element(page.getByRole("region", { name: "BSelect examples" })).toBeVisible();
    await expect
      .element(page.getByRole("region", { name: "BFileUploader examples" }))
      .toBeVisible();
  });

  it("renders plain two-button card footer sample", async () => {
    render(Page);

    await expect.element(page.getByRole("button", { name: "Cancel Footer Sample" })).toBeVisible();
    await expect.element(page.getByRole("button", { name: "Save Footer Sample" })).toBeVisible();
  });
});

describe("/dev/components/+page.svelte select behavior", () => {
  it("renders and updates species select", async () => {
    render(Page);
    const selectRegion = page.getByRole("region", { name: "BSelect examples" });
    const speciesTrigger = selectRegion.getByRole("button", {
      name: "Species (required)",
    });

    await expect.element(speciesTrigger.getByText("Goat")).toBeVisible();
    await speciesTrigger.click();

    const pigOption = page.getByRole("option", { name: "Pig" });
    await pigOption.click();

    await expect.element(speciesTrigger.getByText("Pig")).toBeVisible();
  });
});

describe("/dev/components/+page.svelte file uploader", () => {
  it("renders mobile uploader button and status text", async () => {
    render(Page);
    const uploaderRegion = page.getByRole("region", {
      name: "BFileUploader examples",
    });
    const uploaderButton = uploaderRegion.getByRole("button");

    await expect.element(uploaderButton).toBeVisible();
    await expect.element(uploaderRegion.getByText("Upload photos")).toBeVisible();
    await expect.element(uploaderRegion.getByText("No file selected yet.")).toBeVisible();
  });
});

describe("/dev/components/+page.svelte date quick actions", () => {
  it("sets date to today", async () => {
    render(Page);
    const todayLabel = getTodayLabel();

    await openBirthDatePicker();

    const todayButton = page.getByRole("button", { name: "Today" });
    await todayButton.click();

    const todayText = page.getByText(todayLabel);
    await expect.element(todayText).toBeVisible();
  });

  it("clears the selected date", async () => {
    render(Page);

    await openBirthDatePicker();

    const clearButton = page.getByRole("button", { name: "Clear" });
    await expect.element(clearButton).toBeVisible();
    await clearButton.click();

    const emptyDateText = page.getByText("Pick a date");
    await expect.element(emptyDateText).toBeVisible();
  });
});

describe("/dev/components/+page.svelte date jump controls", () => {
  it("changes month via month select", async () => {
    render(Page);

    await openBirthDatePicker();

    const monthSelect = page.getByRole("combobox", { name: "Month" });
    await monthSelect.selectOptions("1");
    await expect.element(monthSelect).toHaveValue("1");
  });

  it("changes year via year select", async () => {
    render(Page);

    await openBirthDatePicker();

    const yearSelect = page.getByRole("combobox", { name: "Year" });
    await yearSelect.selectOptions("2025");
    await expect.element(yearSelect).toHaveValue("2025");
  });
});
