<script lang="ts">
  import { parseDate } from "@internationalized/date";
  import { Pencil, Plus } from "lucide-svelte";
  import {
    BBadge,
    BButton,
    BCard,
    BDate,
    BField,
    BFileUploader,
    BInput,
    BSelect,
  } from "$lib";

  const sampleBirthDate = parseDate("2026-03-10");
  /* oxlint-disable-next-line eslint/prefer-const */
  let sampleSpecies = $state("goat");
  /* oxlint-disable-next-line eslint/prefer-const */
  let selectedUploadMessage = $state("No file selected yet.");

  const speciesOptions = [
    { label: "Goat", value: "goat" },
    { label: "Pig", value: "pig" },
    { label: "Dog", value: "dog" },
    { label: "Cat", value: "cat" },
    {
      description: "Will be available after species rollout.",
      disabled: true,
      label: "Horse (coming soon)",
      value: "horse",
    },
  ];

  const handlePhotoSelection = (files: File[]): void => {
    const [firstFile, secondFile] = files;

    if (firstFile) {
      const selectedNames = files.map((file) => `• ${file.name}`).join("\n");
      let selectedPrefix = "Selected file";

      if (secondFile) {
        selectedPrefix = "Selected files";
      }

      selectedUploadMessage = `${selectedPrefix}:\n${selectedNames}`;
      return;
    }

    selectedUploadMessage = "No file selected yet.";
  };
</script>

<svelte:head>
  <title>Barn Log Components</title>
</svelte:head>

<main
  class="mx-auto min-h-screen w-full max-w-lg bg-(--ui-color-panel) px-4 py-4 pb-12"
