import type { LoadEvent } from '@sveltejs/kit';
import type { PageLoad } from '../$types';
import type { AnimeList } from '$lib/entity';

export const load = (async ({ fetch, params }: LoadEvent) => {
	const res = await fetch(`https://api.jikan.moe/v4/anime/` + params.id);
	const data = (await res.json()) as AnimeList;
	return {
		InfoAnime: data
	};
}) satisfies PageLoad;
