import { describe, expect, it } from "vitest";
import { page } from "vitest/browser";
import { render } from "vitest-browser-svelte";
import Fixture from "./b-file-uploader.fixture.svelte";

const simulateFileSelection = (input: HTMLInputElement): void => {
  const dataTransfer = new DataTransfer();
  dataTransfer.items.add(new File(["one"], "one.png", { type: "image/png" }));
  dataTransfer.items.add(new File(["two"], "two.png", { type: "image/png" }));
  const filesDescriptor = Object.getOwnPropertyDescriptor(HTMLInputElement.prototype, "files");
  if (filesDescriptor && filesDescriptor.set) {
    filesDescriptor.set.call(input, dataTransfer.files);
  }
  input.dispatchEvent(new Event("change", { bubbles: true }));
};

describe("BFileUploader", () => {
  it("renders a visible trigger and forwards native file input attrs", async () => {
    render(Fixture);

    await expect.element(page.getByText("Choose file").first()).toBeVisible();

    const attachments = page.getByLabelText("Attachments");
    await expect.element(attachments).toHaveAttribute("type", "file");
    await expect.element(attachments).toHaveAttribute("accept", "image/*");
    await expect.element(attachments).toHaveAttribute("name", "attachments");
    await expect.element(attachments).toHaveAttribute("multiple", "");
    await expect.element(attachments).toHaveClass(/b-file-uploader-input/);
  });

  it("applies shared field-control class when wrapped by an error field", async () => {
    render(Fixture);
    const proof = page.getByLabelText("Proof");
    await expect.element(proof).toHaveClass(/b-field-control-input/);
    await expect.element(page.getByText("At least one file is required")).toBeVisible();
  });

  it("updates selected file list and calls onFilesChange on selection", async () => {
    render(Fixture);

    const attachments = document.querySelector<HTMLInputElement>('input[name="attachments"]');
    expect(attachments).toBeTruthy();
    if (!attachments) {
      return;
    }
    simulateFileSelection(attachments);

    await expect.element(page.getByText(/^one\.png$/)).toBeVisible();
    await expect.element(page.getByText(/^two\.png$/)).toBeVisible();
    await expect.element(page.getByText("Attachments callback count: 1")).toBeVisible();
    await expect
      .element(page.getByText("Attachments callback files: one.png, two.png"))
      .toBeVisible();
  });
});
