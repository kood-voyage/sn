<script lang="ts">
	import * as Command from '$lib/components/ui/command';
	import { page } from '$app/stores';
	import type { User, UserType } from '$lib/types/user';
	import { InviteToGroup, type InviteJson } from '$lib/client/api/group-requests';

	const allUsers = $page.data.allusers as UserType[];
	export let userList: UserType[];
	export let invitedUsers: string[];
	export let groupid: string;

	function getUser(userId: string) {
		const foundUser = allUsers.find((user) => user.username === userId);
		if (foundUser) {
			return foundUser;
		} else {
			return undefined;
		}
	}
	async function inviteUser(event: Event) {
		const user = getUser(event.target.innerHTML);
		console.log('THI IS THE USER', user);
		const invite: InviteJson = {
			group_id: groupid,
			target_id: user?.id!,
			message: 'I invited you to a group',
		};

		const resp = await InviteToGroup(invite);
		if (!resp.ok) {
			console.log(resp);
			alert('Something went wrong');
			return;
		}
		console.log('Invited user successfully');
	}

	function isInGroup(user: UserType) {
		// console.log("userid", user)
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
