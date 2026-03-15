<script lang="ts">
  import type { Snippet } from "svelte";
  import type { HTMLButtonAttributes } from "svelte/elements";

  interface Props {
    "aria-label"?: string;
    children: Snippet;
    class?: string;
    disabled?: boolean;
    id?: string;
    onclick?: HTMLButtonAttributes["onclick"];
    size?: "fab" | "lg" | "md" | "sm";
    tone?: "danger" | "primary" | "secondary";
    type?: HTMLButtonAttributes["type"];
  }

  const props: Props = $props();
</script>

<button
  aria-label={props["aria-label"]}
  class={[
    "inline-flex items-center justify-center gap-2 border font-semibold transition-[background-color,border-color,color,box-shadow,transform] duration-150 ease-out focus-visible:outline-none focus-visible:shadow-[var(--ui-shadow-focus)] active:translate-y-px disabled:cursor-not-allowed disabled:opacity-55 disabled:active:translate-y-0",
    props.size === "sm" && "h-9 rounded-lg px-3 text-sm",
    (!props.size || props.size === "md") && "h-11 rounded-xl px-4 text-sm",
    props.size === "lg" && "h-12 rounded-xl px-5 text-base",
    props.size === "fab" &&
      "h-14 w-14 rounded-full p-0 shadow-[var(--ui-shadow-card)]",
    props.tone === "primary" &&
      "border-(--ui-color-primary-strong) bg-(--ui-color-primary) text-white hover:bg-(--ui-color-primary-strong)",
    props.tone === "danger" &&
      "border-(--ui-color-border-danger) bg-[color:color-mix(in_oklab,var(--ui-color-border-danger)_12%,var(--ui-color-surface))] text-(--ui-color-text-danger) hover:bg-[color:color-mix(in_oklab,var(--ui-color-border-danger)_18%,var(--ui-color-surface))]",
    (!props.tone || props.tone === "secondary") &&
      "border-(--ui-color-border-strong) bg-(--ui-color-surface) text-(--ui-color-text) hover:border-(--ui-color-primary) hover:bg-(--ui-color-primary-soft)",
    props.class,
  ]}
  disabled={props.disabled}
  id={props.id}
  onclick={props.onclick}
  type={props.type ?? "button"}
>
  {@render props.children()}
</button>
