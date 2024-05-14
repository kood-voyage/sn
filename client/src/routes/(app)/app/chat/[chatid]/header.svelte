<script lang="ts">
	import UserChatCard from '$lib/components/chat/UserChatCard.svelte';

	export let user;
</script>

<div class="h-12 bg-neutral-50 dark:bg-neutral-800 px-1 w-full">
	<div class="flex p-2 h-full items-center">
		<UserChatCard {user} />

		<div class="border-l-[1px] py-4 h-full mx-2 dark:border-neutral-700"></div>

		<div class="flex space-x-1">
			<p class="rounded-sm dark:bg-neutral-900 bg-neutral-200 px-1 text-xs font-bold">AKA</p>
			<p class="px-1 text-xs">{user.first_name} {user.last_name}</p>
		</div>
	</div>
</div>
