<script lang="ts">
  import {
    getLocalTimeZone,
    today,
    type DateValue,
  } from "@internationalized/date";
  import { DatePicker } from "bits-ui";
  import {
    CalendarDays,
    ChevronDown,
    ChevronLeft,
    ChevronRight,
  } from "lucide-svelte";

  interface Props {
    "aria-describedby"?: string;
    disabled?: boolean;
    id?: string;
    invalid?: boolean;
    name?: string;
    placeholder?: string;
    required?: boolean;
    value?: DateValue;
  }

  /* oxlint-disable eslint/prefer-const */
  let {
    "aria-describedby": ariaDescribedBy,
    disabled,
    id,
    invalid,
    name,
    placeholder,
    required,
    value = $bindable(),
  }: Props = $props();
  /* oxlint-enable eslint/prefer-const */

  const formatter = new Intl.DateTimeFormat(undefined, {
    day: "numeric",
    month: "short",
    year: "numeric",
  });

  let isOpen = $state(false);
  const ariaInvalid = $derived(Boolean(invalid));
  const formValue = $derived(value?.toString() ?? "");

  const displayText = $derived.by(() => {
    if (value) {
      return formatter.format(value.toDate(getLocalTimeZone()));
    }

    if (placeholder) {
      return placeholder;
    }

    return "Select date";
  });

  const displayTextClass = $derived.by(() => {
    if (value) {
      return "text-(--ui-color-text)";
    }

    return "font-normal text-(--ui-color-placeholder)";
  });

  const setToday = (): void => {
    value = today(getLocalTimeZone());
    isOpen = false;
  };
  const clearDate = (): void => {
    value = undefined;
    isOpen = false;
  };
</script>

<DatePicker.Root
  bind:open={isOpen}
  bind:value
  {disabled}
  {required}
  weekdayFormat="short"
