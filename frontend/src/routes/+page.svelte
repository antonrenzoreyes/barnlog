<script lang="ts">
	import { BBadge, BButton, BCard, BIcon } from '$lib/ui';
</script>

<main class="app-shell py-6">
	<section class="mb-4">
		<h1 class="text-2xl font-semibold">Components Showcase</h1>
		<p class="mt-1 text-sm text-barn-muted">
			Table of contents and examples for each UI component.
		</p>
	</section>

	<section class="app-card mb-4 p-4" id="toc">
		<h2 class="mb-3 text-sm font-semibold tracking-wide text-barn-muted uppercase">
			Table of Contents
		</h2>
		<nav aria-label="Component table of contents">
			<ol class="grid gap-2 text-sm sm:grid-cols-2">
				<li><a class="chip" href="#bicon">BIcon</a></li>
				<li><a class="chip" href="#bbutton">BButton</a></li>
				<li><a class="chip" href="#bbadge">BBadge</a></li>
				<li><a class="chip" href="#bcard">BCard</a></li>
			</ol>
		</nav>
	</section>

	<section class="app-card mb-4 p-4" id="bicon">
		<div class="mb-3 flex items-center justify-between">
			<h2 class="text-sm font-semibold tracking-wide text-barn-muted uppercase">BIcon</h2>
			<a class="text-xs text-barn-muted" href="#toc">Back to top</a>
		</div>
		<div class="flex flex-wrap items-center gap-3">
			<span
				class="inline-flex items-center gap-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
			>
				<BIcon name="app" size="button" />
				button
			</span>
			<span
				class="inline-flex items-center gap-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
			>
				<BIcon name="search" size="inline" />
				inline
			</span>
			<span
				class="inline-flex items-center gap-2 rounded-barn-md bg-barn-surface-muted px-3 py-2 text-sm"
			>
				<BIcon name="addLogEvent" size="fab" />
				fab
			</span>
		</div>
	</section>

	<section class="app-card mb-4 p-4" id="bbutton">
		<div class="mb-3 flex items-center justify-between">
			<h2 class="text-sm font-semibold tracking-wide text-barn-muted uppercase">BButton</h2>
			<a class="text-xs text-barn-muted" href="#toc">Back to top</a>
		</div>
		<div class="grid gap-2 sm:grid-cols-2">
			<BButton variant="primary" leadingIcon="addLogEvent">Primary</BButton>
			<BButton variant="secondary" trailingIcon="editAnimal">Secondary</BButton>
			<BButton variant="ghost">Ghost</BButton>
			<BButton variant="danger">Danger</BButton>
			<BButton variant="primary" size="sm">Small</BButton>
			<BButton variant="secondary" disabled={true}>Disabled</BButton>
		</div>
	</section>

	<section class="app-card mb-4 p-4" id="bbadge">
		<div class="mb-3 flex items-center justify-between">
			<h2 class="text-sm font-semibold tracking-wide text-barn-muted uppercase">BBadge</h2>
			<a class="text-xs text-barn-muted" href="#toc">Back to top</a>
		</div>
		<div class="flex flex-wrap gap-2">
			<BBadge>Default Badge</BBadge>
			<BBadge>Longer Badge Label</BBadge>
		</div>
	</section>

	<section class="app-card p-4" id="bcard">
		<div class="mb-3 flex items-center justify-between">
			<h2 class="text-sm font-semibold tracking-wide text-barn-muted uppercase">BCard</h2>
			<a class="text-xs text-barn-muted" href="#toc">Back to top</a>
		</div>
		<BCard>
			{#snippet header()}
				<div class="flex items-center justify-between">
					<div>
						<p class="text-base font-semibold">Card Header</p>
						<p class="text-sm text-barn-muted">Rendered via `header` snippet</p>
					</div>
					<BIcon name="editAnimal" size="button" />
				</div>
			{/snippet}

			{#snippet body()}
				<div class="space-y-3">
					<div class="flex items-center gap-2">
						<BBadge>Inside Body</BBadge>
						<span class="text-xs text-barn-muted">Rendered via `body` snippet</span>
					</div>
					<p class="text-sm text-barn-text">
						Use BCard as a single component with optional named snippets.
					</p>
				</div>
			{/snippet}

			{#snippet footer()}
				<div class="flex gap-2">
					<BButton variant="secondary" class="flex-1">Cancel</BButton>
					<BButton variant="primary" class="flex-1" leadingIcon="addLogEvent">Save</BButton>
				</div>
			{/snippet}
		</BCard>
	</section>
</main>
