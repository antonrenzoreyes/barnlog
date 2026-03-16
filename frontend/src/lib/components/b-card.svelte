<script lang="ts">
  import type { Snippet } from "svelte";

  interface Props {
    "aria-label"?: string;
    as?: "article" | "div" | "header" | "section";
    children: Snippet;
    class?: string;
    id?: string;
    padding?: "lg" | "md" | "none";
    tone?: "panel" | "surface";
  }

  const props: Props = $props();
</script>

<svelte:element
  this={props.as ?? "div"}
  aria-label={props["aria-label"]}
  class={[
    `rounded-3xl border border-(--ui-color-border) shadow-(--ui-shadow-card)`,
    (!props.tone || props.tone === "surface") && "bg-(--ui-color-surface)",
    props.tone === "panel" && "bg-(--ui-color-panel)",
    (!props.padding || props.padding === "md") && "p-4",
    props.padding === "lg" && "p-5",
    props.padding === "none" && "p-0",
    props.class,
  ]}
  id={props.id}
>
  {@render props.children()}
</svelte:element>