>
  <DatePicker.Trigger
    aria-describedby={ariaDescribedBy}
    aria-invalid={ariaInvalid}
    class="group flex h-14 w-full items-center justify-between gap-3 rounded-xl border border-(--ui-color-border) bg-(--ui-color-surface) px-3.5 text-left text-base font-medium shadow-[var(--ui-shadow-input)] outline-none transition-[border-color,box-shadow,background-color] duration-200 ease-out hover:border-(--ui-color-border-strong) hover:bg-[color:color-mix(in_oklab,var(--ui-color-surface)_90%,var(--ui-color-panel))] focus:border-(--ui-color-primary) focus:ring-0 focus:shadow-[var(--ui-shadow-focus)] aria-invalid:border-(--ui-color-border-danger) aria-invalid:focus:border-(--ui-color-border-danger) aria-invalid:focus:shadow-[var(--ui-shadow-focus-danger)] disabled:cursor-not-allowed disabled:border-(--ui-color-border) disabled:bg-[color:color-mix(in_oklab,var(--ui-color-surface)_40%,var(--ui-color-panel))] disabled:text-(--ui-color-text-soft)"
    {id}
  >
    <span class="flex min-w-0 items-center gap-2.5">
      <span
        aria-hidden="true"
        class="grid h-8 w-8 shrink-0 place-items-center rounded-lg border border-(--ui-color-border) bg-(--ui-color-panel) text-(--ui-color-primary-strong) transition group-hover:border-(--ui-color-border-strong)"
      >
        <CalendarDays size={15} strokeWidth={2.25} />
      </span>
      <span class={`truncate ${displayTextClass}`}>{displayText}</span>
    </span>

    <span
      aria-hidden="true"
      class="grid h-7 w-7 shrink-0 place-items-center text-(--ui-color-text-soft) transition-transform duration-200 ease-out group-data-[state=open]:rotate-180"
    >
      <ChevronDown size={15} strokeWidth={2.3} />
    </span>
  </DatePicker.Trigger>

  {#if name}
    <input
      class="sr-only"
      {disabled}
      {name}
      required={Boolean(required)}
      type="date"
      value={formValue}
    />
  {/if}

  <DatePicker.Portal>
    <DatePicker.Content
      class="z-50 w-[min(22rem,calc(100vw-2rem))] overflow-hidden rounded-2xl border border-(--ui-color-border) bg-(--ui-color-surface) p-0 shadow-[var(--ui-shadow-card)]"
      sideOffset={8}
    >
      <DatePicker.Calendar class="space-y-2 p-3">
        {#snippet children({ months, weekdays })}
          {#each months as month (month.value.toString())}
            <DatePicker.Grid class="w-full border-collapse">
              <DatePicker.Header
                class="mb-1.5 flex items-center justify-between gap-2"
              >
                <DatePicker.PrevButton
                  class="grid h-8 w-8 place-items-center rounded-lg border border-(--ui-color-border) text-(--ui-color-text) transition duration-150 hover:border-(--ui-color-border-strong) hover:bg-(--ui-color-panel) focus:shadow-[var(--ui-shadow-focus)] data-[disabled]:cursor-not-allowed data-[disabled]:opacity-40"
                >
                  <ChevronLeft size={16} strokeWidth={2.25} />
                </DatePicker.PrevButton>

                <div class="flex min-w-0 items-center gap-1.5">
                  <DatePicker.MonthSelect
                    aria-label="Month"
                    class="h-9 w-20 rounded-md border border-(--ui-color-border) bg-(--ui-color-surface) py-1 pl-2 pr-7 text-sm font-medium text-(--ui-color-text) outline-none transition hover:border-(--ui-color-border-strong) focus:shadow-[var(--ui-shadow-focus)]"
                    monthFormat="short"
                  />
                  <DatePicker.YearSelect
                    aria-label="Year"
                    class="h-9 w-[4.75rem] rounded-md border border-(--ui-color-border) bg-(--ui-color-surface) py-1 pl-2 pr-7 text-sm font-medium text-(--ui-color-text) outline-none transition hover:border-(--ui-color-border-strong) focus:shadow-[var(--ui-shadow-focus)]"
                  />
                </div>

                <DatePicker.NextButton
                  class="grid h-8 w-8 place-items-center rounded-lg border border-(--ui-color-border) text-(--ui-color-text) transition duration-150 hover:border-(--ui-color-border-strong) hover:bg-(--ui-color-panel) focus:shadow-[var(--ui-shadow-focus)] data-[disabled]:cursor-not-allowed data-[disabled]:opacity-40"
                >
                  <ChevronRight size={16} strokeWidth={2.25} />
                </DatePicker.NextButton>
              </DatePicker.Header>

              <DatePicker.GridHead>
                <DatePicker.GridRow class="grid grid-cols-7 gap-1">
                  {#each weekdays as weekday, weekdayIndex (`${weekday}-${weekdayIndex}`)}
                    <DatePicker.HeadCell
                      class="grid h-7 place-items-center text-xs font-semibold uppercase tracking-[0.05em] text-(--ui-color-text-soft)"
                    >
                      {weekday}
                    </DatePicker.HeadCell>
                  {/each}
                </DatePicker.GridRow>
              </DatePicker.GridHead>

              <DatePicker.GridBody class="grid gap-1">
                {#each month.weeks as week, weekIndex (`${month.value.toString()}-${weekIndex}`)}
                  <DatePicker.GridRow class="grid grid-cols-7 gap-1">
                    {#each week as date (date.toString())}
                      <DatePicker.Cell {date} month={month.value}>
                        <DatePicker.Day
                          class="relative grid h-9 w-full place-items-center rounded-lg border border-transparent text-sm font-semibold text-(--ui-color-text) outline-none transition-[background-color,color,border-color,box-shadow] duration-150 hover:border-(--ui-color-border) hover:bg-(--ui-color-panel) data-[outside-month]:text-(--ui-color-text-soft) data-[today]:border-(--ui-color-primary) data-[today]:ring-1 data-[today]:ring-(--ui-color-primary) data-[today]:after:absolute data-[today]:after:bottom-1 data-[today]:after:h-1 data-[today]:after:w-1 data-[today]:after:rounded-full data-[today]:after:bg-current data-[selected]:border-(--ui-color-primary-strong) data-[selected]:bg-(--ui-color-primary) data-[selected]:text-white data-[selected]:shadow-[0_6px_14px_color-mix(in_oklab,var(--ui-color-primary)_38%,transparent)] data-[selected]:after:hidden data-[disabled]:cursor-not-allowed data-[disabled]:opacity-35 data-[disabled]:hover:border-transparent data-[disabled]:hover:bg-transparent focus-visible:shadow-[var(--ui-shadow-focus)]"
                        />
                      </DatePicker.Cell>
                    {/each}
                  </DatePicker.GridRow>
                {/each}
              </DatePicker.GridBody>
            </DatePicker.Grid>
          {/each}
        {/snippet}
      </DatePicker.Calendar>

      <div
        class="flex items-center gap-2 border-t border-(--ui-color-border) p-2"
      >
        <button
          class="h-9 flex-1 rounded-lg border border-(--ui-color-primary-strong) bg-(--ui-color-primary) px-3 text-sm font-semibold text-white transition hover:bg-(--ui-color-primary-strong) focus:shadow-[var(--ui-shadow-focus)]"
          onclick={setToday}
          type="button"
        >
          Today
        </button>

        <button
          class="h-9 flex-1 rounded-lg border border-transparent bg-transparent px-3 text-sm font-medium text-(--ui-color-text-soft) transition hover:bg-(--ui-color-panel) hover:text-(--ui-color-text) focus:shadow-[var(--ui-shadow-focus)] disabled:cursor-not-allowed disabled:text-(--ui-color-text-soft) disabled:opacity-45"
          disabled={!value}
          onclick={clearDate}
          type="button"
        >
          Clear
        </button>
      </div>
    </DatePicker.Content>
  </DatePicker.Portal>
</DatePicker.Root>
