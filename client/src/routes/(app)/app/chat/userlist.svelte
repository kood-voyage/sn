<script lang="ts">
	import { invalidate, invalidateAll } from '$app/navigation';
	import { currentUserStore } from '$lib/store/user-store';
	import { userStatusStore } from '$lib/store/websocket-store';
	import type { UserModel, UserType } from '$lib/types/user';
	import { onMount } from 'svelte';
	import type { ChatsWithUsers } from './+layout';
	import PeopleSearch from './people-search.svelte';
	import UserChatCard from '$lib/components/chat/UserChatCard.svelte';

	let searchQuery = '';
	let userData = $currentUserStore;
	let userStatus = $userStatusStore;
	export let users: UserType[] = [];
	export let chats: ChatsWithUsers = {};

	let filteredPeople: displayData[] = [];

	$: display_data = getDisplayData(chats);
	$: chats_arr = display_data;

	$: users = users.filter((person) => {
		if (display_data.find((value) => value.user.id == person.id)) {
			return false;
		}
		return true;
	});

	$: if (chats_arr.length != 0 && chats_arr != undefined) {
		filteredPeople = display_data.filter((val) => {
			return val.user.username.toLowerCase().includes(searchQuery.toLowerCase());
		});
	}

	type displayData = {
		chat_id: string;
		user: UserType;
		group?: any;
	};

	function getDisplayData(chats: ChatsWithUsers): displayData[] {
		const users: displayData[] = Object.entries(chats).map(([chat_id, group_data]) => {
			const users = group_data.users;
			const group = group_data.group;

			// BEFORE ALL THIS SHOULD GO GROUP CHECK AS WELL
			const output: displayData = {
				chat_id,
				user: userData
			};

			if (users.length == 1) {
				const other_user = users[0];
				output.user = other_user;
			}

			return output;
		});
		return users;
	}

	async function handleSubmission() {
		invalidate((url) => url.pathname == '/api/v1/auth/chats');
	}

	onMount(() => {
		userStatusStore.subscribe((user_status) => {
			userStatus = user_status;
		});
	});
</script>

<div class="overflow-scroll h-screen w-14 sm:w-60 bg-neutral-50 dark:bg-neutral-900">
	<!-- search group / friends -->
	<div class="border-b-2 border-neutral-300 dark:border-neutral-950 h-22 hidden sm:block">
		{#if users.length != 0}
			<PeopleSearch userInfo={users} on:submit={handleSubmission} />
		{:else}
			<div class="p-2 pb-1">
				<p
					class="text-sm rounded-md h-fit w-full p-1 border select-none bg-neutral-300 dark:bg-neutral-950"
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
	<div class="sm:p-2">
		<!-- header h1 -->
		<div class="hidden sm:block p-2">
			<p class="text-xs select-none">DIRECT MESSAGES</p>
		</div>

		<ol>
			{#if filteredPeople.length != 0}
				{#each filteredPeople as chat (chat)}
					<li>
						<a
							href={'/app/chat/' + chat.chat_id + '/'}
							class="user hover:bg-slate-300 dark:hover:bg-slate-800 text-center select-none flex w-full py-[6px] my-[1px] items-center rounded-sm"
						>
							<UserChatCard user={chat.user} />
						</a>
					</li>
				{/each}
			{/if}
		</ol>
	</div>
</div>
