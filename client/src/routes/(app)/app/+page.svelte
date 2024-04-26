<script lang="ts">
	import { sendMessage } from '$lib/client/websocket';
	import {
		messageStore,
		userStatusStore,
		type MessageStore,
		type UserStatusStore
	} from '$lib/store/websocket-store';
	import { currentUserStore } from '$lib/store/user-store';
	import type { User } from '$lib/types/user';
	import { onDestroy } from 'svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	let messages: MessageStore = [];
	let statuses: UserStatusStore = {};
	let currentUser = $currentUserStore as User;

	messageStore.subscribe((value) => {
		messages = value;
	});

	userStatusStore.subscribe((value) => {
		statuses = value;
	});
	// Establish connection when component mounts

	const getName = (user_id: string) => {
		const user = data.allUsers.data?.find((val) => {
			// console.log(val);
			if (val.id == user_id) return true;
			return false;
		});

		return user?.username;
	};
</script>

<ul>
	{#each messages as message (message)}
		<li>{JSON.stringify(message)}</li>
	{/each}
</ul>

<!-- Example send message form -->
<form
	on:submit|preventDefault={() => {
		sendMessage(
			JSON.stringify({
				type: 'status',
				address: 'broadcast',
				id: 'a',
				source_id: currentUser.id,
				data: 'Hello WebSocket!'
			})
		);
	}}
>
	<button type="submit">Send Message</button>
</form>

<ul>
	{#each Object.entries(statuses) as status (status)}
		<li class="text-sm rounded-md w-fit border bg-sky-500 p-1 m-0.5">{getName(status[0])}</li>
	{/each}
</ul>

<!-- <h1>HERE WILL BE HOME PAGE</h1>
<h1>{currentUser.username}</h1> -->
