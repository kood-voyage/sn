<script lang="ts">
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	import { type CarouselAPI } from '$lib/components/ui/carousel/context.js';
	import * as Carousel from '$lib/components/ui/carousel/index.js';

	let api: CarouselAPI;
	let count = 0;
	let current = 0;

	$: if (api) {
		count = api.scrollSnapList().length;
		current = api.selectedScrollSnap() + 1;
		api.on('select', () => {
			current = api.selectedScrollSnap() + 1;
		});
	}

	// export let article: any;
	export let imgURL: string;
	export let title: string;
	export let body: string;
</script>

<Dialog.Root>
	<Dialog.Trigger class="h-96 w-full sm:rounded-lg mb-1 sm:mb-0 ">
		<div class="p-1">
			<img src={imgURL} alt="image1" />
			<p class="text-lg text-ellipsis w-full text-left">{title}</p>
			<p class="lines3 text-sm text-left text-slate-400">{body}</p>
		</div>
	</Dialog.Trigger>

	<Dialog.Content>
		<Dialog.Header class="">
			<Dialog.Title>{title}</Dialog.Title>

			<Dialog.Description>
				{body}
			</Dialog.Description>

			<Carousel.Root bind:api>
				<Carousel.Content>
					<Carousel.Item><img src={imgURL} alt="image1" /></Carousel.Item>

					>
					<Carousel.Item
						><img
							src="https://resources.finalsite.net/images/v1629478453/usmk12org/yizzw0mr1escg58pmhvp/esports.png"
							alt="image1"
						/></Carousel.Item
					>
				</Carousel.Content>

				<Carousel.Previous class="left-2" />
				<Carousel.Next class="right-2" />
			</Carousel.Root>

			<div class="py-2 text-center text-sm text-muted-foreground">
				Slide {current} of {count}
			</div>
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
