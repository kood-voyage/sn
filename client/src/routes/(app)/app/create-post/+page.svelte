<script lang="ts">
	export let data;
	import PostFormComponent from './post-form.svelte';
</script>

<svelte:head>
	<title>create post</title>
	<meta name="robots" content="noindex nofollow" />
	<html lang="en" />
</svelte:head>

<div>
	<p class="w-full text-2xl py-2">Create post</p>
	<PostFormComponent data={data.form} />
</div>
