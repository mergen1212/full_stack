<script lang="ts">
	import Search from '$lib/components/Search.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>Anime-flex {data.InfoAnime.data.title}</title>
	<meta name="description" content={data.InfoAnime.data.synopsis} />
</svelte:head>
<Search />
<div class="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-5">
	<div class="text-center mb-5">
		{#if data && data.InfoAnime && data.InfoAnime.data}
			<h1 class="text-2xl font-bold mb-2">{data.InfoAnime.data.title}</h1>
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
		{#if data && data.InfoAnime && data.InfoAnime.data}
			<img
				src={data.InfoAnime.data.images.webp.image_url}
				alt="Anime Image"
				class="max-w-full h-auto mb-2"
			/>
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
		<div>
			{#if data && data.InfoAnime && data.InfoAnime.data}
				<p class="text-base">{data.InfoAnime.data.synopsis}</p>
			{:else}
				<p class="text-lg">Loading...</p>
			{/if}
		</div>
	</div>
	<div class="flex flex-col items-center justify-center mt-5">
		{#if data && data.InfoAnime && data.InfoAnime.data}
			{#each data.InfoAnime.data.genres as genre}
				<span class="text-base mb-2">{genre.name}</span>
			{/each}
		{:else}
			<p class="text-lg">Loading...</p>
		{/if}
	</div>
</div>
