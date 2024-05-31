<script lang="ts">
	import { userStatusStore } from '$lib/store/websocket-store';
	import { onMount } from 'svelte';

	export let user;

	onMount(() => {
		userStatusStore.subscribe((user_status) => {
			userStatus = user_status;
		});
	});

	let userStatus = $userStatusStore;

	console.log(user);
</script>

<div class="relative">
	<img src={user.avatar} alt="avatar" class="m-auto sm:mx-2 w-8 h-8 rounded-full" />

	<div
		class="absolute rounded-full w-3 h-3 top-5 left-7 z-50 {userStatus[user.id]
			? 'bg-green-500'
			: 'bg-red-500'} border-2"
	></div>
</div>

<p class="h-fit align-middle justify-center text-center text-sm hidden sm:block">
	{user.username}
</p>
