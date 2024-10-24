<script>
	import { fade } from 'svelte/transition';
</script>

<svelte:head>
	<title>About | My Anime App</title>
	<meta name="description" content="Learn more about our Anime application" />
</svelte:head>

<div in:fade={{ duration: 300 }} class="max-w-2xl mx-auto p-6">
	<h1 class="text-4xl font-bold text-pink-600">About</h1>
	<p class="mt-4 text-gray-800">Welcome to the About page of our Anime application!</p>
	<p class="mt-2 text-gray-700">
		Anime is a captivating form of entertainment that combines art, storytelling, and music to create
		immersive experiences. It has a rich history and a diverse range of genres.
	</p>
	<h2 class="mt-6 text-2xl font-semibold text-purple-600">Key Features of Our Anime App:</h2>
	<ul class="list-disc list-inside mt-2 text-gray-700">
		<li>Extensive Anime Database</li>
		<li>User Reviews and Ratings</li>
		<li>Personalized Recommendations</li>
		<li>Watchlist Management</li>
		<li>Community Discussions</li>
	</ul>
	<p class="mt-4 text-gray-800">
		This page is a simple example of a Svelte component. It demonstrates basic structure, styling,
		and the use of transitions, all while celebrating the world of Anime.
	</p>
</div>

