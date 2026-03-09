<script lang="ts">
  import { icons, type IconNode } from "lucide";
  import { type AppIconKey, iconMap } from "$lib/ui/icons";

  type IconSize = "inline" | "button" | "fab" | "display";
  /**
   * BIcon renders a mapped Lucide icon by app-level name.
   * Defaults: size='inline'.
   */
  interface BIconProps {
    name: AppIconKey;
    size?: IconSize;
    title?: string;
    class?: string;
  }

  const {
    name,
    size = "inline",
    title,
    class: className = "",
  }: BIconProps = $props();

  const iconKey = $derived(iconMap[name]);
  const iconNode = $derived(
    (icons as Record<string, IconNode | undefined>)[iconKey],
  );

  const iconStrokeWidth = 1.75;
  const sizeClass = $derived(`icon-${size}`);
</script>

{#if iconNode}
  <svg
    class={`${sizeClass} ${className}`}
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width={String(iconStrokeWidth)}
    stroke-linecap="round"
    stroke-linejoin="round"
    role={title ? "img" : undefined}
    aria-hidden={title ? undefined : "true"}
    aria-label={title}
  >
    {#if title}
      <title>{title}</title>
    {/if}
    {#each iconNode as [tag, attrs], index (`${tag}-${String(index)}`)}
      <svelte:element this={tag} {...attrs} />
    {/each}
  </svg>
{/if}

<style>
  .icon-inline {
    height: 1rem;
    width: 1rem;
  }

  .icon-button {
    height: 18px;
    width: 18px;
  }

  .icon-fab {
    height: 1.25rem;
    width: 1.25rem;
  }

  .icon-display {
    height: 1.5rem;
    width: 1.5rem;
  }
</style>
