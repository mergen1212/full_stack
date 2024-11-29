<script lang="ts">
	
	import { fade, fly } from 'svelte/transition';
	
	import { search } from '$lib/store/search.svelte';
	import type { AnimeList } from '$lib/models/ApiModel/entity';
	
	let { Anime }: { Anime: AnimeList } = $props();
</script>

<div class="pb-4 bg-gray-100 border-t border-b border-gray-300">
	<div class="mx-auto px-4 w-full max-w-screen-lg">
		<div class="grid gap-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
			{#if search.value!==null}
				{#await search.value}
					<div class="flex justify-center items-center h-64">
						<div
							class="animate-spin rounded-full h-20 w-20 border-t-4 border-b-4 border-blue-600 shadow-lg"
						></div>
					</div>
				{:then res}
					{#if res}
						{#each res.data as i}
							<div
								class="text-center"
								in:fly={{ y: 50, duration: 800, delay: 200 }}
								out:fade={{ duration: 300 }}
							>
								<a href={`uid/${i.mal_id}/`}>
									<img
										src={i.images.webp.image_url}
										alt="loading"
										class="max-w-full max-h-72 h-auto"
									/>
									<p>{i.title}</p>
								</a>
							</div>
						{/each}
					{/if}
				{/await}
			{:else if search.value === null}
				{#each Anime.data as i}
					<div class="text-center" out:fade={{ duration: 50 }}>
						<a href={`uid/${i.mal_id}/`}>
							<img src={i.images.webp.image_url} alt="loading" class="max-w-full max-h-72 h-auto" />
							<p>{i.title}</p>
						</a>
					</div>
				{/each}
			{/if}
		</div>
	</div>
</div>
