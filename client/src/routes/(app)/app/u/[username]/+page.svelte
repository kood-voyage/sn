<script lang="ts">
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';


	import Post from '$lib/components/Post.svelte';

	export let data;

	const { user, posts, params } = data;



	// async function getUser(username: string) {
	// 	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/get/${username}`, {
	// 		method: 'GET',
	// 		headers: {
	// 			'Content-Type': 'application/json',
	// 			'Access-Control-Request-Method': 'GET'
	// 		},
	// 		credentials: 'include'
	// 	});

	// 	if (response.ok) {
	// 		return await response.json();
	// 	} else {
	// 		throw new Error('Failed to fetch users');
	// 	}
	// }

	// async function getFollowers(userId: string) {
	// 	if (userId === undefined) {
	// 		return 0;
	// 	}

	// 	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/followers/${userId}`, {
	// 		method: 'GET',
	// 		headers: {
	// 			'Content-Type': 'application/json',
	// 			'Access-Control-Request-Method': 'GET'
	// 		},
	// 		credentials: 'include'
	// 	});

	// 	console.log(await response.json());

	// 	if (response) {
	// 		return await response.json();
	// 	} else {
	// 		throw new Error('Failed to fetch users');
	// 	}
	// }

	// async function getFollowing(userId: string) {
	// 	if (userId === undefined) {
	// 		return 0;
	// 	}

	// 	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/following/${userId}`, {
	// 		method: 'GET',
	// 		headers: {
	// 			'Content-Type': 'application/json',
	// 			'Access-Control-Request-Method': 'GET'
	// 		},
	// 		credentials: 'include'
	// 	});

	// 	console.log(await response.json());

	// 	if (response) {
	// 		return await response.json();
	// 	} else {
	// 		throw new Error('Failed to fetch folowing');
	// 	}
	// }

	// async function getPosts(userId: string) {
	// 	if (userId === undefined) {
	// 		return 0;
	// 	}

	// 	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/posts/${userId}`, {
	// 		method: 'GET',
	// 		headers: {
	// 			'Content-Type': 'application/json',
	// 			'Access-Control-Request-Method': 'GET'
	// 		},
	// 		credentials: 'include'
	// 	});

	// 	await response.json();

	// 	if (response) {
	// 		return await response.json();
	// 	} else {
	// 		throw new Error('Failed to fetch posts');
	// 	}
	// }

	// onMount(async () => {
	// 	try {
	// 		user = await getUser(data.username);
	// 	} catch (error) {
	// 		console.error(error);
	// 	}

	// 	try {
	// 		postsData = await getPosts(user.data.id);

	// 		console.log(postsData);
	// 	} catch (error) {
	// 		console.error(error);
	// 	}
	// });
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

					<!-- <p class=" text-xs font-bold mr-2">{data.user.id}</p> -->

					<!-- 
					{#if $currentUserStore.id !== data.user.id}
						{#if isCurrentUserFollowing}
							<form action="?/unfollow" method="post">
								<input type="text" hidden name="target_id" value={data.user.id} />

								<button class="text-sm px-5 rounded-md bg-secondary" type="submit">
									unfollow
								</button>
							</form>
						{:else}
							<form action="?/follow" method="post">
								<input type="text" hidden name="target_id" value={data.user.id} />

								<button class="text-sm px-5 rounded-md bg-primary" type="submit"> follow </button>
							</form>
						{/if}
 -->

					<!-- 
					or unfollow button -->

					<!-- 					
					{:else}
						<a href="/app/create-post"
							><button class="text-sm px-5 rounded-md border"> Create Post</button></a
						>
						<a href="/app/settings">
							<button class="text-sm px-5 rounded-md border"> Settings</button></a
						>
					{/if}
 -->
				</div>

				<div class="hidden md:block">
					<div class="flex dark:bg-neutral-900 w-72 rounded-sm absolute right-1 bottom-4">
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">{'none'}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">count</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center hover:bg-neutral-200 dark:hover:bg-neutral-800 p-1"
						>
							<span class="font-bold">count</span> following
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
							<span class="font-bold">{'none'}</span> posts
						</div>
						<div
							class="text-xs w-1/3 border-r text-center p-4 hover:bg-neutral-200 dark:hover:bg-neutral-800"
						>
							<span class="font-bold">followersCount</span> followers
						</div>
						<div
							class="text-xs w-1/3 text-center p-4 hover:rounded-r-lg hover:bg-neutral-200 dark:hover:bg-neutral-800"
						>
							<span class="font-bold">followingCount</span> following
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
			{#if posts !== undefined}
				<p>"qweq"</p>
				{#each posts.data as data}
					<Post {data} />
				{/each}
			{:else}
				<p>Loading...</p>
			{/if}
		</div>
	</div>
</main>

<style>
</style>
