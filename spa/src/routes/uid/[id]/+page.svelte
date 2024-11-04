<script lang="ts">
	import Search from '$lib/components/Search.svelte';
	import VideoPlayer from '$lib/components/VideoPlayer.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	const info=data.i.InfoAnime
</script>

<svelte:head>
	<title>Anime-flex {info.data.title}</title>
	<meta name="description" content={info.data.synopsis} />
</svelte:head>
<Search />
<div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-5">
	<div class="text-center mb-5">
		{#if info && info.data}
			<h1 class="text-2xl font-bold mb-2">{info.data.title}</h1>
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
		{#if info && info.data}
			<img
				src={info.data.images.webp.image_url}
				alt={info.data.title}
				class="max-w-full h-auto mb-2"
			/>
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
		<div>
			{#if info && info.data}
				<p class="text-base">{info.data.synopsis}</p>
			{:else}
				<p class="text-lg">Loading...</p>
			{/if}
		</div>
	</div>
	<div class="flex flex-col items-center justify-center mt-5">
		{#if info && info.data}
			{#each info.data.genres as genre}
				<span class="text-base mb-2">{genre.name}</span>
			{/each}
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
	</div>
</div>
