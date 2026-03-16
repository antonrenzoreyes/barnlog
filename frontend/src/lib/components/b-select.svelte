<script lang="ts">
  import { Select } from "bits-ui";
  import { Check, ChevronDown } from "lucide-svelte";

  interface SelectOption {
    description?: string;
    disabled?: boolean;
    label: string;
    value: string;
  }

  interface Props {
    "aria-describedby"?: string;
    disabled?: boolean;
    id?: string;
    invalid?: boolean;
    name?: string;
    options: SelectOption[];
    placeholder?: string;
    required?: boolean;
    value?: string;
  }

  /* oxlint-disable eslint/prefer-const */
  let {
    "aria-describedby": ariaDescribedBy,
    disabled,
    id,
    invalid,
    name,
    options,
    placeholder,
    required,
    value = $bindable(""),
  }: Props = $props();
  /* oxlint-enable eslint/prefer-const */

  const ariaInvalid = $derived(Boolean(invalid));
  const selectedLabel = $derived.by(() => {
    const selectedOption = options.find((option) => option.value === value);

    return selectedOption?.label;
  });
  const displayText = $derived(selectedLabel ?? placeholder ?? "Select");
</script>

<Select.Root
  type="single"
  bind:value
  {disabled}
  items={options}
  {name}
  {required}
>
  <Select.Trigger
    aria-describedby={ariaDescribedBy}
    aria-invalid={ariaInvalid}
    class="
      group flex h-14 w-full items-center justify-between gap-3 rounded-xl
      border border-(--ui-color-border) bg-(--ui-color-surface) px-3.5 text-left
      text-base font-medium shadow-(--ui-shadow-input)
      transition-[border-color,box-shadow,background-color] duration-200
      ease-out outline-none
      hover:border-(--ui-color-border-strong)
      hover:bg-[color-mix(in_oklab,var(--ui-color-surface)_90%,var(--ui-color-panel))]
      focus:border-(--ui-color-primary) focus:shadow-(--ui-shadow-focus)
      focus:ring-0
      disabled:cursor-not-allowed disabled:border-(--ui-color-border)
      disabled:bg-[color-mix(in_oklab,var(--ui-color-surface)_40%,var(--ui-color-panel))]
      disabled:text-(--ui-color-text-soft)
      aria-invalid:border-(--ui-color-border-danger)
      aria-invalid:focus:border-(--ui-color-border-danger)
      aria-invalid:focus:shadow-(--ui-shadow-focus-danger)
      data-[state=open]:border-(--ui-color-primary)
      data-[state=open]:bg-(--ui-color-panel)
      data-[state=open]:shadow-(--ui-shadow-focus)
    "
    {id}
  >
    <span
      class={[
        "min-w-0 truncate",
        selectedLabel
          ? "text-(--ui-color-text)"
          : "font-normal text-(--ui-color-placeholder)",
      ]}
    >
      {displayText}
    </span>

    <span
      aria-hidden="true"
      class="
        grid size-7 shrink-0 place-items-center text-(--ui-color-text-soft)
        transition-transform duration-200 ease-out
        group-data-[state=open]:rotate-180
      "
    >
      <ChevronDown size={15} strokeWidth={2.3} />
    </span>
  </Select.Trigger>

  <Select.Portal>
    <Select.Content
      class="
        z-50 w-[min(22rem,calc(100vw-2rem))] overflow-hidden rounded-2xl border
        border-(--ui-color-border) bg-(--ui-color-surface) p-1.5
        shadow-(--ui-shadow-card)
      "
      sideOffset={8}
    >
      <Select.Viewport class="max-h-72 space-y-1 overflow-y-auto p-0.5">
        {#if options.length === 0}
          <div
            class="
              flex min-h-11 items-center rounded-lg px-3 text-sm
              text-(--ui-color-text-soft)
            "
          >
            No options available
          </div>
        {:else}
          {#each options as option (option.value)}
            <Select.Item
              disabled={option.disabled}
              label={option.label}
              value={option.value}
            >
              {#snippet children({ highlighted, selected })}
                <div
                  class={[
                    `
                      flex min-h-11 items-center justify-between gap-3
                      rounded-lg px-3 py-2 text-sm font-medium transition
                    `,
                    highlighted && "bg-(--ui-color-panel)",
                    selected &&
                      `
                        bg-(--ui-color-primary-soft)
                        text-(--ui-color-primary-strong)
                      `,
                    option.disabled && "text-(--ui-color-text-soft)",
                  ]}
                >
                  <div class="min-w-0">
                    <span class="block truncate" title={option.label}>
                      {option.label}
                    </span>
                    {#if option.description}
                      <span
                        class="
                          block text-xs font-normal text-(--ui-color-text-soft)
                        "
                      >
                        {option.description}
                      </span>
                    {/if}
                  </div>

                  <span
                    class={[
                      `
                        grid size-5 shrink-0 place-items-center
                        text-(--ui-color-primary-strong) transition-opacity
                      `,
                      selected ? "opacity-100" : "opacity-0",
                    ]}
                  >
                    <Check aria-hidden="true" size={14} strokeWidth={2.5} />
                  </span>
                </div>
              {/snippet}
            </Select.Item>
          {/each}
        {/if}
      </Select.Viewport>
    </Select.Content>
  </Select.Portal>
</Select.Root>
