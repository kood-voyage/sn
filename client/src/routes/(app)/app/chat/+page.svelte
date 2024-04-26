<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import type { UserRowType } from '$lib/server/db/user';
	import { currentUserStore } from '$lib/store/user-store';
	import type { UserModel } from '$lib/types/user';
	import type { PageData } from './$types';
	import type { ChatsWithUsers } from './+page.server';
	import PeopleSearch from './people-search.svelte';

	export let data: PageData;
	// console.log(data.data);

	// $: console.log(data);
	let searchQuery = '';

	let userData = $currentUserStore;
	let people: UserRowType[] = [];
	let chats: ChatsWithUsers = {};
	if (data.ok) {
		people = data.data.usersData;
		chats = data.data.chatsData;
	}

	let filteredPeople: UserModel[] = [];

	$: chats_arr = Object.values(chats);

	// console.log(chats);
	$: if (chats_arr.length != 0 && chats_arr != undefined) {
		filteredPeople = Object.entries(chats).map((chat) => {
			// console.log('CHAT >>>', chat);
			const key = chat[0];
			const value = chat[1];

			const users = value.users;
			const group_id = value.group;
			// console.log(value);

			for (const user of users) {
				// console.log(user);
			}

			// console.log(value.group_id != '' && value.group_id != undefined);
			// if (group_id != '' && group_id != undefined) {
			// }

			// console.log('Key >>>', key);
			// console.log('Value >>>', value);
			// if (chat.length == 0) return false
			// if (chat.length == 1) return true
			// const out = chat.filter
			// return chat[0].username.toLowerCase().includes(searchQuery.toLowerCase());
			return;
		});
		filteredPeople = [];
		// .filter((chat) => {
		// 	// console.log('CHAT >>>', chat);
		// 	// const key = chat[0];
		// 	// const value = chat[1];
		// 	// const users = value.users;
		// 	// const group_id = value.group_id;
		// 	// for (const user of users) {
		// 	// 	console.log(user);
		// 	// }
		// 	// // console.log(value.group_id != '' && value.group_id != undefined);
		// 	// if (value.group_id != '' && value.group_id != undefined) {
		// 	// }
		// 	// console.log('Key >>>', key);
		// 	// console.log('Value >>>', value);
		// 	// if (chat.length == 0) return false
		// 	// if (chat.length == 1) return true
		// 	// const out = chat.filter
		// 	// return chat[0].username.toLowerCase().includes(searchQuery.toLowerCase());
		// })
		// .map((val) => {
		// 	// console.log(val);
		// 	return val[0];
		// });
		// .slice(0, 6);

		// console.log(filteredPeople);
	}

	function handleSubmission(event: CustomEvent) {
		console.log('Submitted from PeopleSearch:', event.detail);
	}
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
			{#if data.ok}
				<PeopleSearch userInfo={people} on:submit={handleSubmission} />
			{:else}
				<div class="p-2 pb-1">
					<p
						class="text-sm rounded-md h-fit w-full p-1 border select-none bg-slate-300 dark:bg-slate-950"
					>
						Not Found
					</p>
				</div>
			{/if}

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
				<p class="text-xs select-none">DIRECT MESSAGES</p>
			</div>

			<ol class="">
				{#if filteredPeople.length != 0}
					{#each filteredPeople as person (person)}
						<!-- {console.log('HERE', person)} -->
						<!-- <img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" /> -->
						<li class="user hover:bg-slate-300 dark:hover:bg-slate-800 text-center select-none">
							<img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" />
							<p class=" align-middle justify-center text-center">
								{person.username}
							</p>
						</li>
						<!-- <li class="person-item select-none text-center hover:bg-slate-500">{person.username}</li> -->
					{/each}
				{/if}
			</ol>
		</div>
	</div>

	<!-- messenger -->

	<div class="w-full">
		<!-- user chat header -->

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

	/* .active {
		@apply bg-slate-400 dark:bg-slate-700;
	} */
	/* 
	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	} */
</style>
