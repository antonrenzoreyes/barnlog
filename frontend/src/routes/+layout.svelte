<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { createIcons, icons } from 'lucide';
	import { onMount } from 'svelte';

	let { children } = $props();

	onMount(() => {
		const renderIcons = () =>
			createIcons({
				icons,
				attrs: {
					'stroke-width': '1.75'
				}
			});

		renderIcons();
		const cleanup = afterNavigate(() => renderIcons());
		return cleanup;
	});
</script>

<svelte:head><link
  href={favicon}
  rel="icon"
/></svelte:head>
{@render children()}
