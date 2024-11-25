<script lang="ts">
	import type { AnimeList } from '../entity';
	import { search } from "./search";
	let { Anime }: { Anime: AnimeList } = $props();
	console.log(search);
</script>

<div class="pb-4 bg-gray-100 border-t border-b border-gray-300">
	<div class="mx-auto px-4 w-full max-w-screen-lg">
		<div class="grid gap-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
			{#if $search}
			    {#await $search}
					<p></p>
				{:then res}
					{#if res}
						{#each res.data as i}
								<div class="text-center">
									<a href={`uid/${i.mal_id}/`}>
										<img src={i.images.webp.image_url} alt="loading" class="max-w-full max-h-72 h-auto" />
										<p>{i.title}</p>
									</a>
								</div>
						{/each}
					{/if}
				{/await}
			{:else if $search === null}
				{#each Anime.data as i}
					<div class="text-center">
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
