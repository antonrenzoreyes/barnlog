<script lang="ts">
  import BIcon from "./b-icon.svelte";
  import type { HTMLInputAttributes } from "svelte/elements";
  /**
   * BFileUploader wraps a native file input with consistent styling and
   * optional selected-file preview text.
   */
  interface BFileUploaderProps extends Omit<
    HTMLInputAttributes,
    "children" | "type" | "files"
  > {
    files?: FileList | null;
    onFilesChange?: (files: FileList | null) => void;
    showSelectedFiles?: boolean;
    triggerLabel?: string;
  }

  // oxlint-disable prefer-const -- bindable prop must remain writable.
  let {
    id,
    files = $bindable<FileList | null>(),
    onchange,
    onFilesChange,
    showSelectedFiles = true,
    triggerLabel = "Choose file",
    class: className = "",
    ...rest
  }: BFileUploaderProps = $props();
  // oxlint-enable prefer-const

  const uploaderState = $state({ fileNames: [] as string[] });
  const uid = $props.id();
  // Intentional divergence: Native file inputs are not reliably styleable
  // Across browsers, so we pair a hidden input with a styled trigger label.
  const inputId = $derived(id ?? `${uid}-file-input`);

  const handleChange: NonNullable<HTMLInputAttributes["onchange"]> = (
    event,
  ): void => {
    const { files: nextFiles } = event.currentTarget as HTMLInputElement;
    files = nextFiles;
    if (onFilesChange) {
      onFilesChange(nextFiles);
    }
    if (onchange) {
      onchange(event);
    }
  };

  $effect(() => {
    const currentFiles = files;
    if (!currentFiles) {
      uploaderState.fileNames = [];
      return;
    }
    uploaderState.fileNames = [...currentFiles].map((file) => file.name);
  });
</script>

<div class={`b-file-uploader ${className}`}>
  <input
    {...rest}
    bind:files
    id={inputId}
    class="b-file-uploader-input b-field-control-input"
    onchange={handleChange}
    type="file"
  />
  <label class="b-file-uploader-trigger b-field-control-input" for={inputId}>
    <BIcon name="uploadPhoto" size="inline" />
    <span>{triggerLabel}</span>
  </label>

  {#if showSelectedFiles && uploaderState.fileNames.length > 0}
    <ul class="b-file-uploader-list">
      {#each uploaderState.fileNames as fileName, index (`${index}-${fileName}`)}
        <li>{fileName}</li>
      {/each}
    </ul>
  {/if}
</div>

<style>
  .b-file-uploader {
    display: grid;
    gap: 0.5rem;
  }

  .b-file-uploader-input {
    position: absolute;
    width: 1px;
    height: 1px;
    margin: -1px;
    overflow: hidden;
    border: 0;
    padding: 0;
    clip: rect(0 0 0 0);
    clip-path: inset(50%);
    white-space: nowrap;
  }

  .b-file-uploader-trigger {
    display: inline-flex;
    min-height: 3rem;
    width: 100%;
    cursor: pointer;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    border-radius: var(--radius-barn-md);
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    padding: 0.5rem 0.75rem;
    color: var(--color-barn-text);
    font-weight: 500;
  }

  .b-file-uploader-input:focus-visible + .b-file-uploader-trigger {
    outline: 2px solid color-mix(in srgb, var(--color-barn-primary) 45%, white);
    outline-offset: 1px;
  }

  .b-file-uploader-list {
    display: grid;
    gap: 0.25rem;
    margin: 0;
    padding: 0;
    list-style: none;
    font-size: 0.75rem;
    color: var(--color-barn-muted);
  }
</style>
