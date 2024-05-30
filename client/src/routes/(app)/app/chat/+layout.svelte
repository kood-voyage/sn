<script lang="ts">
	import type { UserType } from '$lib/types/user';
	import type { LayoutData } from './$types';
	import type { ChatsWithUsers } from './+layout';
	import Userlist from './userlist.svelte';

	export let data: LayoutData;

	let users: UserType[] = [];
	let chats: ChatsWithUsers = {};
	$: if (data.ok) {
		users = data.chatLoadData.usersData;
		chats = data.chatLoadData.chatsData;
	}
</script>

<div class="flex h-screen w-full">
	<Userlist {users} {chats} />

	<slot />
</div>
