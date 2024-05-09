<script lang="ts">
	import type { UserType } from '$lib/types/user';
	import type { LayoutData } from './$types';
	import type { ChatsWithUsers } from './+layout';
	import Userlist from './userlist.svelte';

	export let data: LayoutData;

	let people: UserType[] = [];
	let chats: ChatsWithUsers = {};
	$: if (data.ok) {
		people = data.chatLoadData.usersData;
		chats = data.chatLoadData.chatsData;
	}
</script>

<div class="flex w-full h-screen">
	<Userlist {people} {chats} />

	<slot />
</div>
