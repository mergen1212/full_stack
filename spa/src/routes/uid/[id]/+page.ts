import type { LoadEvent } from '@sveltejs/kit';

export interface Root {
	data: Data;
}

export interface Data {
	mal_id: number;
	url: string;
	images: Images;
	trailer: Trailer;
	approved: boolean;
	titles: Title[];
	title: string;
	title_english: string;
	title_japanese: string;
	title_synonyms: any[];
	type: string;
	source: string;
	episodes: number;
	status: string;
	airing: boolean;
	aired: Aired;
	duration: string;
	rating: string;
	score: number;
	scored_by: number;
	rank: number;
	popularity: number;
	members: number;
	favorites: number;
	synopsis: string;
	background: string;
	season: string;
	year: number;
	broadcast: Broadcast;
	producers: Producer[];
	licensors: Licensor[];
	studios: Studio[];
	genres: Genre[];
	explicit_genres: any[];
	themes: Theme[];
	demographics: any[];
}

export interface Images {
	jpg: Jpg;
	webp: Webp;
}

export interface Jpg {
	image_url: string;
	small_image_url: string;
	large_image_url: string;
}

export interface Webp {
	image_url: string;
	small_image_url: string;
	large_image_url: string;
}

export interface Trailer {
	youtube_id: string;
	url: string;
	embed_url: string;
	images: Images2;
}

export interface Images2 {
	image_url: string;
	small_image_url: string;
	medium_image_url: string;
	large_image_url: string;
	maximum_image_url: string;
}

export interface Title {
	type: string;
	title: string;
}

export interface Aired {
	from: string;
	to: string;
	prop: Prop;
	string: string;
}

export interface Prop {
	from: From;
	to: To;
}

export interface From {
	day: number;
	month: number;
	year: number;
}

export interface To {
	day: number;
	month: number;
	year: number;
}

export interface Broadcast {
	day: string;
	time: string;
	timezone: string;
	string: string;
}

export interface Producer {
	mal_id: number;
	type: string;
	name: string;
	url: string;
}

export interface Licensor {
	mal_id: number;
	type: string;
	name: string;
	url: string;
}

export interface Studio {
	mal_id: number;
	type: string;
	name: string;
	url: string;
}

export interface Genre {
	mal_id: number;
	type: string;
	name: string;
	url: string;
}

export interface Theme {
	mal_id: number;
	type: string;
	name: string;
	url: string;
}

export const load = async ({ fetch, params }: LoadEvent) => {
	const res = await fetch(`https://api.jikan.moe/v4/anime/` + params.id);
	const data = (await res.json()) as Root;
	return {
		i: { InfoAnime: data },
		maxAge: 60
	};
};
