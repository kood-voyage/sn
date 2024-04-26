<script context="module">
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';
</script>

<script lang="ts">
	import { onMount } from 'svelte';

	let data: RowType[] | null = null;

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

	{#if data === null}
		<p>Loading...</p>
	{:else}


		{#each data.data as user (user.id)}
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
						<p class="text-xs text-slate-400 dard:text-slate-600">
							{user.first_name}
							{user.last_name}
						</p>
						<!-- <p class="text-xs text-slate-600">{user.description == null ? "..." : user.description}</p> -->
					</div>
				</a>
			</div>
		{/each}


	{/if}


</div>
