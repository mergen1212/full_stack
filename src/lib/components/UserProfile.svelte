<script lang="ts">
	const {
		name = '',
		age = null,
		favouriteColors = [],
		isAvailable = false
	}: {
		name?: string;
		age?: number | null;
		favouriteColors?: string[];
		isAvailable?: boolean;
	} = $props();
	console.log(name, age, favouriteColors, isAvailable);
</script>

<p>My name is {name !== '' ? `${name}` : 'имя не указано'}!</p>
<p>My age is {age !== null ? { age } : 'возраст не указан'}!</p>
<p>My favourite colors are {favouriteColors.join(', ')}!</p>
<p>I am {isAvailable ? 'available' : 'not available'}</p>
