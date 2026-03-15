<script lang="ts">
  import { Label } from "bits-ui";

  interface Props {
    forId: string;
    required?: boolean;
    text: string;
  }

  const props: Props = $props();
</script>

<Label.Root
  class="block text-[0.95rem] font-semibold tracking-[0.01em] text-(--ui-color-text)"
  for={props.forId}
>
  {props.text}
  {#if props.required}
    <span aria-hidden="true" class="ml-1 text-(--ui-color-required)">*</span>
    <span class="sr-only">(required)</span>
  {/if}
</Label.Root>
