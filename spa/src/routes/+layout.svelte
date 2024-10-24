<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
	let { children } = $props();
</script>

<svelte:head>
	<title>Anime</title>
	<meta name="description" content="Anime innformation" />
</svelte:head>

<Header />
{@render children()}
