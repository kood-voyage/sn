<script lang="ts">
	import type { RowType } from '$lib/server/db/user';

	export let data;
	const userList = data.data as RowType[];

	console.log('NEW DATA CONSOLE >>>', data);

	// The array now contains 100 names
</script>

<svelte:head>
	<title>users</title>
</svelte:head>

<div class="p-12">
	<p class="text-md">Suggested for you</p>
	{#each userList as user}
		<div class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md">
			<a href="/app/u/{user.username}" class="flex">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={user.username}"
					alt="user-avatar"
					class="rounded-full w-12 mr-2"
				/>
				<div>
					<p class="font-bold">{user.username}</p>
					<p class="text-sm text-slate-500">Full Name or description</p>
				</div>
			</a>
			<!-- <div class="flex">
				<button class="text-sm p-1 bg-blue-500">follow</button>
			</div> -->
		</div>
	{/each}
</div>
