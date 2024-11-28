<script lang="ts">
	import { fade } from 'svelte/transition';
</script>

<svelte:head>
	<title>О нас | Anime-flex</title>
	<meta name="description" content="Learn more about our Anime application" />
</svelte:head>

<div in:fade={{ duration: 300 }} class="max-w-2xl mx-auto p-6">
	<h1 class="text-4xl font-bold text-pink-600">О нас</h1>
	<p class="mt-4 text-gray-800">Добро пожаловать на страницу «О нас» нашего аниме-приложения!</p>
	<p class="mt-2 text-gray-700">
		Аниме — это увлекательная форма развлечения, сочетающая в себе искусство, повествование и музыку
		для создания захватывающих впечатлений. Он имеет богатую историю и разнообразие жанров.
	</p>
	<h2 class="mt-6 text-2xl font-semibold text-purple-600">
		Ключевые особенности нашего аниме-приложения:
	</h2>
	<ul class="list-disc list-inside mt-2 text-gray-700">
		<li>Обширная база данных аниме</li>
		<li>Управление списком наблюдения</li>
		<li>Обсуждения сообщества</li>
	</ul>
	<p class="mt-4 text-gray-800">
		Наша цель — предоставить вам лучший опыт просмотра аниме и создать дружелюбное сообщество
	</p>
</div>
