<script lang="ts">
	import UserHover from '$lib/components/feed/UserHover.svelte';
	import PostTimeAgo from './PostTimeAgo.svelte';
	import GroupHover from './GroupHover.svelte';
	export let post;

	const preview =
		post?.image_path !== null
			? post?.image_path[0]
			: 'https://cdn.vectorstock.com/i/500p/65/30/default-image-icon-missing-picture-page-vector-40546530.jpg';

</script>

<div
	class="w-full max-w-[680px] max-h-[800px] h-full flex flex-col rounded-md bg-neutral-200 shadow-lg dark:bg-neutral-900 gap-y-2 my-2"
>
	<div class="sm:p-4">
		<div class="flex">
			<UserHover postAuthor={post.user_information} username={false} avatar={true} />

			<div class="ml-2">
				<p class="align-middle justify-center text-sm">
					<UserHover postAuthor={post.user_information} username={true} avatar={false} />
				</p>

				<PostTimeAgo created_at={post.created_at} />
			</div>

			{#if post.group_name !== ''}
				<GroupHover group_name={post.group_name} />
			{/if}
		</div>

		<div>
			<p>{post.title}</p>

			<p>{post.content}</p>
		</div>
	</div>

	<body class=" bg-black flex justify-center">
		<a href={`/app/post/${post.id}`}>
			<img src={preview} alt="post_image" class="max-h-[600px]" />
		</a>
	</body>

	<footer class="p-4"></footer>
</div>
