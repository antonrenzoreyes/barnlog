<script lang="ts">
  import type { HTMLFieldsetAttributes } from "svelte/elements";
  /**
   * BChoose renders a single-select tile group using radio inputs.
   * Use `selectedId` + `onValueChange` to track the selected option.
   */
  interface BChooseOption {
    value: string;
    label: string;
    id: string;
    disabled?: boolean;
  }

  interface BChooseProps extends Omit<
    HTMLFieldsetAttributes,
    "children" | "class"
  > {
    legend?: string;
    name: string;
    options: BChooseOption[];
    selectedId?: string;
    value?: string;
    onValueChange?: (optionId: string, value: string) => void;
    required?: boolean;
    disabled?: boolean;
    class?: string;
  }

  const {
    legend,
    name,
    options,
    selectedId,
    value,
    onValueChange,
    required = false,
    disabled = false,
    class: className = "",
    ...rest
  }: BChooseProps = $props();
  const chooserDomIdPrefix = $props.id();

  const resolvedSelectedId = $derived.by(() => {
    if (typeof selectedId === "string") {
      return selectedId;
    }

    if (typeof value !== "string") {
      return "";
    }

    const matchedOption = options.find((option) => option.value === value);
    if (matchedOption) {
      return matchedOption.id;
    }

    return "";
  });
</script>

<fieldset {...rest} class={`b-choose ${className}`} {disabled}>
  {#if legend}
    <legend class="b-choose-legend">{legend}</legend>
  {/if}
  {#each options as option (option.id)}
    {@const optionDomId = `${chooserDomIdPrefix}-${option.id}`}
    <label
      class="b-choose-option"
      class:b-choose-option-selected={resolvedSelectedId === option.id}
      class:b-choose-option-disabled={disabled || option.disabled}
      for={optionDomId}
    >
      <input
        id={optionDomId}
        checked={resolvedSelectedId === option.id}
        class="b-choose-input"
        disabled={disabled || option.disabled}
        {name}
        onchange={() => {
          if (onValueChange) {
            onValueChange(option.id, option.value);
          }
        }}
        {required}
        type="radio"
        value={option.id}
        data-value={option.value}
      />
      <span>{option.label}</span>
    </label>
  {/each}
</fieldset>

<style>
  .b-choose {
    margin: 0;
    padding: 0;
    min-width: 0;
    border: 0;
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 0.625rem;
  }

  .b-choose-legend {
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

  .b-choose-option {
    position: relative;
    display: inline-flex;
    min-height: 3.75rem;
    cursor: pointer;
    align-items: center;
    justify-content: center;
    border-radius: var(--radius-barn-md);
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    padding-inline: 0.75rem;
    text-align: center;
    font-size: 1.125rem;
    color: var(--color-barn-text);
    transition:
      background-color 150ms ease,
      border-color 150ms ease,
      color 150ms ease;
  }

  .b-choose-option-selected {
    border-color: var(--color-barn-primary);
    background: var(--color-barn-primary);
    color: rgb(245 245 244);
  }

  .b-choose-option-disabled {
    cursor: not-allowed;
    opacity: 0.6;
  }

  .b-choose-input {
    position: absolute;
    inset: 0;
    opacity: 0;
    cursor: inherit;
  }

  .b-choose-option:has(.b-choose-input:focus-visible) {
    outline: 2px solid color-mix(in srgb, var(--color-barn-primary) 45%, white);
    outline-offset: 1px;
  }
</style>
