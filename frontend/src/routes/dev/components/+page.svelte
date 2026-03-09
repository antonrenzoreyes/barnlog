<script lang="ts">
  import {
    BBadge,
    BButton,
    BCard,
    BChoose,
    BDate,
    BField,
    BFileUploader,
    BIcon,
    BInput,
  } from "$lib/ui";

  const speciesOptions = [
    { id: "showcase-choose-goat", label: "Goat", value: "goat" },
    { id: "showcase-choose-pig", label: "Pig", value: "pig" },
    { id: "showcase-choose-dog", label: "Dog", value: "dog" },
    { id: "showcase-choose-cat", label: "Cat", value: "cat" },
  ];

  const formState = $state({
    selectedSpecies: "pig",
    selectedSpeciesId: "showcase-choose-pig",
  });
</script>

<main class="app-shell app-shell-wide py-6" id="top">
  <section class="mb-5">
    <p class="text-xs font-semibold tracking-wide text-barn-muted uppercase">
      Dev Route
    </p>
    <h1 class="mt-1 text-2xl font-semibold">UI Components Showcase</h1>
    <p class="mt-1 text-sm text-barn-muted">
      Reference page for the current Barn UI primitives.
    </p>
    <div
      class="mt-3 rounded-barn-md border border-barn-line bg-barn-surface px-3 py-2 text-xs text-barn-muted"
    >
      Path: /dev/components
    </div>
  </section>

  <div class="grid gap-4 lg:grid-cols-[230px_minmax(0,1fr)] lg:items-start">
    <aside class="lg:sticky lg:top-4">
      <nav
        aria-label="Component table of contents"
        class="rounded-barn-md border border-barn-line bg-barn-surface p-3"
      >
        <p
          class="mb-3 text-xs font-semibold tracking-wide text-barn-muted uppercase"
        >
          Components
        </p>
        <ol class="grid gap-1 text-sm">
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bicon">BIcon</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bfield">BField</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bchoose">BChoose</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bdate">BDate</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bfileuploader">BFileUploader</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#binput">BInput</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bbutton">BButton</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bbadge">BBadge</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#bcard">BCard</a
            >
          </li>
          <li>
            <a
              class="block rounded-barn-sm px-2 py-1 text-barn-muted transition-colors hover:bg-barn-surface-muted hover:text-barn-text"
              href="#notes">Notes</a
            >
          </li>
        </ol>
      </nav>
    </aside>

    <div>
      <BCard id="bicon" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BIcon
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Name-based icon rendering using mapped app icon keys.
        </p>
        <div class="grid gap-2 sm:grid-cols-2">
          <div class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm">
            <div class="mb-2 text-xs text-barn-muted">inline</div>
            <div class="flex items-center gap-2">
              <BIcon name="search" size="inline" />
              <span>search</span>
            </div>
          </div>
          <div class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm">
            <div class="mb-2 text-xs text-barn-muted">button</div>
            <div class="flex items-center gap-2">
              <BIcon name="editAnimal" size="button" />
              <span>editAnimal</span>
            </div>
          </div>
          <div class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm">
            <div class="mb-2 text-xs text-barn-muted">fab</div>
            <div class="flex items-center gap-2">
              <BIcon name="addLogEvent" size="fab" />
              <span>addLogEvent</span>
            </div>
          </div>
          <div class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm">
            <div class="mb-2 text-xs text-barn-muted">display</div>
            <div class="flex items-center gap-2">
              <BIcon name="app" size="display" />
              <span>app</span>
            </div>
          </div>
        </div>
        <div
          class="mt-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
        >
          <div class="mb-1 text-xs text-barn-muted">a11y with title</div>
          <div class="flex items-center gap-2">
            <BIcon
              name="timelineTime"
              size="button"
              title="Timeline time icon"
            />
            <span>Icon with readable label</span>
          </div>
        </div>
      </BCard>

      <BCard id="bfield" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BField
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Wrapper for label, control and hint/error messaging around any form
          control.
        </p>
        <div class="grid gap-3">
          <BField
            label="Species"
            for="showcase-field-species"
            hint="Pick the closest match"
          >
            {#snippet children(field)}
              <select {...field} class="input-base">
                <option>Goat</option>
                <option>Pig</option>
                <option>Dog</option>
              </select>
            {/snippet}
          </BField>
          <BField
            error="Notes are required"
            for="showcase-field-notes"
            label="Notes"
            required={true}
          >
            {#snippet children(field)}
              <textarea
                {...field}
                class="input-base min-h-24 py-2"
                placeholder="Add context for this log event"
              ></textarea>
            {/snippet}
          </BField>
        </div>
      </BCard>

      <BCard id="bchoose" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BChoose
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Single-select tile chooser for option grids (prototype-aligned species
          picker).
        </p>
        <BField label="Species" required={true}>
          {#snippet children(field)}
            <BChoose
              {...field}
              name="showcase-species"
              onValueChange={(nextId, nextValue) => {
                formState.selectedSpeciesId = nextId;
                formState.selectedSpecies = nextValue;
              }}
              options={speciesOptions}
              selectedId={formState.selectedSpeciesId}
            />
          {/snippet}
        </BField>
        <p class="mt-3 text-sm text-barn-muted">
          Selected: {formState.selectedSpecies}
        </p>
      </BCard>

      <BCard id="bdate" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BDate
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Native date input wrapper for date-only fields.
        </p>
        <div class="grid gap-3">
          <BField
            label="Birth date"
            for="showcase-birth-date"
            hint="YYYY-MM-DD"
          >
            {#snippet children(field)}
              <BDate {...field} name="birthDate" value="2026-03-01" />
            {/snippet}
          </BField>
          <BField
            error="Date is required"
            for="showcase-check-date"
            label="Check date"
          >
            {#snippet children(field)}
              <BDate {...field} />
            {/snippet}
          </BField>
        </div>
      </BCard>

      <BCard id="bfileuploader" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BFileUploader
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Native file input wrapper for attachments and media uploads.
        </p>
        <div class="grid gap-3">
          <BField
            label="Attachments"
            for="showcase-attachments"
            hint="Upload one or more image files"
          >
            {#snippet children(field)}
              <BFileUploader
                {...field}
                accept="image/*"
                multiple={true}
                name="attachments"
              />
            {/snippet}
          </BField>
          <BField
            error="At least one file is required"
            for="showcase-proof"
            label="Proof"
          >
            {#snippet children(field)}
              <BFileUploader {...field} name="proof" />
            {/snippet}
          </BField>
        </div>
      </BCard>

      <BCard id="binput" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BInput
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Base text-like input primitive for common form fields.
        </p>
        <div class="grid gap-3">
          <BField
            label="Animal name"
            for="showcase-input-animal-name"
            hint="Use a short, readable name"
          >
            {#snippet children(field)}
              <BInput
                {...field}
                name="animalName"
                placeholder="Bella"
                type="text"
              />
            {/snippet}
          </BField>
          <BField
            error="Weight is required"
            for="showcase-input-weight"
            label="Weight"
          >
            {#snippet children(field)}
              <BInput {...field} placeholder="41.2" type="number" />
            {/snippet}
          </BField>
        </div>
      </BCard>

      <BCard id="bbutton" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BButton
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Supports variants, sizes, and optional leading/trailing icons.
        </p>
        <div class="grid gap-2">
          <BButton variant="primary" leadingIcon="addLogEvent">Primary</BButton>
          <BButton variant="secondary" trailingIcon="editAnimal"
            >Secondary</BButton
          >
          <BButton variant="danger">Danger</BButton>
        </div>
        <div class="mt-2 grid gap-2 sm:grid-cols-2">
          <BButton variant="primary" size="sm" leadingIcon="search"
            >Small + Icon</BButton
          >
          <BButton variant="secondary" disabled={true}>Disabled</BButton>
        </div>
        <div
          class="mt-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
        >
          Use `type="button"` unless the button should submit a form.
        </div>
      </BCard>

      <BCard id="bbadge" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BBadge
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Compact label/tag primitive for metadata and quick status text.
        </p>
        <div class="flex flex-wrap gap-2">
          <BBadge>Goat</BBadge>
          <BBadge>Medication Due</BBadge>
          <BBadge class="bg-barn-chip-bg text-barn-chip-text">Healthy</BBadge>
        </div>
        <div
          class="mt-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
        >
          <div class="mb-1 text-xs text-barn-muted">
            Long content wrap behavior
          </div>
          <BBadge>
            Needs follow-up checkup after feed and medication adjustments this
            afternoon
          </BBadge>
        </div>
      </BCard>

      <BCard id="bcard" class="mb-4">
        <div class="mb-3 flex items-center justify-between">
          <h2 class="text-sm font-semibold tracking-wide text-barn-muted">
            BCard
          </h2>
          <a class="text-xs text-barn-muted" href="#top">Back to top</a>
        </div>
        <p class="mb-3 text-sm text-barn-muted">
          Use snippets for header, body, and footer to structure content.
        </p>
        <div class="mb-3 grid gap-2 text-xs sm:grid-cols-2">
          <div
            class="rounded-barn-md border border-barn-line bg-barn-surface px-3 py-2"
          >
            <span class="font-semibold">Structured mode:</span> header + body + footer
            snippets
          </div>
          <div
            class="rounded-barn-md border border-barn-line bg-barn-surface px-3 py-2"
          >
            <span class="font-semibold">Simple mode:</span> pass children only
          </div>
        </div>
        <BCard>
          {#snippet header()}
            <div class="flex items-center justify-between">
              <div>
                <p class="text-base font-semibold">Bella (Goat)</p>
                <p class="text-sm text-barn-muted">ID: A-1024</p>
              </div>
              <BIcon name="editAnimal" size="button" />
            </div>
          {/snippet}

          {#snippet body()}
            <div class="space-y-3">
              <div class="grid gap-2 sm:grid-cols-2">
                <div
                  class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
                >
                  Last feed: 7:30 AM
                </div>
                <div
                  class="rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
                >
                  Last weight: 41.2 kg
                </div>
              </div>
              <p class="text-sm text-barn-text">
                Card body can render either the body snippet or default
                children.
              </p>
            </div>
          {/snippet}

          {#snippet footer()}
            <div class="grid grid-cols-2 gap-2">
              <BButton variant="secondary" class="flex-1">Cancel</BButton>
              <BButton
                variant="primary"
                class="flex-1"
                leadingIcon="addLogEvent">Save</BButton
              >
            </div>
          {/snippet}
        </BCard>
        <div class="mt-3">
          <BCard>
            <p class="text-sm">Simple children-only card content.</p>
          </BCard>
        </div>
      </BCard>

      <BCard id="notes">
        <div class="mb-2 flex items-center gap-2">
          <BIcon name="timelineTime" size="inline" />
          <h2
            class="text-sm font-semibold tracking-wide text-barn-muted uppercase"
          >
            Notes
          </h2>
        </div>
        <ul class="grid gap-1 text-sm text-barn-muted">
          <li>
            Use this page for quick visual verification of component states.
          </li>
          <li>When adding a component, add a section and TOC link.</li>
        </ul>
      </BCard>
    </div>
  </div>
</main>
