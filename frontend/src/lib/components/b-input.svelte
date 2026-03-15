<script lang="ts">
  import type { HTMLInputAttributes } from "svelte/elements";

  interface Props {
    "aria-describedby"?: string;
    autocomplete?: HTMLInputAttributes["autocomplete"];
    disabled?: boolean;
    id?: string;
    invalid?: boolean;
    name?: string;
    placeholder?: string;
    required?: boolean;
    type?: HTMLInputAttributes["type"];
    value?: string;
  }

  /* oxlint-disable eslint/prefer-const */
  let {
    "aria-describedby": ariaDescribedBy,
    autocomplete,
    disabled,
    id,
    invalid,
    name,
    placeholder,
    required,
    type = "text",
    value = $bindable(""),
  }: Props = $props();
  /* oxlint-enable eslint/prefer-const */

  const ariaInvalid = $derived(Boolean(invalid));
</script>

<input
  bind:value
  aria-describedby={ariaDescribedBy}
  aria-invalid={ariaInvalid}
  {autocomplete}
  class="h-14 w-full rounded-xl border border-(--ui-color-border) bg-(--ui-color-surface) px-3.5 text-base font-medium text-(--ui-color-text) shadow-[var(--ui-shadow-input)] outline-none transition-[border-color,box-shadow,background-color] duration-200 ease-out placeholder:font-normal placeholder:text-(--ui-color-placeholder) hover:border-(--ui-color-border-strong) focus:border-(--ui-color-primary) focus:ring-0 focus:shadow-[var(--ui-shadow-focus)] aria-invalid:border-(--ui-color-border-danger) aria-invalid:focus:border-(--ui-color-border-danger) aria-invalid:focus:shadow-[var(--ui-shadow-focus-danger)] disabled:cursor-not-allowed disabled:border-(--ui-color-border) disabled:bg-[color:color-mix(in_oklab,var(--ui-color-surface)_40%,var(--ui-color-panel))] disabled:text-(--ui-color-text-soft)"
  {disabled}
  {id}
  {name}
  {placeholder}
  {required}
  {type}
/>
