<script lang="ts">
	export let postAuthor;
	export let created_at;

	import Calendar from 'svelte-radix/Calendar.svelte';
	import * as HoverCard from '$lib/components/ui/hover-card/index.js';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	import { Globe } from 'svelte-radix';

	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');

</script>

<div class="flex h-12 justify-between p-4">
	<HoverCard.Root>
		<HoverCard.Trigger
			target="_blank"
			rel="noreferrer noopener"
			class="rounded-sm underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black h-full"
		>
			<a href="/app/u/{postAuthor.username}" class="flex h-full self-center items-center">
				<img src={postAuthor.avatar} alt="avatarURL" class="w-8 h-8 object-cover rounded-full" />
				<div class="ml-2 items-center self-center">
					<p class="text-sm">{postAuthor.username}</p>
					<p class="text-xs font-thin">{postAuthor.first_name} {postAuthor.last_name}</p>
				</div>
			</a>
		</HoverCard.Trigger>

		<HoverCard.Content class="w-80">
			<div class="flex flex-col w-full">
				<!-- <div class="w-full h-16 mb-2">
					<img src={postAuthor.cover} alt="cover" class="w-full h-16 object-cover rounded-sm" />
				</div> -->

				<div class="flex">
					<img
						src={postAuthor.avatar}
						alt="avatarURL"
						class="w-16 h-16 aspect-square object-cover rounded-full mr-2"
					/>
					<div class="space-y-1 p-0">
						<h4 class="text-md font-semibold">@{postAuthor.username}</h4>
						<p class="text-sm font-thin">{postAuthor.first_name} {postAuthor.last_name}</p>
						<p class="text-sm text-muted-foreground">"{postAuthor.description}"</p>
						<div class="flex items-center pt-2">
							<Calendar class="mr-2 h-4 w-4 opacity-70" />
							<span class="text-xs text-muted-foreground"
								>Joined {timeAgo.format(new Date(postAuthor.timestamp))}</span
							>
						</div>
					</div>
				</div>
			</div>
		</HoverCard.Content>
	</HoverCard.Root>

	<div class="flex">
		<Tooltip.Root>
			<Tooltip.Trigger>
				<div class="flex">
					<p class="text-muted-foreground text-sm items-center self-center flex">
						{timeAgo.format(new Date(created_at), 'mini')}
					</p>
				</div>
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p class="z-[99999]">{new Date(created_at)}</p>
			</Tooltip.Content>
		</Tooltip.Root>

		<Tooltip.Root>
			<Tooltip.Trigger>
				<div class="flex">
					<Globe class="w-5 h-5 ml-1 p-[3px] items-center self-center flex text-muted-foreground" />
				</div>
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p>{'public'}</p>
			</Tooltip.Content>
		</Tooltip.Root>
	</div>
</div>
