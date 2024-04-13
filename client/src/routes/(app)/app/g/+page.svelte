<script lang="ts">
	import type { GroupJson } from '$lib/server/api/group-requests';
	import Plus from 'svelte-radix/Plus.svelte';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { Dash } from 'svelte-radix';

	export let data: PageData;

	const groups = data.groups.data;
	// console.log('DATA', groupsResp);

	// const groups = [
	// 	{ name: 'Team Valor', members: 5, description: 'Fearless warriors fighting for victory.' },
	// 	{
	// 		name: 'Pixel Rangers',
	// 		members: 8,
	// 		description: 'Guardians of the gaming realm, pixel by pixel.'
	// 	},
	// 	{ name: 'Quest Keepers', members: 3, description: 'Explorers on an epic quest for glory.' },
	// 	{ name: 'Level Legends', members: 6, description: 'Ascending to legendary gaming heights.' },
	// 	{ name: 'Game Guardians', members: 4, description: 'Protectors of the gaming universe.' },
	// 	{
	// 		name: 'Console Crusaders',
	// 		members: 7,
	// 		description: 'Crusading through realms via console adventures.'
	// 	},
	// 	{ name: 'Joystick Jugglers', members: 2, description: 'Masters of precision and control.' },
	// 	{
	// 		name: 'Elite Empire',
	// 		members: 9,
	// 		description: 'Ruling the gaming empire with elite skills.'
	// 	},
	// 	{
	// 		name: 'Respawn Rebels',
	// 		members: 1,
	// 		description: 'Rebels with a cause, rising again and again.'
	// 	},
	// 	{
	// 		name: 'Warrior Wizards',
	// 		members: 5,
	// 		description: 'Wizards wielding gaming magic on the battlefield.'
	// 	},
	// 	{
	// 		name: 'Warrior Wizards1',
	// 		members: 4,
	// 		description: 'Wizards wielding gaming magic on the battlefield.'
	// 	},
	// 	{
	// 		name: 'Warrior Wizards2',
	// 		members: 6,
	// 		description: 'Wizards wielding gaming magic on the battlefield.'
	// 	},
	// 	{
	// 		name: 'Warrior Wizards3',
	// 		members: 1,
	// 		description: 'Wizards wielding gaming magic on the battlefield.'
	// 	}
	// ];
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
									<span>
										{group.members}
									</span>
								{:else}
									<span>0</span>
								{/if}
							</p>
							<p class="text-sm text-slate-500">{group.description}</p>
						</div>
					</a>
					<!-- <div class="flex">
				<button class="text-sm p-1 bg-blue-500">follow</button>
			</div> -->
				</div>
			{/each}
		{/if}
	</div>
</div>
