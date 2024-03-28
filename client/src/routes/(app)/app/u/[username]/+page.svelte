<script lang="ts">
	import Post from '$lib/components/Post.svelte';
	import type { PageData } from './$types';
	import { currentUserFollowing, currentUserStore } from '$lib/store/user-store';
	export let data: PageData;

	let isCurrentUserFollowing = false;

	// console.log($currentUserFollowing);

	if ($currentUserFollowing.data !== null && $currentUserFollowing.data !== undefined) {
		for (const following of $currentUserFollowing.data) {
			if (following.id === data.user.id) {
				isCurrentUserFollowing = true;
			}
		}
	}

	const followersCount = data.followers ? data.followers.length : 0;
	const followingCount = data.following ? data.following.length : 0;

	const { posts } = data;

	console.log(posts);
</script>

<svelte:head>
	<title>u/{data.username}</title>
</svelte:head>

<!-- user profile page -->

<main class=" overflow-scroll">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- banner img  -->
			<div class="m-auto h-24 sm:h-40 max-w-[1096px] p-0 sm:p-4">
				<img
					class="h-full w-full rounded-none sm:rounded-xl object-cover"
					src={data.user.cover}
					alt="cover"
				/>
			</div>

			<div class="h-8 relative mx-0 sm:mx-4">
				<img
					src={data.user.avatar}
					alt="avatar"
					class="absolute bottom-[1px] left-12 h-20 w-20 rounded-full border-4 border-white object-cover dark:border-slate-950"
				/>

				<div
					class="absolute bottom-1 sm:bottom-3 left-[140px] bg-white dark:bg-slate-950 rounded-2xl flex"
				>
					<p class=" md:text-2xl font-bold mr-2">{data.user.username}</p>
					<!-- <p class=" text-xs font-bold mr-2">{data.user.id}</p> -->
					<!-- <p class="text-xs">
						<span class="bg-slate-500 px-1 rounded-sm">AKA</span>
						{data.first_name}
						{data.last_name}
					</p> -->

					{#if $currentUserStore.id !== data.user.id}
						{#if isCurrentUserFollowing}
							<form action="?/unfollow" method="post">
								<input type="text" hidden name="target_id" value={data.user.id} />

								<button class="text-sm px-5 rounded-md bg-red-500" type="submit"> unfollow </button>
							</form>
						{:else}
							<form action="?/follow" method="post">
								<input type="text" hidden name="target_id" value={data.user.id} />

								<button class="text-sm px-5 rounded-md bg-sky-500" type="submit"> follow </button>
							</form>
						{/if}

						<!-- 
					or unfollow button -->
					{:else}
						<a href="/app/create-post"
							><button class="text-sm px-5 rounded-md border"> Create Post</button></a
						>
						<a href="/app/settings">
							<button class="text-sm px-5 rounded-md border"> Settings</button></a
						>
					{/if}
				</div>

				<div class="hidden md:block">
					<div class="flex dark:bg-slate-900 w-72 rounded-sm absolute right-1 bottom-4">
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1"
						>
							<span class="font-bold">{'none'}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1"
						>
							<span class="font-bold">{followersCount}</span> followers
						</div>
						<div class="text-xs w-1/3 text-center hover:bg-slate-200 dark:hover:bg-slate-800 p-1">
							<span class="font-bold">{followingCount}</span> following
						</div>
					</div>
				</div>
			</div>

			<div class="sm:mx-4 h-8 block md:hidden">
				<div class="">
					<div class="flex bg-slate-100 dark:bg-slate-900 w-full rounded-lg">
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-slate-200 dark:hover:bg-slate-800 hover:rounded-l-lg"
						>
							<span class="font-bold">{'none'}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-slate-200 dark:hover:bg-slate-800"
						>
							<span class="font-bold">{followersCount}</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center p-4 hover:rounded-r-lg hover:bg-slate-200 dark:hover:bg-slate-800"
						>
							<span class="font-bold">{followingCount}</span> following
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- <p>Followers</p>

		{#each data.followers as follower}
			<p>{follower.id}</p>
		{/each}

		<p>Following</p>

		{#each data.following as following}
			<p>{following.id}</p>
		{/each} -->

		<!-- profile activity / posts -->

		
		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-1 p-0 sm:p-4 mt-5 md:mt-0">
			{#if posts !== null}
				{#each posts as data}
					<Post {data} />
				{/each}
			{/if}
		</div>
	</div>
</main>

<style>
</style>
