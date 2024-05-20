<script lang="ts">
	export let data;

	let empty = data.feed.data.data === null;

	import PostFeed from '$lib/components/feed/PostFeed.svelte';
</script>

<div class=" flex flex-col items-center h-screen overflow-y-scroll">
	{#if !empty}
		{#each data?.feed?.data?.data as post, i (post.id)}
			<PostFeed {post} />
		{/each}
	{:else}
		<div class="flex items-center h-full">
			<div class="text-center">
				<img src={'confused.png'} alt="no feed" />
				<p class="text-2xl">Opps...Empty Feed</p>
				<p>Follow someone or Join a Group</p>
			</div>
		</div>
	{/if}
</div>
