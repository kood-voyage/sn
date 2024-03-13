<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
</script>

<Dialog.Root>
	<Dialog.Trigger class="bg-red-500 h-96 w-full sm:rounded-lg mb-1 sm:mb-0 ">
		<div class="p-1">
			<img
				src="https://images.twinkl.co.uk/tw1n/image/private/t_630/u/ux/twinkl-video-games-2_ver_1.jpg"
			/>
			<p>Game industry standarts</p>
			<p class="lines3 text-sm">
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer lobortis erat nec massa
				convallis scelerisque. Suspendisse sed accumsan neque. Sed malesuada, velit nec
				venenatisdasdasdawq
			</p>
		</div>
	</Dialog.Trigger>
	<Dialog.Content class="w-full">
		<Dialog.Header>
			<Dialog.Title>Edit profile</Dialog.Title>
			<Dialog.Description>
				Make changes to your profile here. Click save when you're done.
			</Dialog.Description>
		</Dialog.Header>
		<div class="w-full h-full"></div>
	</Dialog.Content>
</Dialog.Root>

<style>
	.lines3 {
		display: -webkit-box;
		-webkit-line-clamp: 3; /* number of lines to show */
		-webkit-box-orient: vertical;
		overflow: hidden;
	}
</style>