>
  <BCard as="header" class="mb-5 rounded-2xl">
    <div class="flex items-center gap-3">
      <div
        class="grid h-11 w-11 place-items-center rounded-xl border border-(--ui-color-border) bg-(--ui-color-primary-soft)"
      >
        <span
          aria-hidden="true"
          class="text-sm font-bold tracking-[0.05em] text-(--ui-color-primary-strong)"
        >
          BL
        </span>
      </div>
      <div>
        <h1 class="text-2xl font-semibold leading-tight text-(--ui-color-text)">
          Barn Log Components
        </h1>
        <p class="text-sm text-(--ui-color-text-soft)">
          Foundation forms for BInput, BDate, BField, and BFileUploader.
        </p>
      </div>
    </div>
  </BCard>

  <BCard as="section" aria-label="BField and BInput examples">
    <div class="mb-4 flex items-center justify-between gap-2">
      <div>
        <p
          class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
        >
          Form Primitives
        </p>
        <h2 class="text-lg font-semibold text-(--ui-color-text)">
          BField + BInput + BDate
        </h2>
      </div>
      <BBadge text="Preview" tone="primary" />
    </div>

    <div class="space-y-5">
      <BField
        id="dev-bfield-name"
        hint="Use the animal's display name."
        label="Name"
        required
      >
        {#snippet children({ id, describedBy, invalid })}
          <BInput
            {id}
            aria-describedby={describedBy}
            {invalid}
            placeholder="e.g., Clover"
            value="Mabel"
          />
        {/snippet}
      </BField>

      <BField
        error="Tag must be at least 2 characters."
        hint="Farm short code used for quick lookup."
        id="dev-bfield-tag"
        label="Tag"
      >
        {#snippet children({ id, describedBy, invalid })}
          <BInput
            {id}
            aria-describedby={describedBy}
            {invalid}
            placeholder="e.g., G-12"
            value="A"
          />
        {/snippet}
      </BField>

      <BField
        hint="Date opens a mobile-first picker popover."
        id="dev-bfield-birth-date"
        label="Birth date"
      >
        {#snippet children({ id, describedBy, invalid })}
          <BDate
            {id}
            aria-describedby={describedBy}
            {invalid}
            placeholder="Pick a date"
            value={sampleBirthDate}
          />
        {/snippet}
      </BField>
    </div>
  </BCard>

  <BCard as="section" aria-label="BBadge examples" class="mt-4">
    <div class="mb-4">
      <p
        class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
      >
        Display Tokens
      </p>
      <h2 class="text-lg font-semibold text-(--ui-color-text)">BBadge</h2>
    </div>

    <div class="flex flex-wrap gap-2">
      <BBadge text="Goat" />
      <BBadge text="Checked In" tone="primary" />
      <BBadge text="Needs Review" tone="danger" />
    </div>
  </BCard>

  <BCard as="section" aria-label="BButton examples" class="mt-4">
    <div class="mb-4">
      <p
        class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
      >
        Actions
      </p>
      <h2 class="text-lg font-semibold text-(--ui-color-text)">BButton</h2>
    </div>

    <div class="flex flex-wrap items-center gap-2">
      <BButton tone="primary">Save Animal</BButton>
      <BButton>Cancel</BButton>
      <BButton tone="danger">Archive</BButton>
      <BButton>View Notes</BButton>
    </div>

    <div class="mt-4 flex items-center gap-3">
      <BButton aria-label="Add animal" size="fab" tone="primary">
        <Plus aria-hidden="true" size={22} strokeWidth={2.4} />
      </BButton>
      <BButton aria-label="Edit animal" size="fab">
        <Pencil aria-hidden="true" size={20} strokeWidth={2.3} />
      </BButton>
    </div>
  </BCard>

  <BCard as="section" aria-label="BCard footer sample" class="mt-4">
    <div class="mb-4">
      <p
        class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
      >
        Footer Actions
      </p>
      <h2 class="text-lg font-semibold text-(--ui-color-text)">BCard Footer</h2>
      <p class="mt-1 text-sm text-(--ui-color-text-soft)">
        Edit-animal style action row built with plain card layout.
      </p>
    </div>

    <div class="border-t border-(--ui-color-border) pt-3">
      <div class="grid grid-cols-2 gap-2">
        <BButton>Cancel Footer Sample</BButton>
        <BButton tone="primary">Save Footer Sample</BButton>
      </div>
    </div>
  </BCard>

  <BCard as="section" aria-label="BSelect examples" class="mt-4">
    <div class="mb-4">
      <p
        class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
      >
        Choice Inputs
      </p>
      <h2 class="text-lg font-semibold text-(--ui-color-text)">BSelect</h2>
    </div>

    <div class="space-y-5">
      <BField
        hint="Used for filtering and species-specific treatment records."
        id="dev-bfield-species"
        label="Species"
        required
      >
        {#snippet children({ id, describedBy, invalid })}
          <BSelect
            aria-describedby={describedBy}
            {id}
            {invalid}
            name="species"
            options={speciesOptions}
            placeholder="Choose species"
            bind:value={sampleSpecies}
          />
        {/snippet}
      </BField>

      <BField
        error="Please select a species before saving."
        id="dev-bfield-species-invalid"
        label="Species (invalid)"
      >
        {#snippet children({ id, describedBy, invalid })}
          <BSelect
            aria-describedby={describedBy}
            {id}
            {invalid}
            options={speciesOptions}
            placeholder="Select species"
            value=""
          />
        {/snippet}
      </BField>
    </div>
  </BCard>

  <BCard as="section" aria-label="BFileUploader examples" class="mt-4">
    <div class="mb-4">
      <p
        class="text-[0.72rem] font-semibold uppercase tracking-[0.08em] text-(--ui-color-text-soft)"
      >
        Media Input
      </p>
      <h2 class="text-lg font-semibold text-(--ui-color-text)">
        BFileUploader
      </h2>
    </div>

    <BField
      hint="Mobile-first attach button powered by Uppy restrictions."
      id="dev-bfield-photo"
      label="Profile photo"
      required
    >
      {#snippet children({ id, describedBy, invalid })}
        <BFileUploader
          accept="image/*"
          aria-describedby={describedBy}
          buttonLabel="Upload photos"
          {id}
          {invalid}
          multiple
          name="photo"
          onFilesChange={handlePhotoSelection}
        />
      {/snippet}
    </BField>

    <p class="mt-3 whitespace-pre-line text-sm text-(--ui-color-text-soft)">
      {selectedUploadMessage}
    </p>
  </BCard>
</main>
