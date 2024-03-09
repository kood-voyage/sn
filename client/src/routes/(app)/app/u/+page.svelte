<script lang="ts">
	const list = [
		'Alfred',
		'Bella',
		'Charlie',
		'Diana',
		'Ethan',
		'Fiona',
		'George',
		'Hannah',
		'Isaac',
		'Julia',
		'Kevin',
		'Lily',
		'Mason',
		'Nora',
		'Oliver',
		'Penelope',
		'Quentin',
		'Rachel',
		'Samuel',
		'Tessa',
		'Ulysses',
		'Victoria',
		'Winston',
		'Xena',
		'Yasmine',
		'Zachary',
		'Alice',
		'Bob',
		'Catherine',
		'David',
		'Emily',
		'Frank',
		'Grace',
		'Henry',
		'Ivy',
		'Jack',
		'Katherine',
		'Leo',
		'Mia',
		'Noah',
		'Oscar',
		'Pamela'
	];

	// The array now contains 100 names
</script>

<svelte:head>
	<title>users</title>
</svelte:head>

<div class="p-12">
	<p class="text-md">Suggested for you</p>
	{#each list as name (name)}
		<div class="flex m-auto w-[420px] p-2 hover:bg-slate-200 dark:hover:bg-slate-900 rounded-md">
			<a href="/app/u/{name}" class="flex">
				<img
					src="https://api.dicebear.com/7.x/bottts-neutral/svg?seed={name}"
					alt="user-avatar"
					class="rounded-full w-12 mr-2"
				/>
				<div>
					<p class="font-bold">{name}</p>
					<p class="text-sm text-slate-500">Full Name or description</p>
				</div>
			</a>
			<!-- <div class="flex">
				<button class="text-sm p-1 bg-blue-500">follow</button>
			</div> -->
		</div>
	{/each}
</div>
