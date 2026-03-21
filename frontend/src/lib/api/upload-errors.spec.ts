import { describe, expect, it } from "vitest";
import { parseUploadErrorCode, uploadErrorCode, uploadErrorMessageFromCode } from "./upload-errors";

describe("parseUploadErrorCode", () => {
  it("returns known backend upload error code", () => {
    const code = parseUploadErrorCode("unsupported_file_type");
    expect(code).toBe(uploadErrorCode.unsupportedFileType);
  });

  it("returns undefined for unknown code", () => {
    const code = parseUploadErrorCode("not_a_real_code");
    expect(code).toBeUndefined();
  });

  it("returns undefined for non-string code", () => {
    const code = parseUploadErrorCode(413);
    expect(code).toBeUndefined();
  });
});

describe("uploadErrorMessageFromCode", () => {
  it("maps upload_aborted to a user-facing message", () => {
    const message = uploadErrorMessageFromCode(uploadErrorCode.uploadAborted);
    expect(message).toContain("canceled");
  });

  it("maps file_too_large to a user-facing message", () => {
    const message = uploadErrorMessageFromCode(uploadErrorCode.fileTooLarge);
    expect(message).toContain("too large");
  });

  it("maps network_error to a user-facing message", () => {
    const message = uploadErrorMessageFromCode(uploadErrorCode.networkError);
    expect(message).toContain("Network error");
  });
});
