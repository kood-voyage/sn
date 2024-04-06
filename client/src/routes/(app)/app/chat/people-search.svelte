<script lang="ts">
	import type { UserRowType } from '$lib/server/db/user';
	import type { User } from '$lib/types/user';

	type SampleData = {
		user_id: string;
		username: string;
	};
	// {{{

	let people: UserRowType[] | SampleData[] = [
		{ user_id: '1', username: 'Alice Smith' },
		{ user_id: '2', username: 'Bob Johnson' },
		{ user_id: '3', username: 'Charlie Brown' },
		{ user_id: '4', username: 'Dana White' },
		{ user_id: '5', username: 'Evan Green' },
		{ user_id: '6', username: 'Fiona Grey' },
		{ user_id: '7', username: 'George Black' },
		{ user_id: '8', username: 'Hannah Blue' },
		{ user_id: '9', username: 'Ian Yellow' },
		{ user_id: '10', username: 'Jenna Purple' },
		{ user_id: '11', username: 'Karl Orange' },
		{ user_id: '12', username: 'Lena Pink' },
		{ user_id: '13', username: 'Mike Red' },
		{ user_id: '14', username: 'Nina Violet' },
		{ user_id: '15', username: 'Oscar Indigo' },
		{ user_id: '16', username: 'Paula Brown' },
		{ user_id: '17', username: 'Quentin Maroon' },
		{ user_id: '18', username: 'Rachel Olive' },
		{ user_id: '19', username: 'Steven Cyan' },
		{ user_id: '20', username: 'Tina Magenta' },
		{ user_id: '21', username: 'Ursula Lavender' },
		{ user_id: '22', username: 'Victor Cream' },
		{ user_id: '23', username: 'Wendy Peach' },
		{ user_id: '24', username: 'Xander Lime' },
		{ user_id: '25', username: 'Yolanda Mint' },
		{ user_id: '26', username: 'Zachary Emerald' },
		{ user_id: '27', username: 'Amelia Ruby' },
		{ user_id: '28', username: 'Noah Sapphire' },
		{ user_id: '29', username: 'Isabella Jade' },
		{ user_id: '30', username: 'Liam Coral' }
	];
	/// }}}
	let searchQuery = '';

	export let userInfo: UserRowType[] | undefined;

	if (userInfo != undefined) people = userInfo;

	// Computed property that filters the list based on the search query
	$: filteredPeople = people
		.filter((person) => person.username.toLowerCase().includes(searchQuery.toLowerCase()))
		.slice(0, 6);
</script>

<div class="mx-4">
	<input type="text" placeholder="Search people..." class="search-box" bind:value={searchQuery} />

	<!-- <ul class="person-list"> -->

	<ol class="">
		{#each filteredPeople as person (person)}
			<!-- <img src={person.avatar} alt="avatar" class="m-auto sm:mx-2" /> -->
			<li class="user hover:bg-slate-300 dark:hover:bg-slate-800 text-center">
				<img src="/static/favicon.png" alt="avatar" class="m-auto sm:mx-2" />
				<p class=" align-middle justify-center text-center">{person.username}</p>
			</li>
			<!-- <li class="person-item select-none text-center hover:bg-slate-500">{person.username}</li> -->
		{/each}
	</ol>
	<!-- </ul> -->
</div>

<style>
	.search-box {
		padding: 6px;
		margin-bottom: 20px;
		width: 100%;
		box-sizing: border-box;
	}
	.person-list {
		list-style-type: none;
		padding: 0;
	}
	.person-item {
		padding: 5px;
		margin: 5px;
		border-radius: 0.5rem;
		border: 1px solid #ccc;
	}
	div ol li.user {
		@apply flex h-12 cursor-pointer self-center rounded-sm;
	}

	div ol li.user p {
		@apply hidden items-center text-center text-slate-700 dark:text-slate-300;
	}

	@media (min-width: 640px) {
		div ol li.user p {
			@apply ml-2 flex items-center text-center align-middle;
		}
	}

	div ol li.user img {
		@apply h-8 w-8 items-center justify-center rounded-full;
	}

	.active {
		@apply bg-slate-400 dark:bg-slate-700;
	}

	.user .active p {
		@apply text-slate-900 dark:text-slate-100;
	}
</style>
