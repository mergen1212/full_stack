<script lang="ts">
	import CurrentAnimeSisson from '$lib/components/CurrentAnimeSisson.svelte';
	import Info from '$lib/components/Info.svelte';
	import Search from '$lib/components/Search.svelte';
	
	import type { PageData } from './$types';
	
	let { data }: { data: PageData } = $props();
</script>
<svelte:head>
	<title>Anime-flex</title>
	<meta name="description" content="Anime innformation" />
</svelte:head>
<Search />
<Info />
<CurrentAnimeSisson Anime={data.props.data} />
