import type { AnimeList } from '$lib/entity';
import type { LoadEvent } from '@sveltejs/kit';

export const load = async ({ fetch }: LoadEvent) => {
	const res = await fetch('https://api.jikan.moe/v4/anime');
	const data = (await res.json()) as AnimeList;
	return {
		InfoAnime: data
	};
};
