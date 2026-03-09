<script lang="ts">
  import type { Snippet } from "svelte";
  import type { HTMLButtonAttributes } from "svelte/elements";
  import BIcon from "./b-icon.svelte";
  import type { AppIconKey } from "$lib/ui/icons";

  type Variant = "primary" | "secondary" | "danger";
  type Size = "sm" | "md";
  type ButtonType = "button" | "submit" | "reset";
  /**
   * BButton is the base actionable control for forms and UI actions.
   * Defaults: variant='primary', size='md', type='button', disabled=false.
   */
  interface BButtonProps extends Omit<HTMLButtonAttributes, "children"> {
    variant?: Variant;
    size?: Size;
    type?: ButtonType;
    disabled?: boolean;
    leadingIcon?: AppIconKey;
    trailingIcon?: AppIconKey;
    children?: Snippet;
  }

  const {
    variant = "primary",
    size = "md",
    type = "button",
    disabled = false,
    leadingIcon,
    trailingIcon,
    class: className = "",
    children,
    ...rest
  }: BButtonProps = $props();

  const variantClass = $derived(`btn-${variant}`);
  const sizeClass = $derived(`btn-${size}`);
  const disabledClass = $derived.by(() => {
    if (disabled) {
      return "btn-disabled";
    }
    return "";
  });
</script>

<button
  {...rest}
  {type}
  {disabled}
  class={`b-button ${variantClass} ${sizeClass} ${disabledClass} ${className}`}
>
  {#if leadingIcon}
    <BIcon name={leadingIcon} size="button" />
  {/if}

  {#if children}
    {@render children()}
  {/if}

  {#if trailingIcon}
    <BIcon name={trailingIcon} size="button" />
  {/if}
</button>

<style>
  .b-button {
    display: inline-flex;
    gap: 0.5rem;
    align-items: center;
    justify-content: center;
    border-radius: 0.875rem;
    font-weight: 600;
    transition:
      background-color 150ms ease,
      color 150ms ease,
      opacity 150ms ease;
  }

  .btn-sm {
    min-height: 2.25rem;
    padding-inline: 0.75rem;
    font-size: 0.875rem;
  }

  .btn-md {
    min-height: 2.75rem;
    padding-inline: 1rem;
    font-size: 0.94rem;
  }

  .btn-primary {
    background: var(--color-barn-primary);
    color: rgb(245 245 244);
  }

  .btn-primary:not(:disabled):hover,
  .btn-primary:not(:disabled):active {
    background: var(--color-barn-primary-pressed);
  }

  .btn-secondary {
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    color: var(--color-barn-text);
  }

  .btn-secondary:not(:disabled):hover,
  .btn-secondary:not(:disabled):active {
    background: var(--color-barn-surface-muted);
  }

  .btn-danger {
    background: var(--color-barn-danger);
    color: rgb(245 245 244);
  }

  .btn-danger:not(:disabled):hover,
  .btn-danger:not(:disabled):active {
    opacity: 0.9;
  }

  .btn-disabled {
    cursor: not-allowed;
    opacity: 0.5;
  }
</style>
