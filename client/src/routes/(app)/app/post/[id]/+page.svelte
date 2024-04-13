<script lang="ts">
	import CarouselItem from '$lib/components/ui/carousel/carousel-item.svelte';
	import CarouselNext from '$lib/components/ui/carousel/carousel-next.svelte';
	import CarouselPrevious from '$lib/components/ui/carousel/carousel-previous.svelte';
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import { EnterFullScreen, ExitFullScreen } from 'svelte-radix';

	export let data;

	const { post } = data;

	let toggle = false;
</script>

<div class="h-screen flex flex-col lg:flex-row">
	<div class="bg-black w-full h-full flex px-16 m-auto items-center relative justify-center">
		<Carousel.Root
			class="flex flex-col h-full my-auto justify-center"
			opts={{ loop: true, skipSnaps: true, watchDrag: false, dragThreshold: 0 }}
		>
			<Carousel.Content class="my=-auto h-full">
				{#each post.image_path as image, i (i)}
					<CarouselItem class="my-auto h-full">
						<div class="">
							<img src={image} class="m-auto" alt={'' + i} />
						</div>
					</CarouselItem>
				{/each}
			</Carousel.Content>
			<CarouselPrevious />
			<CarouselNext />
		</Carousel.Root>

		<button
			class="absolute right-4 top-4 rounded-full p-2 hover:bg-neutral-800"
			on:click={() => (toggle = !toggle)}
		>
			{#if toggle}
				<ExitFullScreen class="text-white " />
			{:else}
				<EnterFullScreen class="text-white " />
			{/if}
		</button>
	</div>

	<div class="w-full lg:w-[480px] flex flex-col {toggle && 'hidden'}">
		<header class="0 w-full h-96 overflow-y-scroll p-1">{@html post.content}</header>

		<div class="h-full w-full lg:w-[480px] overflow-y-scroll">
			comment section
			<div class="text-2xl w-full">
				<div class=" w-full h-[1000px]"></div>
			</div>
		</div>

		<footer class=" w-full h-40">editor</footer>
	</div>
</div>
