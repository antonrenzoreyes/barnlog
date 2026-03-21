export const uploadErrorCode = {
  fileRequired: "file_required",
  fileTooLarge: "file_too_large",
  internalError: "internal_error",
  invalidFile: "invalid_file",
  invalidMultipart: "invalid_multipart",
  invalidResponse: "invalid_response",
  multipleFilesNotAllowed: "multiple_files_not_allowed",
  networkError: "network_error",
  unknownError: "unknown_error",
  unsupportedFileType: "unsupported_file_type",
  uploadAborted: "upload_aborted",
} as const;

export type UploadErrorCode = (typeof uploadErrorCode)[keyof typeof uploadErrorCode];

const knownBackendUploadErrorCodes = new Set<UploadErrorCode>([
  uploadErrorCode.fileRequired,
  uploadErrorCode.multipleFilesNotAllowed,
  uploadErrorCode.unsupportedFileType,
  uploadErrorCode.fileTooLarge,
  uploadErrorCode.invalidMultipart,
  uploadErrorCode.invalidFile,
  uploadErrorCode.internalError,
]);

export const parseUploadErrorCode = (code: unknown): UploadErrorCode | undefined => {
  if (typeof code !== "string") {
    return undefined;
  }

  if (!knownBackendUploadErrorCodes.has(code as UploadErrorCode)) {
    return undefined;
  }

  return code as UploadErrorCode;
};

const uploadErrorMessages: Record<UploadErrorCode, string> = {
  [uploadErrorCode.uploadAborted]: "Upload canceled.",
  [uploadErrorCode.fileRequired]: "Please choose a file before uploading.",
  [uploadErrorCode.multipleFilesNotAllowed]: "Only one file can be uploaded at a time.",
  [uploadErrorCode.unsupportedFileType]:
    "Unsupported file type. Please upload a JPG, PNG, WEBP, or GIF image.",
  [uploadErrorCode.fileTooLarge]: "File is too large. Please upload a smaller image.",
  [uploadErrorCode.invalidMultipart]: "Upload request was invalid. Please try again.",
  [uploadErrorCode.invalidFile]: "Uploaded file is invalid. Please choose a different file.",
  [uploadErrorCode.internalError]: "Upload failed due to a server error. Please try again.",
  [uploadErrorCode.networkError]:
    "Network error while uploading. Check your connection and try again.",
  [uploadErrorCode.invalidResponse]: "Upload failed because of an unexpected server response.",
  [uploadErrorCode.unknownError]: "Upload failed. Please try again.",
};

export const uploadErrorMessageFromCode = (code: UploadErrorCode): string =>
  uploadErrorMessages[code] ?? uploadErrorMessages[uploadErrorCode.unknownError];
