<script lang="ts">
	import Plus from 'svelte-radix/Plus.svelte';
	import type { PageData } from './$types';
	import { currentUserStore } from '$lib/store/user-store';
	import type { UserType } from '$lib/types/user';
	import { writable } from 'svelte/store';

	export let data: PageData;

	const currentUser = $currentUserStore as UserType;
	const groups = data.groups;
	let renderedGroupIds: string[] = [];
	const searchQuery = writable('');

	if (groups) {
		groups.forEach((group) => {
			if (currentUser && 'id' in currentUser) {
				group.members?.forEach((member) => {
					if (member.id == currentUser.id) {
						renderedGroupIds.push(group.id);
					} else {
						if (group.creator_id == currentUser.id) {
							renderedGroupIds.push(group.id);
						}
					}
				});
			}
		});
	}

	$: filteredGroups = groups ? groups.filter(group => {
		const lowerSearchQuery = $searchQuery.toLowerCase();
		return group.name.toLowerCase().includes(lowerSearchQuery);
	}) : [];
</script>

<svelte:head>
	<title>groups</title>
</svelte:head>

<div class="sm:p-12">
	<input
		type="text"
		class="flex mx-auto my-2 w-[420px] p-2 rounded-md border dark:bg-neutral-900"
		placeholder="Search..."
		bind:value={$searchQuery}
	/>
	<hr class="w-[440px] m-auto" />

	<div class="overflow-hidden h">
		<div
			class="flex m-auto w-[420px] p-2 mt-1 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md overflow-hidden"
		>
			<a href="/app/create-group" class="flex w-full">
				<Plus alt="user-avatar" class="w-16 mr-2" />
				<div class="">
					<p class="font-bold h-full content-center justify-center align-middle">Create Group</p>
				</div>
			</a>
		</div>
		{#if filteredGroups.length > 0}
			{#each filteredGroups as group}
				{#if ((group.privacy == 'public' || group.members) && (!group.members || group.members.some((member) => member.id == currentUser.id))) || group.privacy == 'public'}
					<div
						class="flex m-auto w-[420px] p-2 mb-2 hover:bg-neutral-200 dark:hover:bg-neutral-900 border rounded-md overflow-hidden {group.creator_id ===
						$currentUserStore.id
							? 'bg-emerald-900/20'
							: 'bg-neutral-900/20'}"
					>
						<a href="/app/g/{group.name.replace(/\s/g, '_')}" class="flex w-full">
							<img src={group.image_path} alt="user-avatar" class="w-16 mr-2" />
							<div>
								<p class="font-bold">{group.name}</p>
								<p class="text-sm text-blue-500">
									Members:
									{#if group.members != null}
										<span>{group.members.length}</span>
									{:else}
										<span>0</span>
									{/if}
								</p>
								<p class="text-sm text-slate-500">{group.description}</p>
							</div>

							<p>{group.creator_id === $currentUserStore.id ? 'creator' : ''}</p>
							<p>
								{group.members.some((member) => member.id == currentUser.id) ? ' ->member' : ''}
							</p>
						</a>
					</div>
				{/if}
			{/each}
			{#each filteredGroups as group}
				{#if group.privacy == 'private'}
					{#if !renderedGroupIds.includes(group.id)}
						<div
							class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md overflow-hidden"
						>
							<div class="flex w-full opacity-50">
								<img src={group.image_path} alt="user-avatar" class="w-16 mr-2" />
								<div>
									<p class="font-bold">{group.name} <span class="text-red-500">Private</span></p>
									<p class="text-sm text-blue-500">
										Members:
										{#if group.members != null}
											<span>{group.members.length}</span>
										{:else}
											<span>0</span>
										{/if}
									</p>
									<p class="text-sm text-slate-500">{group.description}</p>
								</div>
							</div>
						</div>
					{/if}
				{/if}
			{/each}
		{:else}
			<p class="text-center text-gray-500">No groups found</p>
		{/if}
	</div>
</div>
