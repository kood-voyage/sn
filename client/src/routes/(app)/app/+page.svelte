<script lang="ts">
	export let data;

	let feed = data?.feed?.data?.data;

	import PostFeed from '$lib/components/feed/PostFeed.svelte';
</script>

<div class=" flex flex-col items-center h-screen overflow-y-scroll">
	{#each feed as post, i (post.id)}
		<PostFeed {post} />
	{/each}
</div>
