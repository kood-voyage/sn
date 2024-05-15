<script lang="ts">
	import { sendMessage, sendNotification } from '$lib/client/websocket';
	import {
		messageStore,
		notificationStore,
		userStatusStore,
		type MessageStore,
		type NotificationStore,
		type UserStatusStore
	} from '$lib/store/websocket-store';
	import { currentUserStore } from '$lib/store/user-store';
	import type { UserType } from '$lib/types/user';
	import { onDestroy } from 'svelte';
	import type { PageData } from './$types';
	import { GetAllUsers } from '$lib/client/api/user-requests';
	import {
		createNotification,
		deleteNotification,
		getUserNotifications
	} from '$lib/client/api/notification-requests';
	import toast from 'svelte-french-toast';

	export let data: PageData;

	let messages: MessageStore = [];
	let notifications: NotificationStore = [];
	let statuses: UserStatusStore = {};
	let currentUser = $currentUserStore as UserType;
	$: display_statuses = Object.entries(statuses).filter(([_, bool]) => bool);

	messageStore.subscribe((value) => {
		if (value) messages = value;
	});

	userStatusStore.subscribe((value) => {
		if (value) statuses = value;
	});

	notificationStore.subscribe((value) => {
		if (value) notifications = value;
	});

	let allUsers: UserType[] = data.allUsers;

	// $: for (const [key, value] of Object.entries(statuses)) {
	// 	if value
	// }
	console.log(statuses['asdasd']);

	const getName = (user_id: string) => {
		const user = allUsers.find((val) => {
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
	on:submit|preventDefault={async () => {
		const userNotifResp = await getUserNotifications();
		console.log(userNotifResp);
		if (!userNotifResp.ok) {
			toast.error(userNotifResp.message);
			return;
		}
		console.log(userNotifResp.notifications);
		if (userNotifResp.notifications) notificationStore.set(userNotifResp.notifications);
	}}
>
	<button type="submit">Send Message</button>
</form>

<ul class="border-b-2 border-neutral-300 dark:border-neutral-950">
	{#if allUsers != undefined && allUsers.length != 0}
		{#each display_statuses as status (status)}
			<li class="text-sm rounded-md w-fit border bg-sky-500 p-1 m-0.5 cursor-pointer">
				<button
					on:click={async () => {
						const notifResp = await createNotification(
							currentUser.id,
							status[0],
							`HEy first notification be like... ${notifications.length + 1}`
						);
						if (!notifResp.ok) {
							toast.error(notifResp.message);
							return;
						}
						console.log(notifResp);
						sendNotification(status[0], currentUser.id, notifResp.createdNotif);
					}}
				>
					{getName(status[0])}
				</button>
			</li>
		{/each}
	{/if}
</ul>
<ul class="border-b-2 border-neutral-300 dark:border-neutral-950">
	{#if notifications != undefined && notifications.length != 0}
		{#each notifications as status (status)}
			<li class="text-sm rounded-md w-fit border bg-sky-500 p-1 m-0.5">
				<button
					on:click={async () => {
						const notifResp = await deleteNotification(status.id);
						console.log('notifResp >>>', notifResp);
						if (notifResp.ok)
							notificationStore.update((old) => old.filter((notif) => notif.id != status.id));
					}}
				>
					{status.message}
				</button>
			</li>
		{/each}
	{/if}
</ul>
<!-- <h1>HERE WILL BE HOME PAGE</h1>
<h1>{currentUser.username}</h1> -->
