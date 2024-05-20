<script lang="ts">
	import { getUserById } from '$lib/client/api/user-requests';
	import { onMount } from 'svelte';
	import Follower from './Follower.svelte';
	import TimeAgo from 'javascript-time-ago';
	import en from 'javascript-time-ago/locale/en';

	export let user_id;
	export let message;

	let promise = getUserById(user_id);

	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');
</script>

<a
	href={`/app/chat/${message.chat_id}`}
	class="flex items-center p-2 border hover:bg-neutral-100 hover:dark:bg-neutral-800 dark:border-neutral-800 rounded my-1 cursor-pointer mx-auto w-full max-w-96"
>
	<div class="w-full max-w-96">
		<div class="flex justify-between w-full">
			{#await promise}
				load..
			{:then user}
				<div class="flex justify-between items-center gap-2 w-full">
					<Follower user={user.data} avatar={true} />

					<p class="text-xs">send you a message</p>
					<p class="text-sm">{timeAgo.format(new Date(message.created_at), 'mini')}</p>
				</div>
			{/await}
		</div>

		

		<div class="p-2 dark:bg-neutral-900 rounded-lg w-full max-w-96">{@html message.message}</div>
	</div>
</a>
