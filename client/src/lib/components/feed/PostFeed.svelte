<script lang="ts">
	import TimeAgo from 'javascript-time-ago';
	import en from 'javascript-time-ago/locale/en';

	import UserHover from '$lib/components/feed/UserHover.svelte';

	import PostTimeAgo from './PostTimeAgo.svelte';
	import GroupHover from './GroupHover.svelte';

	export let post;

	console.log(post);

	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');
</script>

<div
	class="w-full max-w-[680px] max-h-[800px] h-full flex flex-col rounded-md bg-neutral-200 shadow-lg dark:bg-neutral-900 gap-y-2 my-2"
>
	<div class="p-4">
		<div class="flex">
			<UserHover postAuthor={post.user_information} username={false} avatar={true} />

			<div class="ml-2">
				<p class="align-middle justify-center text-sm">
					<UserHover postAuthor={post.user_information} username={true} avatar={false} />
				</p>

				<PostTimeAgo created_at={post.created_at} />
			</div>

			{#if post.group_name !== ''}
				<!-- <a href={`/app/g/${post.group_name}`} class="justify-end ml-auto cursor-pointer"
					><Badge>{post.group_name}</Badge></a
				> -->

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
			<img src={post.image_path[0]} alt="post_image" class="max-h-[600px]" />
		</a>
	</body>

	<footer class="p-4"></footer>
</div>

<!-- <div class="flex">
		<Tooltip.Root>
			<Tooltip.Trigger>
				<div class="flex">
					<p class="text-muted-foreground text-sm items-center self-center flex">
						{timeAgo.format(new Date(created_at), 'mini')}
					</p>
				</div>
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p class="z-[99999]">{new Date(created_at)}</p>
			</Tooltip.Content>
		</Tooltip.Root>

		<Tooltip.Root>
			<Tooltip.Trigger>
				<div class="flex">
					<Globe class="w-5 h-5 ml-1 p-[3px] items-center self-center flex text-muted-foreground" />
				</div>
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p>{'public'}</p>
			</Tooltip.Content>
		</Tooltip.Root>
	</div> -->
