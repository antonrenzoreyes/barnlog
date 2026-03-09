<script lang="ts">
  import type { Snippet } from "svelte";
  /**
   * BField standardizes label + control + hint/error messaging for form inputs.
   */
  interface BFieldControlAttributes {
    id: string;
    required?: boolean;
    "aria-required"?: "true";
    "aria-describedby"?: string;
    "aria-invalid"?: "true";
  }

  interface BFieldProps {
    label?: string;
    for?: string;
    hint?: string;
    error?: string;
    invalid?: boolean;
    required?: boolean;
    class?: string;
    children?: Snippet<[BFieldControlAttributes]>;
  }

  const {
    label,
    for: forId,
    hint,
    error,
    invalid = false,
    required = false,
    class: className = "",
    children,
  }: BFieldProps = $props();
  const uid = $props.id();
  const controlId = $derived(forId || `${uid}-control`);
  const messageId = $derived.by(() => {
    if (error || hint) {
      return `${controlId}-message`;
    }
    return "";
  });

  const hasErrorState = $derived.by(() => {
    if (error) {
      return true;
    }
    if (invalid) {
      return true;
    }
    return false;
  });
  const controlAttributes = $derived.by(() => {
    const nextAttributes: BFieldControlAttributes = { id: controlId };
    if (required) {
      nextAttributes.required = true;
      nextAttributes["aria-required"] = "true";
    }
    if (messageId) {
      nextAttributes["aria-describedby"] = messageId;
    }
    if (hasErrorState) {
      nextAttributes["aria-invalid"] = "true";
    }
    return nextAttributes;
  });
</script>

<div class={["b-field", hasErrorState && "b-field-invalid", className]}>
  {#if label}
    <label class="b-field-label" for={controlId}>
      {label}
      {#if required}
        <span class="b-field-required-indicator" aria-hidden="true"> *</span>
      {/if}
    </label>
  {/if}

  {#if children}
    <div class="b-field-control">
      {@render children(controlAttributes)}
    </div>
  {/if}

  {#if error}
    <p class="b-field-message b-field-error" aria-live="polite" id={messageId}>
      {error}
    </p>
  {:else if hint}
    <p class="b-field-message" id={messageId}>{hint}</p>
  {/if}
</div>

<style>
  .b-field {
    display: grid;
    gap: 0.5rem;
  }

  .b-field-control :global(.b-field-control-input:focus-visible) {
    outline: 2px solid color-mix(in srgb, var(--color-barn-primary) 45%, white);
    outline-offset: 1px;
  }

  .b-field-label {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--color-barn-text);
  }

  .b-field-required-indicator {
    color: var(--color-barn-danger);
  }

  .b-field-invalid .b-field-control :global(.b-field-control-input) {
    border-color: var(--color-barn-danger);
  }

  .b-field-message {
    font-size: 0.75rem;
    color: var(--color-barn-muted);
  }

  .b-field-error {
    color: var(--color-barn-danger);
  }
</style>
