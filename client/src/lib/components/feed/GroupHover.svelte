<script lang="ts">
	import * as HoverCard from '$lib/components/ui/hover-card/index.js';

	import { onMount } from 'svelte';
	import Badge from '../ui/badge/badge.svelte';
	import { GetGroup } from '$lib/client/api/group-requests';

	export let group_name;

	let group;

	onMount(async () => {
		group = await GetGroup(group_name);
	});

	$: console.log(group);
</script>

<HoverCard.Root>
	<HoverCard.Trigger
		target="_blank"
		rel="noreferrer noopener"
		class="ml-auto underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black"
	>
		<a href={`/app/g/${group_name}`} class=" ml-auto cursor-pointer text-sm">{group_name}</a>
	</HoverCard.Trigger>

	<HoverCard.Content class="w-80">
		<div class="flex flex-col w-full">
			<div class="w-full h-16 mb-2">
				<img src={group.group.image_path} alt="cover" class="w-full h-16 object-cover rounded-sm" />
			</div>

			<div class="flex">
				<div class="space-y-1 p-0">
					<div>
						<h4 class="text-md font-semibold">{group_name}</h4>
						<p class="text-[11px]">Members: {group.group.members.length}</p>
					</div>

					<p class="text-sm text-muted-foreground">"{group.group.description}"</p>
				</div>
			</div>
		</div>
	</HoverCard.Content>
</HoverCard.Root>
