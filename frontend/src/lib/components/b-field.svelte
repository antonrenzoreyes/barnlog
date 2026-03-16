<script lang="ts">
  import type { Snippet } from "svelte";
  import BLabel from "./b-label.svelte";

  interface FieldRenderArgs {
    id: string;
    describedBy?: string;
    invalid: boolean;
    required: boolean;
  }

  interface Props {
    children: Snippet<[FieldRenderArgs]>;
    error?: string;
    hint?: string;
    id: string;
    label: string;
    required?: boolean;
  }

  const props: Props = $props();
  const hint = $derived(props.hint);
  const error = $derived(props.error);

  const hintId = $derived.by(() => {
    if (hint) {
      return `${props.id}-hint`;
    }
  });

  const errorId = $derived.by(() => {
    if (error) {
      return `${props.id}-error`;
    }
  });

  const describedBy = $derived.by(() => {
    const joinedIds = [hintId, errorId].filter(Boolean).join(" ");

    if (joinedIds) {
      return joinedIds;
    }
  });

  const invalid = $derived(Boolean(error));
</script>

<div class="space-y-2.5">
  <BLabel forId={props.id} required={props.required} text={props.label} />

  {@render props.children({
    id: props.id,
    describedBy,
    invalid,
    required: Boolean(props.required),
  })}

  {#if hint}
    <p
      id={hintId}
      class="text-[0.95rem] leading-snug text-(--ui-color-text-muted)"
    >
      {hint}
    </p>
  {/if}

  {#if error}
    <p
      id={errorId}
      class="
        text-[0.95rem] leading-snug font-semibold text-(--ui-color-text-danger)
      "
      role="alert"
    >
      {error}
    </p>
  {/if}
</div>
