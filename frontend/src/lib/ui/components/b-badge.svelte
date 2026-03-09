<script lang="ts">
  import type { Snippet } from "svelte";
  import type { HTMLAttributes } from "svelte/elements";
  /**
   * BBadge shows compact metadata text (tags, status, labels).
   */
  interface BBadgeProps extends Omit<
    HTMLAttributes<HTMLSpanElement>,
    "children"
  > {
    children?: Snippet;
  }

  const { class: className = "", children, ...rest }: BBadgeProps = $props();
</script>

<span {...rest} class={`b-badge ${className}`}>
  {#if children}
    {@render children()}
  {/if}
</span>

<style>
  .b-badge {
    display: inline-flex;
    align-items: center;
    border-radius: 9999px;
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface-muted);
    padding: 0.25rem 0.625rem;
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--color-barn-muted);
  }
</style>
