<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import { page } from '$app/stores';

	export let comments;

	import en from 'javascript-time-ago/locale/en';
	import TimeAgo from 'javascript-time-ago';
	TimeAgo.addDefaultLocale(en);
	const timeAgo = new TimeAgo('en-US');

	let previousUserId = comments[0].user_id;

	// console.log($page.data.data.id);



    console.loh
</script>

<div class="h-full px-4">
	{#each comments as comment}
		<div class="flex p-2">
			{#if $page.data.data.id === comment.user_id}
				<img
					src={$page.data.data.avatar}
					alt="avatar"
					class="w-9 h-9 object-cover rounded-full"
				/>
			{:else}
				<img
					src="https://sm.ign.com/ign_nordic/cover/a/avatar-gen/avatar-generations_prsz.jpg"
					alt="avatar"
					class="w-9 h-9 object-cover rounded-full"
				/>
			{/if}

			<div class="flex flex-col">
				<div class="flex flex-col p-2 mx-2 bg-neutral-200 dark:bg-neutral-700 rounded-2xl">
					<p class="text-[10px]">{$page.data.data.id === comment.user_id ? $page.data.data.username :comment.user_id}</p>
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
</div>
