<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	import { onMount } from 'svelte';

	export let comments;

	onMount(() => {
		TimeAgo.addDefaultLocale(en);
	});

	const timeAgo = new TimeAgo('en-US');
</script>

<div class="h-full px-4">
	{#if comments && comments.length > 0}
		{#each comments as comment}
			<div class="flex p-2">
				<img src={comment.user_avatar} alt="avatar" class="w-9 h-9 object-cover rounded-full" />

				<div class="flex flex-col">
					<div class="flex flex-col p-2 mx-2 bg-neutral-200 dark:bg-neutral-700 rounded-xl">
						<p class="text-[10px]">
							{comment.user_name}
						</p>
						<div class="w-full text-md">{@html comment.content}</div>
					</div>

					<Tooltip.Root>
						<Tooltip.Trigger
							class="rounded-sm underline-offset-4 hover:underline focus-visible:outline-2 focus-visible:outline-offset-8 focus-visible:outline-black ml-4"
						>
							<div class="flex">
								<p class="text-muted-foreground text-sm items-center self-center flex">
									{timeAgo.format(new Date(comment.created_at), 'mini')}
								</p>
							</div>
						</Tooltip.Trigger>
						<Tooltip.Content>
							<p class="z-[99999]">{new Date(comment.created_at)}</p>
						</Tooltip.Content>
					</Tooltip.Root>
				</div>
			</div>
		{/each}
	{:else}
		<div class="flex items-center h-full justify-center">No comments yet...</div>
	{/if}
</div>
