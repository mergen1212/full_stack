<script lang="ts">
	import type { AnimeList } from '$lib/models/ApiModel/entity';
	import { search } from '$lib/store/search.svelte';
	let input: string = $state('');
	
	const fetchAnime = async (query: string): Promise<AnimeList> => {
		const url = `https://api.jikan.moe/v4/anime?q=${encodeURIComponent(query)}`;
		const res = await fetch(url);
		if (!res.ok) {
			throw new Error('Network response was not ok');
		}
		return res.json();
	};
	
	const searchAnimeEvent = (e: KeyboardEvent) => {
		switch (e.key) {
			case 'Enter':
				searchAnime();
				break;
			case 'Escape':
				search.value = null;
				break;
			default:
				break;
		}
	};
	$effect(() => {
		input
		const id = setTimeout(() => {
			
			if (input!=='') {
				
				search.value = fetchAnime(input);
			}
		}, 1500);

		return () => {
			
			clearTimeout(id)
		};
	});

	const searchAnime = () => {
		search.value = fetchAnime(input);
	};
</script>

<svelte:window on:keydown={searchAnimeEvent} />
<div class="relative">
	<div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
		<svg
			class="w-4 h-4 text-gray-500"
			aria-hidden="true"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 20 20"
		>
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
			/>
		</svg>
	</div>
	<input
		type="search"
		bind:value={input}
		placeholder="Search..."
		class="w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500"
	/>
	<button
		onclick={searchAnime}
		class="absolute right-2.5 bottom-2.5 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2"
	>
		Search
	</button>
</div>
