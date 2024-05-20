<script lang="ts">
	import Calendar from 'svelte-radix/Calendar.svelte';
	import * as HoverCard from '$lib/components/ui/hover-card/index.js';
	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	export let user;
	export let avatar;

	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');
</script>

<HoverCard.Root>
	<HoverCard.Trigger
		target="_blank"
		rel="noreferrer noopener"
		class="rounded-sm underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black h-full"
	>
		<a href="/app/u/{user.username}" class="flex h-full self-center items-center">
			{#if avatar}
				<img
					src={user.avatar}
					alt="avatarURL"
					class="w-8 h-8 aspect-square object-cover rounded-full mr-2"
				/>
			{/if}
			<p class="text-sm">{user.username}</p>
		</a>
	</HoverCard.Trigger>

	<HoverCard.Content class="w-80">
		<div class="flex flex-col w-full">
			<div class="flex">
				<img
					src={user.avatar}
					alt="avatarURL"
					class="w-16 h-16 aspect-square object-cover rounded-full mr-2"
				/>
				<div class="space-y-1 p-0">
					<h4 class="text-md font-semibold">@{user.username}</h4>
					<p class="text-sm font-thin">
						{user.first_name}
						{user.last_name}
					</p>
					<p class="text-sm text-muted-foreground">"{user.description}"</p>
					<div class="flex items-center pt-2">
						<Calendar class="mr-2 h-4 w-4 opacity-70" />
						<span class="text-xs text-muted-foreground"
							>Joined {timeAgo.format(new Date(user.timestamp))}</span
						>
					</div>
				</div>
			</div>
		</div>
	</HoverCard.Content>
</HoverCard.Root>
