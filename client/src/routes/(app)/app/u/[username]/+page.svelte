<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { follow, getUserFollowing, unfollow } from '$lib/client/api/user-requests';

	import Post from '$lib/components/Post.svelte';
	import { currentUserStore, currentUserFollowing } from '$lib/store/user-store.js';
	import toast from 'svelte-french-toast';

	export let data;

	const { user, posts, followers, following } = data;

	let isCurrentUserFollowing = false;

	currentUserFollowing.subscribe((val) => {
		if (val !== null && val !== undefined) {
			if (val.find((following) => following.id == user.data.id) !== undefined) {
				isCurrentUserFollowing = true;
			} else {
				isCurrentUserFollowing = false;
			}
		} else {
			isCurrentUserFollowing = false;
		}
	});

	// $: if ($currentUserFollowing !== null && $currentUserFollowing !== undefined) {
	// 	for (const following of $currentUserFollowing) {
	// 		if (following.id === user.data.id) {
	// 			isCurrentUserFollowing = true;
	// 		}
	// 	}
	// }

	$: followersCount = followers.data !== null ? followers.data.length : 0;
	$: followingCount = following.data !== null ? following.data.length : 0;
	$: postsCount = posts.data !== null ? posts.data.length : 0;
</script>

<svelte:head>
	<title>u/{user.data.username}</title>
</svelte:head>

<!-- user profile page -->

<main class=" overflow-scroll">
	<div class=" m-auto h-full w-full max-w-[1096px]">
		<!-- profile info header -->
		<div class="profile-info relative">
			<!-- cover img  -->
			<div class="m-auto h-24 sm:h-40 max-w-[1096px] p-0 sm:p-4">
				<!-- i know how i can change image position -->
				<img
					class="h-full w-full rounded-none sm:rounded-xl object-cover relative"
					style="object-position: 0% 0%"
					src={user === undefined ? '' : user.data.cover}
					alt="cover"
				/>
			</div>

			<div class="h-8 relative mx-0 sm:mx-4">
				<img
					src={user === undefined ? '' : user.data.avatar}
					alt="avatar"
					class="absolute bottom-[1px] left-12 h-20 w-20 rounded-full border-4 border-white object-cover dark:border-slate-950"
				/>

				<div
					class="absolute bottom-1 sm:bottom-3 left-[140px] bg-white dark:bg-neutral-950 rounded-2xl flex"
				>
					{#if user === undefined}
						<p>Loading...</p>
					{:else}
						<p class=" md:text-2xl font-bold mr-2">{user.data.username}</p>
					{/if}

					{#if $currentUserStore.id !== user.data.id}
						{#if isCurrentUserFollowing}
							<button
								class="text-sm px-5 rounded-md bg-secondary"
								on:click={async () => {
									await unfollow(user.data.id);
									const followingResp = await getUserFollowing($currentUserStore.id);
									if (!followingResp.ok) {
										toast.error("Couldn't get following " + followingResp.message);
										return;
									}
									currentUserFollowing.set(followingResp.data);
								}}
							>
								unfollow
							</button>
						{:else}
							<button
								class="text-sm px-5 rounded-md bg-primary"
								on:click={async () => {
									await follow(user.data.id);
									const followingResp = await getUserFollowing($currentUserStore.id);
									if (!followingResp.ok) {
										toast.error("Couldn't get following " + followingResp.message);
										return;
									}
									currentUserFollowing.set(followingResp.data);
								}}
							>
								follow
							</button>
						{/if}
					{:else}
						<a href="/app/create-post" class="flex"
							><button class="text-sm px-5 rounded-md border"> Create Post</button></a
						>
						<!-- <a href="/app/settings">
							<button class="text-sm px-5 rounded-md border"> Settings</button></a
						> -->

						<!-- <div class="flex items-center">
							
							<SettingsForm />
						</div> -->
					{/if}
				</div>

				<div class="hidden md:block">
					<div class="flex dark:bg-neutral-900 w-72 rounded-sm absolute right-1 bottom-4">
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">{postsCount}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">{followersCount}</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">{followingCount}</span> following
						</div>
					</div>
				</div>
			</div>

			<div class="sm:mx-4 h-8 block md:hidden">
				<div class="">
					<div class="flex bg-neutral-100 dark:bg-neutral-900 w-full rounded-lg">
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-neutral-200 dark:hover:bg-neutral-800 hover:rounded-l-lg"
						>
							<span class="font-bold">{postsCount}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-neutral-200 dark:hover:bg-neutral-800"
						>
							<span class="font-bold">{followersCount}</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center p-4 hover:rounded-r-lg hover:bg-neutral-200 dark:hover:bg-neutral-800"
						>
							<span class="font-bold">{followingCount}</span> following
						</div>
					</div>
				</div>
			</div>
		</div>

		<!-- profile activity / posts -->

		<div class="h-full w-full sm:grid sm:grid-cols-2 md:grid-cols-3 gap-1 p-0 sm:p-4 mt-5 md:mt-0">
			{#if posts.data !== null}
				{#each posts.data as data}
					<Post {data} />
				{/each}
			{:else}
				<p>No posts</p>
			{/if}
		</div>
	</div>
</main>

<style>
</style>
