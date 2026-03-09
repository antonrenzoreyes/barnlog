<script lang="ts">
  import BField from "./b-field.svelte";
  import BInput from "./b-input.svelte";
</script>

<div>
  <BField
    label="Animal name"
    for="fixture-animal-name"
    hint="Use a short, readable name"
  >
    {#snippet children(field)}
      <BInput {...field} name="animalName" placeholder="Bella" type="text" />
    {/snippet}
  </BField>

  <BField
    class="mt-4"
    error="Weight is required"
    for="fixture-weight"
    label="Weight"
  >
    {#snippet children(field)}
      <BInput {...field} placeholder="41.2" type="number" />
    {/snippet}
  </BField>
</div>
