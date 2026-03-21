import { parseUploadErrorCode, uploadErrorCode } from "./upload-errors";
import type { components } from "./generated/types";

type ErrorResponse = components["schemas"]["httpapi.errorResponse"];
type UploadErrorCode = (typeof uploadErrorCode)[keyof typeof uploadErrorCode];

export interface UploadedAnimalPhoto {
  content_type: string;
  file_id: string;
  file_name: string;
  size_bytes: number;
}

interface UploadAnimalPhotoSuccess {
  data: UploadedAnimalPhoto;
  ok: true;
}

interface UploadAnimalPhotoFailure {
  code: UploadErrorCode;
  ok: false;
  status?: number;
}

export type UploadAnimalPhotoResult = UploadAnimalPhotoFailure | UploadAnimalPhotoSuccess;

interface UploadAnimalPhotoOptions {
  signal?: AbortSignal;
}

interface UploadResponsePayload {
  content_type?: unknown;
  file_id?: unknown;
  file_name?: unknown;
  size_bytes?: unknown;
}

interface UploadRequestSuccess {
  response: Response;
}

interface UploadRequestFailure {
  failure: UploadAnimalPhotoFailure;
}

type UploadRequestResult = UploadRequestFailure | UploadRequestSuccess;

const endpointPath = "/uploads/animal-photos";
const fileFieldName = "file";
const postMethod = "POST";
const requestEntityTooLargeStatus = 413;

const isAbortError = (error: unknown): boolean => {
  if (error instanceof DOMException && error.name === "AbortError") {
    return true;
  }
  if (typeof error !== "object" || error === null) {
    return false;
  }
  return "name" in error && (error as { name?: unknown }).name === "AbortError";
};

const createUploadFormData = (file: File): FormData => {
  const formData = new FormData();
  formData.set(fileFieldName, file);
  return formData;
};

const getUploadAnimalPhotoUrl = (): string => {
  const base = import.meta.env.PUBLIC_API_BASE_URL?.trim() ?? "";
  const normalizedBase = base.replace(/\/+$/, "");

  if (normalizedBase === "") {
    return endpointPath;
  }

  return `${normalizedBase}${endpointPath}`;
};

const readString = (value: unknown): string | undefined => {
  if (typeof value === "string") {
    return value;
  }

  return undefined;
};

const readNumber = (value: unknown): number | undefined => {
  if (typeof value === "number") {
    return value;
  }

  return undefined;
};

const parseUploadedAnimalPhoto = (payload: unknown): UploadedAnimalPhoto | undefined => {
  if (!payload || typeof payload !== "object") {
    return undefined;
  }

  const response = payload as UploadResponsePayload;
  const contentType = readString(response.content_type);
  const fileID = readString(response.file_id);
  const fileName = readString(response.file_name);
  const sizeBytes = readNumber(response.size_bytes);

  if (!contentType || !fileID || !fileName || sizeBytes === undefined) {
    return undefined;
  }

  return {
    content_type: contentType,
    file_id: fileID,
    file_name: fileName,
    size_bytes: sizeBytes,
  };
};

const parseErrorCodeFromResponse = async (
  response: Response,
): Promise<UploadErrorCode | undefined> => {
  try {
    const payload = (await response.json()) as ErrorResponse;
    return parseUploadErrorCode(payload?.error);
  } catch {
    return undefined;
  }
};

const sendUploadRequest = async (
  file: File,
  signal?: AbortSignal,
): Promise<UploadRequestResult> => {
  try {
    const response = await fetch(getUploadAnimalPhotoUrl(), {
      body: createUploadFormData(file),
      method: postMethod,
      signal,
    });
    return { response };
  } catch (error) {
    if (isAbortError(error)) {
      return { failure: { code: uploadErrorCode.uploadAborted, ok: false } };
    }
    return { failure: { code: uploadErrorCode.networkError, ok: false } };
  }
};

const parseUploadSuccess = async (response: Response): Promise<UploadAnimalPhotoResult> => {
  try {
    const payload = await response.json();
    const parsed = parseUploadedAnimalPhoto(payload);

    if (!parsed) {
      return {
        code: uploadErrorCode.invalidResponse,
        ok: false,
        status: response.status,
      };
    }

    return { data: parsed, ok: true };
  } catch {
    return {
      code: uploadErrorCode.invalidResponse,
      ok: false,
      status: response.status,
    };
  }
};

const parseUploadFailure = async (response: Response): Promise<UploadAnimalPhotoFailure> => {
  const parsedCode = await parseErrorCodeFromResponse(response);
  if (parsedCode) {
    return { code: parsedCode, ok: false, status: response.status };
  }

  if (response.status === requestEntityTooLargeStatus) {
    return { code: uploadErrorCode.fileTooLarge, ok: false, status: response.status };
  }

  return { code: uploadErrorCode.unknownError, ok: false, status: response.status };
};

export const uploadAnimalPhoto = async (
  file: File,
  options?: UploadAnimalPhotoOptions,
): Promise<UploadAnimalPhotoResult> => {
  const result = await sendUploadRequest(file, options?.signal);

  if ("failure" in result) {
    return result.failure;
  }
  if (result.response.ok) {
    return parseUploadSuccess(result.response);
  }

  return parseUploadFailure(result.response);
};
