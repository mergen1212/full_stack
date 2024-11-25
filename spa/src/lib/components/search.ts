import type { AnimeList } from '$lib/entity';
import { writable } from 'svelte/store';
export const search = writable<Promise<AnimeList> | null>(null)