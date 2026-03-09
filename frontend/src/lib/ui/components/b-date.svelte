<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  /**
   * BDate is the base date input primitive.
   * It wraps a native date input.
   */
  interface BDateProps extends Omit<
    HTMLInputAttributes,
    "children" | "type" | "value"
  > {
    class?: string;
    value?: string;
  }

  // oxlint-disable prefer-const -- bindable props must remain writable.
  let {
    class: className = "",
    value = $bindable(),
    ...rest
  }: BDateProps = $props();
  // oxlint-enable prefer-const
</script>

<input
  {...rest}
  bind:value
  class={`b-date b-field-control-input ${className}`}
  type="date"
/>

<style>
  .b-date {
    min-height: 3rem;
    width: 100%;
    border-radius: var(--radius-barn-md);
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    padding-inline: 0.75rem;
    font-size: 1rem;
    color: var(--color-barn-text);
  }
</style>
