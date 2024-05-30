<script lang="ts">
	import PostFormComponent from '../../../../lib/components/forms/PostForm.svelte';
	export let data;
</script>

<svelte:head>
	<title>create post</title>
	<html lang="en" />
</svelte:head>

<div class="w-full">
	<p class="text-2xl py-2 flex items-center justify-center">Create post</p>

	<div class="w-2/3 mx-auto">
		<PostFormComponent data={data.form} community_id={''} />
	</div>
</div>
