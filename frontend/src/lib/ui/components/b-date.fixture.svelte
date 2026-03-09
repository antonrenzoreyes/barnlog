<script lang="ts">
  import BDate from "./b-date.svelte";
  import BField from "./b-field.svelte";

  const formState = $state({
    birthDate: "2026-03-01",
  });
</script>

<div>
  <BField label="Birth date" for="fixture-birth-date" hint="YYYY-MM-DD">
    {#snippet children(field)}
      <BDate {...field} name="birthDate" bind:value={formState.birthDate} />
    {/snippet}
  </BField>
  <p>Birth date value: {formState.birthDate}</p>

  <BField
    class="mt-4"
    error="Date is required"
    for="fixture-check-date"
    label="Check date"
  >
    {#snippet children(field)}
      <BDate {...field} />
    {/snippet}
  </BField>
</div>
