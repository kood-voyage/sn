<script lang="ts">
	import type { RowType } from '$lib/server/db/user';

	export let data;
	const userList = data.data as RowType[];
</script>

<svelte:head>
	<title>users</title>
</svelte:head>

<div class="p-4">
	<input
		type="text"
		class="flex mx-auto my-2 w-[420px] p-2 rounded-md border dark:bg-neutral-900"
		placeholder="Search..."
	/>
	<hr class="w-[440px] m-auto" />

	{#each userList as user}
		<div
			class="flex m-auto w-[420px] p-2 hover:bg-neutral-200 dark:hover:bg-neutral-900 rounded-md"
		>
			<a href="/app/u/{user.username}" class="flex">
				<img src={user.avatar} alt="user-avatar" class="rounded-full w-12 h-12 mr-2 object-cover" />
				<div>
					<p class="font-bold">{user.username}</p>
					<p class="text-xs text-slate-400 dard:text-slate-600">
						{user.first_name}
						{user.last_name}
					</p>
					<!-- <p class="text-xs text-slate-600">{user.description == null ? "..." : user.description}</p> -->
				</div>
			</a>
		</div>
	{/each}
</div>
