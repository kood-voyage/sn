<script lang="ts">
	import { page } from '$app/stores';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';

	export let Icon;
	export let href;
	export let msg;

	$: active = href.split('/')[2] === $page.url.pathname.split("/")[2];


</script>

<Tooltip.Root>
	<Tooltip.Trigger>
		<a {href}>
			<div class="button {active ? 'bg-primary' : ''}">
				<Icon class="h-[1.2rem] w-[1.2rem] m-auto self-center" />
			</div>
		</a></Tooltip.Trigger
	>
	<Tooltip.Content side="right">
		<p>{msg}</p>
	</Tooltip.Content>
</Tooltip.Root>

<style>
	.button {
		@apply hover:bg-primary hover:dark:bg-primary flex h-[58px] w-[58px] cursor-pointer rounded  transition-all duration-300;
	}
</style>
