<script lang="ts">
	import * as Command from '$lib/components/ui/command';
	import { page } from '$app/stores';
	import type { UserType } from '$lib/types/user';
	import { InviteToGroup, type InviteJson } from '$lib/client/api/group-requests';

	const allUsers = $page.data.allusers as UserType[];
	export let invitedUsers: string[] | undefined;
	export let userList: UserType[];
	export let groupid: string;

	function inviteUser(event: Event) {
		let user = findUserId(event.target.innerHTML);
		// console.log('user', user);
		let invite: InviteJson = {
			group_id: groupid,
			target_id: user,
			message: "have invited you to a group"
		}
		const resp = InviteToGroup(invite)
	}

	function isInGroup(user: UserType) {
		return userList.some((groupUser) => groupUser.id === user.id);
	}

	function isInvited(userId: string) {
		if (invitedUsers) {
			return invitedUsers.some((user) => {
				if (user === userId) {
					return true;
				}
				return false;
			});
		}
		return false;
	}

	function findUserId(userName: string): string {
		const foundUser = allUsers.find((user) => user.username === userName);
		if (foundUser) {
			return foundUser.id;
		}
		return '';
	}

	const filteredUsers = allUsers.filter((user) => !isInGroup(user) && !isInvited(user.id!));
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
