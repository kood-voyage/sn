<script lang="ts">
	import PostFormComponent from '../../../../lib/components/forms/PostForm.svelte';
	export let data;
</script>

<svelte:head>
	<title>create post</title>
	<html lang="en" />
</svelte:head>

<div class="w-full">
	<p class="w-full text-2xl py-2">Create post</p>

	<div class="w-2/3">
		<PostFormComponent data={data.form} community_id={''} />
	</div>
</div>
