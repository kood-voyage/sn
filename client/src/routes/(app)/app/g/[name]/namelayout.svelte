<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Command from '$lib/components/ui/command';
    import { page } from `$app/stores`
	import type { User } from '$lib/types/user';
	import { LOCAL_PATH } from '$env/static/private';


    const allUsers = $page.data.allusers
	// const users = mainGetAllUsers();
    const groupUsers = $page.data.group.data.members
	function inviteUser() {
		const resp = fetch(`${LOCAL_PATH}/api/v1/auth/group/invite`, {
            headers: 
        })
	}

    function isInGroup(user: User) {
        return groupUsers.some(groupUser => groupUser.id === user.id);
    }
    const filteredUsers = allUsers.filter(user => !isInGroup(user));
    console.log("THIS IS A PAGE", filteredUsers)
</script>

<Command.Root>
	<Command.Input placeholder="Type a command or search..." />
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
