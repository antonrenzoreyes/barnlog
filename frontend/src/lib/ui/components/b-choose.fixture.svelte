<script lang="ts">
  import BChoose from "./b-choose.svelte";
  import BField from "./b-field.svelte";

  const speciesOptions = [
    { id: "fixture-species-goat", label: "Goat", value: "goat" },
    { id: "fixture-species-pig", label: "Pig", value: "pig" },
    {
      id: "fixture-species-goat-sheep-first",
      label: "Goat / Sheep",
      value: "goat / sheep",
    },
    {
      id: "fixture-species-goat-sheep-second",
      label: "Goat / Sheep",
      value: "goat / sheep",
    },
  ];

  const formState = $state({
    speciesId: "fixture-species-pig",
    speciesValue: "pig",
  });
</script>

<BField label="Species" required={true}>
  {#snippet children(field)}
    <BChoose
      {...field}
      legend="Species"
      name="fixture-species"
      onValueChange={(nextId, nextValue) => {
        formState.speciesId = nextId;
        formState.speciesValue = nextValue;
      }}
      options={speciesOptions}
      selectedId={formState.speciesId}
    />
  {/snippet}
</BField>

<p>Selected species: {formState.speciesValue}</p>
