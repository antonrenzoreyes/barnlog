<script lang="ts">
  import type { Snippet } from "svelte";
  import type { HTMLAttributes } from "svelte/elements";
  /**
   * BCard is a container primitive with optional header/body/footer snippets.
   * If `body` is omitted, `children` is rendered as the body content.
   */
  interface BCardProps extends Omit<HTMLAttributes<HTMLElement>, "children"> {
    children?: Snippet;
    header?: Snippet;
    body?: Snippet;
    footer?: Snippet;
  }

  const {
    class: className,
    children,
    header,
    body,
    footer,
    ...rest
  }: BCardProps = $props();
</script>

<section {...rest} class={["b-card", className]}>
  {#if header}
    <header class="border-b border-barn-line p-4">
      {@render header()}
    </header>
  {/if}

  {#if body}
    <div class="p-4">
      {@render body()}
    </div>
  {:else if children}
    <div class="p-4">
      {@render children()}
    </div>
  {/if}

  {#if footer}
    <footer class="border-t border-barn-line p-4">
      {@render footer()}
    </footer>
  {/if}
</section>

<style>
  .b-card {
    border-radius: 0.875rem;
    border: 1px solid var(--color-barn-line);
    background: var(--color-barn-surface);
    box-shadow: var(--shadow-barn-card);
  }
</style>
