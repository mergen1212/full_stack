<script lang="ts">
	import { base } from '$app/paths';
</script>

<header class="flex bg-gray-800 text-white p-2">
	<nav class="mt-2">
		<ul class="flex list-none p-0 m-0">
			<li class="mx-4">
				<a
					href="{base}/"
					class="text-white no-underline transition-colors duration-200 hover:text-gray-300"
					>Главная</a
				>
			</li>
			<li class="mx-4">
				<a
					href="{base}/about"
					class="text-white no-underline transition-colors duration-200 hover:text-gray-300"
					>web</a
				>
			</li>
		</ul>
	</nav>
</header>
