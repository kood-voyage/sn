<script lang="ts">
	import Calendar from 'svelte-radix/Calendar.svelte';
	import * as HoverCard from '$lib/components/ui/hover-card/index.js';
	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	import { toast } from 'svelte-sonner';



    



    
	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');

	export let status;
</script>

<HoverCard.Root>
	<HoverCard.Trigger
		target="_blank"
		rel="noreferrer noopener"
		class="rounded-sm underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black h-full"
	>
		<a
			href="/app/u/{status.source_information.username}"
			class="flex h-full self-center items-center"
		>
			<p class="text-sm">{status.source_information.username}</p>
		</a>
	</HoverCard.Trigger>

	<HoverCard.Content class="w-80">
		<div class="flex flex-col w-full">
			<!-- <div class="w-full h-16 mb-2">
					<img src={postAuthor.cover} alt="cover" class="w-full h-16 object-cover rounded-sm" />
				</div> -->

			<div class="flex">
				<img
					src={status.source_information.avatar}
					alt="avatarURL"
					class="w-16 h-16 aspect-square object-cover rounded-full mr-2"
				/>
				<div class="space-y-1 p-0">
					<h4 class="text-md font-semibold">@{status.source_information.username}</h4>
					<p class="text-sm font-thin">
						{status.source_information.first_name}
						{status.source_information.last_name}
					</p>
					<p class="text-sm text-muted-foreground">"{status.source_information.description}"</p>
					<div class="flex items-center pt-2">
						<Calendar class="mr-2 h-4 w-4 opacity-70" />
						<span class="text-xs text-muted-foreground"
							>Joined {timeAgo.format(new Date(status.source_information.timestamp))}</span
						>
					</div>
				</div>
			</div>
		</div>
	</HoverCard.Content>
</HoverCard.Root>
