<script lang="ts">
  interface Props {
    tone?: "danger" | "neutral" | "primary";
    text: string;
  }

  const props: Props = $props();
</script>

<span
  class={[
    "inline-flex items-center rounded-full border px-2.5 py-1 text-xs font-semibold tracking-[0.02em]",
    props.tone === "primary" &&
      "border-(--ui-color-primary) bg-(--ui-color-primary-soft) text-(--ui-color-primary-strong)",
    props.tone === "danger" &&
      "border-(--ui-color-border-danger) bg-[color:color-mix(in_oklab,var(--ui-color-border-danger)_12%,var(--ui-color-surface))] text-(--ui-color-text-danger)",
    (!props.tone || props.tone === "neutral") &&
      "border-(--ui-color-border-strong) bg-(--ui-color-surface) text-(--ui-color-text-muted)",
  ]}
>
  {props.text}
</span>
