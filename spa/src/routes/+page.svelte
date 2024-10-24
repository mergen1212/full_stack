<script lang="ts">
	import CurrentAnimeSisson from '$lib/components/CurrentAnimeSisson.svelte';
	import Info from '$lib/components/Info.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<Info />
<CurrentAnimeSisson Anime={data.InfoAnime} />
