<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';

	export let Icon;
	export let href;
	export let msg;
</script>

<Tooltip.Root>
	<Tooltip.Trigger>
		<a {href}>
			<div class="button">
				<Icon class="h-[1.2rem] w-[1.2rem] m-auto self-center " />
			</div>
		</a></Tooltip.Trigger
	>
	<Tooltip.Content side="right">
		<p>{msg}</p>
	</Tooltip.Content>
</Tooltip.Root>

<style>
	.button {
		@apply flex h-[58px] w-[58px] cursor-pointer rounded transition-all duration-300  hover:bg-primary hover:dark:bg-primary;
	}
</style>
