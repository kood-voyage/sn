<script lang="ts">
	import type { GroupJson } from '$lib/server/api/group-requests';
	import Plus from 'svelte-radix/Plus.svelte';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { Dash } from 'svelte-radix';
	import { currentUserStore } from '$lib/store/user-store';

	export let data: PageData;

	const currentUser = $currentUserStore;
	const groups = data.groups.data;
	let isMember = false;
	let renderedGroupIds: string[] = [];

	groups?.forEach((group) => {
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
</script>

<svelte:head>
	<title>groups</title>
</svelte:head>

<div class="p-12">
	<input
		type="text"
		class="flex mx-auto my-2 w-[420px] p-2 rounded-md border dark:bg-slate-900"
		placeholder="Search..."
	/>
	<hr class="w-[440px] m-auto" />

	<div class="overflow-scroll h">
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
		{#if groups}
			{#each groups as group}
				{#if ((group.privacy == 'public' || group.members) && (!group.members || group.members.some((member) => member.id == currentUser.id))) || group.privacy == 'public'}
					<div
						class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md overflow-hidden"
					>
						<a href="/app/g/{group.name.replace(/\s/g, '_')}" class="flex w-full">
							<img
								src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={group.name}"
								alt="user-avatar"
								class="w-16 mr-2"
							/>
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
						</a>
					</div>
				{/if}
			{/each}
			{#each groups as group}
				{#if group.privacy == 'private'}
					{#if !renderedGroupIds.includes(group.id)}
						<div
							class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md overflow-hidden"
						>
							<div class="flex w-full opacity-50">
								<img
									src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={group.name}"
									alt="user-avatar"
									class="w-16 mr-2"
								/>
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
		{/if}
	</div>
</div>
