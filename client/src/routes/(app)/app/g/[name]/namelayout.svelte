<script lang="ts">
	import * as Command from '$lib/components/ui/command';
	import { page } from '$app/stores';
	import type { User, UserType } from '$lib/types/user';

	const allUsers = $page.data.allusers as UserType[];
	export let userList: UserType[];
	function inviteUser(event: Event) {
		// const resp = fetch(`${LOCAL_PATH}/api/v1/auth/group/invite`, {
		//     headers:
		// })
		// if(event.target){
		// 	console.log("THis is an evet", event.target.innerHTML)
		// }
		console.log("Need to make API Stuff first")
	}

	function isInGroup(user: UserType) {
		return userList.some((groupUser) => groupUser.id === user.id);
	}
	const filteredUsers = allUsers.filter((user) => !isInGroup(user));
</script>

<Command.Root>
	<Command.Input placeholder="Search a user..." />
	<Command.List>
		<Command.Empty>No results found.</Command.Empty>
		<Command.Group heading="Suggestions">
			{#each filteredUsers as s}
				<Command.Item
					><button class="w-full" on:click|preventDefault={inviteUser}>{s.username}</button
					></Command.Item
				>
			{/each}
		</Command.Group>
	</Command.List>
</Command.Root>
