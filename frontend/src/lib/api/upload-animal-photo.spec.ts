import { afterEach, describe, expect, it, vi } from "vitest";
import { uploadAnimalPhoto } from "./upload-animal-photo";
import { uploadErrorCode } from "./upload-errors";

const makeFile = (name = "pepper.png", type = "image/png"): File =>
  new File(["image-content"], name, { type });

const stubFetch = (response: unknown) => {
  const fetchMock = vi.fn().mockResolvedValue(response);
  vi.stubGlobal("fetch", fetchMock);
  return fetchMock;
};

const expectUploadRequest = (fetchMock: ReturnType<typeof vi.fn>, file: File) => {
  expect(fetchMock).toHaveBeenCalledTimes(1);
  const [url, init] = fetchMock.mock.calls[0] as [string, RequestInit];
  expect(url).toBe("/uploads/animal-photos");
  expect(init.method).toBe("POST");
  expect(init.body).toBeInstanceOf(FormData);
  expect((init.body as FormData).get("file")).toBe(file);
};

afterEach(() => {
  vi.restoreAllMocks();
  vi.unstubAllGlobals();
});

describe("uploadAnimalPhoto success", () => {
  it("uploads multipart file and returns parsed response", async () => {
    const fetchMock = stubFetch({
      json: vi.fn().mockResolvedValue({
        content_type: "image/png",
        file_id: "file_abc123",
        file_name: "pepper.png",
        size_bytes: 1024,
      }),
      ok: true,
      status: 201,
    });
    const file = makeFile();

    const result = await uploadAnimalPhoto(file);

    expectUploadRequest(fetchMock, file);
    expect(result).toEqual({
      data: {
        content_type: "image/png",
        file_id: "file_abc123",
        file_name: "pepper.png",
        size_bytes: 1024,
      },
      ok: true,
    });
  });
});

describe("uploadAnimalPhoto invalid success payload (missing file_id)", () => {
  it("returns invalid_response", async () => {
    stubFetch({
      json: vi.fn().mockResolvedValue({
        content_type: "image/png",
        file_name: "pepper.png",
        photo_id: "legacy_photo_123",
        size_bytes: 1024,
      }),
      ok: true,
      status: 201,
    });

    const result = await uploadAnimalPhoto(makeFile());
    expect(result).toEqual({
      code: uploadErrorCode.invalidResponse,
      ok: false,
      status: 201,
    });
  });
});

describe("uploadAnimalPhoto invalid success payload (missing file_name)", () => {
  it("returns invalid_response", async () => {
    stubFetch({
      json: vi.fn().mockResolvedValue({
        content_type: "image/png",
        file_id: "file_abc123",
        size_bytes: 1024,
      }),
      ok: true,
      status: 201,
    });

    const result = await uploadAnimalPhoto(makeFile());
    expect(result).toEqual({
      code: uploadErrorCode.invalidResponse,
      ok: false,
      status: 201,
    });
  });
});

describe("uploadAnimalPhoto failure from backend error code", () => {
  it("returns parsed backend code", async () => {
    stubFetch({
      json: vi.fn().mockResolvedValue({ error: "unsupported_file_type" }),
      ok: false,
      status: 400,
    });

    const result = await uploadAnimalPhoto(makeFile("notes.txt", "text/plain"));
    expect(result).toEqual({
      code: uploadErrorCode.unsupportedFileType,
      ok: false,
      status: 400,
    });
  });
});

describe("uploadAnimalPhoto failure from request size", () => {
  it("maps status 413 to file_too_large when body is not parseable", async () => {
    stubFetch({
      json: vi.fn().mockRejectedValue(new Error("invalid json")),
      ok: false,
      status: 413,
    });

    const result = await uploadAnimalPhoto(makeFile());
    expect(result).toEqual({
      code: uploadErrorCode.fileTooLarge,
      ok: false,
      status: 413,
    });
  });
});

describe("uploadAnimalPhoto network failure", () => {
  it("returns upload_aborted when request is canceled", async () => {
    const fetchMock = vi
      .fn()
      .mockRejectedValue(new DOMException("The operation was aborted.", "AbortError"));
    vi.stubGlobal("fetch", fetchMock);

    const result = await uploadAnimalPhoto(makeFile(), { signal: new AbortController().signal });
    expect(result).toEqual({
      code: uploadErrorCode.uploadAborted,
      ok: false,
    });
  });

  it("returns network_error when request fails before response", async () => {
    const fetchMock = vi.fn().mockRejectedValue(new TypeError("network down"));
    vi.stubGlobal("fetch", fetchMock);

    const result = await uploadAnimalPhoto(makeFile());
    expect(result).toEqual({
      code: uploadErrorCode.networkError,
      ok: false,
    });
  });
});
