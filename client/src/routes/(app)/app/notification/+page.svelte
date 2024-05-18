<script lang="ts">
	import { sendNotification } from '$lib/client/websocket';
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
	import type { PageData } from './$types';
	import {
		createNotification,
		deleteNotification,
		getUserNotifications
	} from '$lib/client/api/notification-requests';
	import toast from 'svelte-french-toast';
	import { onMount } from 'svelte';

	import Follower from '$lib/components/notification/Follower.svelte';
	import { Cross1, CrossCircled } from 'svelte-radix';

	export let data: PageData;

	import { fly } from 'svelte/transition';

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

	console.log(data.allUsers);

	onMount(async () => {
		const userNotifResp = await getUserNotifications();
		console.log(userNotifResp);
		if (!userNotifResp.ok) {
			toast.error(userNotifResp.message);
			return;
		}
		console.log(userNotifResp.notifications);
		if (userNotifResp.notifications) notificationStore.set(userNotifResp.notifications);
	});

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
<!-- <form
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
</form> -->

<!-- <ul class="border-b-2 border-neutral-300 dark:border-neutral-950">
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
</ul> -->

<div class="w-full flex flex-col h-full">
	<p class="mx-auto">Notifications</p>

	<ul class="mx-auto">
		{#if notifications != undefined && notifications.length != 0}
			{#each notifications as status (status.id)}
				<li transition:fly={{ y: 200, duration: 500 }} class=" w-96">
					<div
						class="flex items-center p-2 border hover:bg-neutral-100 hover:dark:bg-neutral-800 dark:border-neutral-800 rounded w-full my-1 cursor-pointer"
					>
						<span class="m-1">👋</span>

						<div class="w-full relative">
							<div class="flex items-center">
								<Follower {status} />
								<p class="text-xs ml-2">send request to follow you</p>
							</div>
						</div>

						<button
							on:click={async () => {
								const notifResp = await deleteNotification(status.id);
								console.log('notifResp >>>', notifResp);
								if (notifResp.ok)
									notificationStore.update((old) => old.filter((notif) => notif.id != status.id));
							}}
						>
							<Cross1 class="text-neutral-500 w-3" />
						</button>
					</div>
				</li>
			{/each}
		{:else}
			<p class="w-full h-full my-auto">No notifications</p>
		{/if}
	</ul>
</div>
