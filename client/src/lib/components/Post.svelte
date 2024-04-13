<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js';

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

	export let data;

	const { image_path, title, content, id, created_at } = data;

	function formatDate(isoDateString: string): string {
		const date: Date = new Date(isoDateString);
		const now: Date = new Date();

		const diffTime: number = Math.abs(now.getTime() - date.getTime());
		const diffDays: number = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

		if (date.toDateString() === now.toDateString()) {
			return 'today';
		}

		if (diffDays <= 14) {
			// If it's within two weeks, return the number of days ago
			return `${diffDays} days ago`;
		} else {
			// Otherwise, return the date in the format dd.mm.yyyy
			const formattedDate: string = date.toLocaleDateString('en-GB', {
				day: '2-digit',
				month: '2-digit',
				year: 'numeric'
			});
			return formattedDate;
		}
	}
</script>

<Dialog.Root>
	<Dialog.Trigger class="h-full w-full sm:rounded mb-1 sm:mb-0 aspect-square group hover:shadow-lg">
		<div class="flex w-full h-full relative">
			{#if image_path != null}
				<div class="hover:bg-slate-900/50 w-full h-full absolute"></div>
				<img src={image_path[0]} alt="image1" class="h-full w-full object-cover sm:rounded" />
			{/if}

			<div
				class="absolute w-full bg-slate-100 dark:bg-slate-950 bottom-0 p-2 opacity-0 group-hover:opacity-100 h-8 overflow-hidden group-hover:h-16 transition-all ease-in-out duration-100 rounded-t flex flex-col"
			>
				<p class="text-sm md:text-lg font-medium text-ellipsis w-full text-left line-clamp-1">
					{title}
				</p>
				<p class="text-xs text-ellipsis w-full text-left line-clamp-1 text-right">
					created at {formatDate(created_at)}
				</p>
				<!-- <p class="text-sm text-left text-slate-400">{content}</p> -->
			</div>
		</div>
	</Dialog.Trigger>

	<Dialog.Content>
		<Dialog.Header class="">
			<Dialog.Title>{title}</Dialog.Title>

			<Dialog.Description>
				{@html content}
			</Dialog.Description>

			<Carousel.Root bind:api>
				<Carousel.Content>
					{#each image_path as image}
						<Carousel.Item><img src={image} alt="image1" /></Carousel.Item>
					{/each}
				</Carousel.Content>

				<Carousel.Previous class="left-2" />
				<Carousel.Next class="right-2" />
			</Carousel.Root>

			<div class="py-2 text-center text-sm text-muted-foreground">
				Slide {current} of {count}
			</div>
		</Dialog.Header>
		<div class="w-full h-full"></div>

		<a href={`/app/post/${id}`} class="w-full h-4">to post</a>
	</Dialog.Content>
</Dialog.Root>
