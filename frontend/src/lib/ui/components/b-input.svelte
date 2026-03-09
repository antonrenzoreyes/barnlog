<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";
  /**
   * BInput is the base text-like input primitive for forms.
   * Supports native input attributes.
   */
  type TextInputType =
    | "text"
    | "email"
    | "password"
    | "search"
    | "tel"
    | "url"
    | "number";

  interface BInputProps extends Omit<
    HTMLInputAttributes,
    "children" | "value" | "type"
  > {
    class?: string;
    type?: TextInputType;
    value?: string;
  }

  // oxlint-disable prefer-const -- bindable props must remain writable.
  let {
    class: className = "",
    type = "text",
    value = $bindable(""),
    ...rest
  }: BInputProps = $props();
  // oxlint-enable prefer-const
</script>

<input
  {...rest}
  bind:value
  class={`b-input b-field-control-input ${className}`}
  {type}
/>

<style>
  .b-input {
    min-height: 3rem;
    width: 100%;
    border-radius: var(--radius-barn-md);
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    padding-inline: 0.75rem;
    font-size: 1rem;
    color: var(--color-barn-text);
  }

  .b-input::placeholder {
    color: color-mix(in srgb, var(--color-barn-muted) 85%, transparent);
  }
</style>
