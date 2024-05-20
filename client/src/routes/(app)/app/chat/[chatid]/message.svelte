<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	import { onMount } from 'svelte';

	export let user;
	export let msg;
	export let time;

	// onMount(() => {
	// 	TimeAgo.addDefaultLocale(en);

	// });

	// const timeAgo = new TimeAgo('en-US');
</script>

<div class="w-full flex mb-4 hover:bg-neutral-200/50 dark:hover:bg-neutral-900/50 p-1">
	<img src={user.avatar} class="w-8 h-8 rounded-full mx-2" alt={user.username + 'image'} />

	<div class="flex flex-col">
		<div class="flex gap-2">
			<p class=" text-xs font-bold">{user.username}</p>
			<Tooltip.Root>
				<Tooltip.Trigger
					class="rounded-sm underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black ml-1"
				>
					<div class="flex">
						<p class="text-muted-foreground text-[11px] items-center self-center flex">
							{new Date(time).toDateString()}
						</p>
					</div>
				</Tooltip.Trigger>
				<Tooltip.Content>
					<p class="z-[99999]">{new Date(time)}</p>
				</Tooltip.Content>
			</Tooltip.Root>
		</div>

		<div contenteditable="false" class="text-sm py-2 px-1 font-light">{@html msg}</div>
	</div>
</div>
