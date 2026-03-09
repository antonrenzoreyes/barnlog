<script lang="ts">
  import BField from "./b-field.svelte";
</script>

<div>
  <BField label="Animal name" for="animal-name" hint="Use the barn tag name">
    {#snippet children(field)}
      <input {...field} class="input-base" />
    {/snippet}
  </BField>

  <BField
    class="mt-4"
    label="Weight"
    for="weight"
    hint="Use kg"
    error="Weight is required"
    required={true}
  >
    {#snippet children(field)}
      <input {...field} class="input-base" />
    {/snippet}
  </BField>
</div>
