<script lang="ts">
  import BField from "./b-field.svelte";
  import BFileUploader from "./b-file-uploader.svelte";

  const fixtureState = $state({
    changeCount: 0,
    latestFileNames: "none",
  });
</script>

<div>
  <BField
    label="Attachments"
    for="fixture-attachments"
    hint="Upload one or more image files"
  >
    {#snippet children(field)}
      <BFileUploader
        {...field}
        accept="image/*"
        multiple={true}
        name="attachments"
        onFilesChange={(files) => {
          fixtureState.changeCount += 1;
          if (!files || files.length === 0) {
            fixtureState.latestFileNames = "none";
            return;
          }
          fixtureState.latestFileNames = [...files]
            .map((file) => file.name)
            .join(", ");
        }}
      />
    {/snippet}
  </BField>

  <BField
    class="mt-4"
    error="At least one file is required"
    for="fixture-proof"
    label="Proof"
  >
    {#snippet children(field)}
      <BFileUploader {...field} name="proof" />
    {/snippet}
  </BField>

  <p>Attachments callback count: {fixtureState.changeCount}</p>
  <p>Attachments callback files: {fixtureState.latestFileNames}</p>
</div>
