<script lang="ts">
	import type { UserType } from '$lib/types/user';
	import * as Dialog from '$lib/components/ui/dialog';
	import { enhance } from '$app/forms';
	import { createEventDispatcher } from 'svelte';
	import { newChatCreate } from './new-chat';
	import { currentUserStore } from '$lib/store/user-store';
	import { sendMessage } from '$lib/client/websocket';

	const dispatch = createEventDispatcher();

	let dialogOpen = false;
	let people: UserType[] = [];
	let searchQuery = '';
	let filteredPeople: UserType[] = [];

	export let userInfo: UserType[] | undefined;

	if (userInfo != undefined) {
		people = userInfo;
		people = people.filter((person) => $currentUserStore.id != person.id);
	}

	$: if (people != undefined && people.length != 0) {
		filteredPeople = people
			.filter((person) => person.username.toLowerCase().includes(searchQuery.toLowerCase()))
			.slice(0, 6);
	}
</script>

<Dialog.Root bind:open={dialogOpen}>
	<div class="p-2 pb-1">
		<Dialog.Trigger
			class="text-sm rounded-md h-fit w-full  py-1 border dark:hover:bg-slate-800 bg-slate-300 dark:bg-slate-950 "
			>New Chat</Dialog.Trigger
		>
	</div>

	<Dialog.Content class="w-fit h-fit max-h-96">
		<div class="mx-4">
			<input
				type="text"
				placeholder="Search people..."
				class="search-box"
				bind:value={searchQuery}
			/>

			<!-- <ul class="person-list"> -->

			<ol class="">
				{#each filteredPeople as person (person)}
					<!-- <img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" /> -->

					<form
						method="post"
						use:enhance={async ({ formData, controller, cancel }) => {
							// console.log('FormData >>', formData);
							formData.set('target', person.id);
							// console.log(person.id);

							const createResp = await newChatCreate(formData);
							if (!createResp.ok) {
								controller.abort();
								cancel();
								return;
							}
							dispatch('submit', { detail: 'Data or message from PeopleSearch' });
							dialogOpen = false;
							sendMessage('status', 'direct', person.id, $currentUserStore.id, 2);

							controller.abort();
							cancel();
						}}
					>
						<button type="submit" class="h-full w-full">
							<!-- {person.username} -->
							<li
								class="user hover:bg-slate-300 select-none flex dark:hover:bg-slate-800 text-center"
							>
								<img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" />
								<p class="  align-middle justify-center text-center">
									{person.username}
								</p>
							</li>
						</button>
					</form>
					<!-- <li class="person-item select-none text-center hover:bg-slate-500">{person.username}</li> -->
				{/each}
			</ol>
			<!-- </ul> -->
		</div>
	</Dialog.Content>
</Dialog.Root>

<style>
	.search-box {
		padding: 6px;
		margin-bottom: 20px;
		width: 100%;
		box-sizing: border-box;
	}
	/* .person-list {
		list-style-type: none;
		padding: 0;
	}
	.person-item {
		padding: 5px;
		margin: 5px;
		border-radius: 0.5rem;
		border: 1px solid #ccc;
	} */
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

	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	}
</style>
