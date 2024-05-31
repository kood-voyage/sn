<script context="module">
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';
</script>

<script lang="ts">
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let data: RowType[] | null = null;
	const searchQuery = writable('');

	async function getAllUsers() {
		const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/all`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'Access-Control-Request-Method': 'GET'
			},
			credentials: 'include'
		});

		if (response.ok) {
			return await response.json();
		} else {
			throw new Error('Failed to fetch users');
		}
	}

	onMount(async () => {
		try {
			data = await getAllUsers();
		} catch (error) {
			console.error(error);
		}
	});

	$: filteredUsers = data ? data.data.filter(user => {
		const lowerSearchQuery = $searchQuery.toLowerCase();
		return user.username.toLowerCase().includes(lowerSearchQuery);
	}) : [];
</script>

<svelte:head>
	<title>users</title>
</svelte:head>

<div class="sm:p-4">
	<input
		type="text"
		class="flex mx-auto my-2 w-[420px] p-2 rounded-md border dark:bg-neutral-900"
		placeholder="Search..."
		bind:value={$searchQuery}
	/>
	<hr class="w-[440px] m-auto" />

	{#if data === null}
		<p>Loading...</p>
	{:else}
		{#if filteredUsers.length > 0}
			{#each filteredUsers as user (user.id)}
				<div
					class="flex m-auto w-[420px] p-2 hover:bg-neutral-200 dark:hover:bg-neutral-900 rounded-md"
				>
					<a href="/app/u/{user.username}" class="flex">
						<img
							src={user.avatar}
							alt="user-avatar"
							class="rounded-full w-12 h-12 mr-2 object-cover"
						/>
						<div>
							<p class="font-bold">{user.username}</p>
							<p class="text-xs text-slate-400 dark:text-slate-600">
								{user.first_name}
								{user.last_name}
							</p>
							<!-- <p class="text-xs text-slate-600">{user.description == null ? "..." : user.description}</p> -->
						</div>
					</a>
				</div>
			{/each}
		{:else}
			<p class="text-center text-gray-500">No users found</p>
		{/if}
	{/if}
</div>
