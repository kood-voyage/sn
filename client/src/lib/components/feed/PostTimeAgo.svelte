<script lang="ts">
	import TimeAgo from 'javascript-time-ago';
	import en from 'javascript-time-ago/locale/en';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import { Globe } from 'svelte-radix';
	export let created_at
	TimeAgo.addLocale(en);
	const timeAgo = new TimeAgo('en-US');


</script>


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
