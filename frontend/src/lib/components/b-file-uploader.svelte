<script lang="ts">
  import Uppy, { type Body, type Meta } from "@uppy/core";
  import { Upload } from "lucide-svelte";
  import { onDestroy } from "svelte";
  import type { HTMLInputAttributes } from "svelte/elements";

  const DEFAULT_MAX_FILE_SIZE_BYTES = 10_485_760;
  const FALLBACK_ERROR_MESSAGE = "Could not attach file.";
  const EMPTY_FILE_COUNT = 0;
  const SINGLE_FILE_LIMIT = 1;

  interface Props {
    "aria-describedby"?: string;
    accept?: string;
    buttonLabel?: string;
    capture?: HTMLInputAttributes["capture"];
    disabled?: boolean;
    id?: string;
    invalid?: boolean;
    maxFileSizeBytes?: number;
    maxFiles?: number;
    multiple?: boolean;
    name?: string;
    onFilesChange?: (files: File[]) => void;
    onUploadError?: (message: string) => void;
    required?: boolean;
  }

  const props: Props = $props();
  /* oxlint-disable-next-line eslint/prefer-const */
  let inputElement = $state<HTMLInputElement>();
  let uploader = $state<Uppy<Meta, Body>>();
  let uploadErrorMessage = $state("");

  const allowedFileTypes = $derived.by(() => {
    if (!props.accept) {
      return;
    }

    const fileTypes = props.accept
      .split(",")
      .map((type) => type.trim())
      .filter(Boolean);

    if (fileTypes.length === EMPTY_FILE_COUNT) {
      return;
    }

    return fileTypes;
  });

  const maxFileSizeBytes = $derived(
    props.maxFileSizeBytes ?? DEFAULT_MAX_FILE_SIZE_BYTES,
  );

  const restrictions = $derived.by(() => {
    if (props.multiple) {
      return {
        allowedFileTypes,
        maxFileSize: maxFileSizeBytes,
        maxNumberOfFiles: props.maxFiles,
      };
    }

    return {
      allowedFileTypes,
      maxFileSize: maxFileSizeBytes,
      maxNumberOfFiles: props.maxFiles ?? SINGLE_FILE_LIMIT,
    };
  });

  const getUploader = (): Uppy<Meta, Body> => {
    if (!uploader) {
      uploader = new Uppy<Meta, Body>({
        allowMultipleUploadBatches: false,
        autoProceed: false,
        restrictions,
      });
    }

    return uploader;
  };

  $effect(() => {
    if (uploader) {
      uploader.setOptions({ restrictions });
    }
  });

  const openFileDialog = (): void => {
    inputElement?.click();
  };
  const getErrorMessage = (error: unknown): string => {
    if (error instanceof Error) {
      return error.message;
    }

    return FALLBACK_ERROR_MESSAGE;
  };
  const addCandidateFile = (
    nextUploader: Uppy<Meta, Body>,
    file: File,
  ): boolean => {
    try {
      nextUploader.addFile({
        data: file,
        name: file.name,
        type: file.type,
      });

      return true;
    } catch (error) {
      const errorMessage = getErrorMessage(error);
      uploadErrorMessage = errorMessage;
      props.onUploadError?.(errorMessage);
      return false;
    }
  };
  const getSelectedFiles = (nextUploader: Uppy<Meta, Body>): File[] =>
    nextUploader
      .getFiles()
      .map((selectedFile) => selectedFile.data)
      .filter(
        (selectedData): selectedData is File => selectedData instanceof File,
      );
  const collectAcceptedFiles = (
    nextFiles: File[],
  ): { acceptedFiles: File[]; hasRejectedFiles: boolean } => {
    const nextUploader = getUploader();
    uploadErrorMessage = "";
    let hasRejectedFiles = false;

    if (!props.multiple) {
      nextUploader.clear();
    }

    nextFiles.forEach((file) => {
      const wasAccepted = addCandidateFile(nextUploader, file);

      if (!wasAccepted) {
        hasRejectedFiles = true;
      }
    });

    return {
      acceptedFiles: getSelectedFiles(nextUploader),
      hasRejectedFiles,
    };
  };
  const handleInputChange = (event: Event): void => {
    const target = event.currentTarget;

    if (!(target instanceof HTMLInputElement)) {
      return;
    }

    const nextFiles = [...(target.files ?? [])];

    if (nextFiles.length === EMPTY_FILE_COUNT) {
      return;
    }

    const { acceptedFiles, hasRejectedFiles } = collectAcceptedFiles(nextFiles);

    if (hasRejectedFiles && props.name) {
      target.value = "";
    }

    props.onFilesChange?.(acceptedFiles);
  };

  onDestroy(() => {
    uploader?.destroy();
  });
</script>

<button
  aria-describedby={props["aria-describedby"]}
  class={[
    "inline-flex h-14 w-full items-center justify-center gap-2 rounded-xl border bg-(--ui-color-surface) px-4 text-base font-semibold transition-[background-color,border-color,color,box-shadow,transform] duration-150 ease-out focus-visible:outline-none active:translate-y-px disabled:cursor-not-allowed disabled:opacity-55 disabled:active:translate-y-0",
    props.invalid
      ? "border-(--ui-color-border-danger) text-(--ui-color-text-danger) focus-visible:shadow-[var(--ui-shadow-focus-danger)]"
      : "border-(--ui-color-border-strong) text-(--ui-color-text) hover:border-(--ui-color-primary) hover:bg-(--ui-color-primary-soft) focus-visible:shadow-[var(--ui-shadow-focus)]",
  ]}
  disabled={props.disabled}
  id={props.id}
  onclick={openFileDialog}
  type="button"
>
  <Upload aria-hidden="true" size={17} strokeWidth={2.3} />
  <span>{props.buttonLabel ?? "Upload file"}</span>
</button>

<input
  accept={props.accept}
  bind:this={inputElement}
  capture={props.capture}
  class="sr-only"
  disabled={props.disabled}
  multiple={props.multiple}
  name={props.name}
  onchange={handleInputChange}
  required={props.required}
  tabindex={-1}
  type="file"
/>

{#if uploadErrorMessage}
  <p
    class="mt-2 text-sm font-medium text-(--ui-color-text-danger)"
    role="alert"
  >
    {uploadErrorMessage}
  </p>
{/if}
