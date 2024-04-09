<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import type { UserRowType } from '$lib/server/db/user';
	import type { PageData } from './$types';
	import PeopleSearch from './people-search.svelte';

	export let data: PageData;
	// console.log(data.data);

	let searchQuery = '';

	let people = data.data as UserRowType[];

	$: filteredPeople = people
		.filter((person) => person.username.toLowerCase().includes(searchQuery.toLowerCase()))
		.slice(0, 6);
</script>

<svelte:head>
	<title>Chat</title>
</svelte:head>

<!-- chat page -->
<div class="flex w-full">
	<!-- user list -->
	<div class="overflow-scroll h-screen w-14 sm:w-60 bg-slate-50 dark:bg-slate-900">
		<!-- search group / friends -->
		<div class="border-b-2 border-slate-300 dark:border-slate-950 h-22 hidden sm:block">
			<Dialog.Root>
				<div class="p-2 pb-1">
					<Dialog.Trigger
						class="text-sm rounded-md h-fit w-full  py-1 border dark:hover:bg-slate-800 bg-slate-300 dark:bg-slate-950 "
						>New Chat</Dialog.Trigger
					>
				</div>

				<Dialog.Content class="w-fit h-fit max-h-96">
					{#if data.ok}
						<PeopleSearch userInfo={data.data} />
					{:else}
						<p>Data is not available or an error occurred.</p>
					{/if}
				</Dialog.Content>
			</Dialog.Root>

			<div class="p-2 pt-1">
				<input
					type="text"
					placeholder="Find friend..."
					class="bg-slate-300 dark:bg-slate-950 w-full rounded py-1 px-2 cursor-pointer"
					bind:value={searchQuery}
				/>
			</div>
		</div>

		<!-- user list -->
		<div class=" sm:p-2">
			<!-- header h1 -->

			<div class="hidden sm:block p-2">
				<p class="text-xs">DIRECT MESSAGES</p>
			</div>

			<ol class="">
				{#each filteredPeople as person (person)}
					<!-- <img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" /> -->
					<li class="user hover:bg-slate-300 dark:hover:bg-slate-800 text-center">
						<img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" />
						<p class=" align-middle justify-center text-center">{person.username}</p>
					</li>
					<!-- <li class="person-item select-none text-center hover:bg-slate-500">{person.username}</li> -->
				{/each}
			</ol>
		</div>
	</div>

	<!-- messenger -->

	<div class="w-full">
		<!-- user chat header -->
		<div class="h-12 bg-slate-50 dark:bg-slate-800 px-4">
			<div class="flex p-2 h-full">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed=Nikita"
					alt="avatar"
					class="w-8 rounded-full mr-2"
				/>
				<p class="h-full text-center">Nikita</p>
			</div>
		</div>

		<!-- direct chat -->
		<div class="p-4"></div>
	</div>
</div>

<style>
	div ol li.user {
		@apply flex h-12 cursor-pointer self-center rounded-sm;
	}

	div ol li.user p {
		@apply hidden items-center text-center text-slate-700 dark:text-slate-300;
	}

	@media (min-width: 640px) {
		div ol li.user p {
			@apply ml-2 flex items-center text-center align-middle;
		}
	}

	div ol li.user img {
		@apply h-8 w-8 items-center justify-center rounded-full;
	}

	.active {
		@apply bg-slate-400 dark:bg-slate-700;
	}

	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	}
</style>
