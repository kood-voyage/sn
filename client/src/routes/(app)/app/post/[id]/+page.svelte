<script lang="ts">
	import CarouselItem from '$lib/components/ui/carousel/carousel-item.svelte';
	import CarouselNext from '$lib/components/ui/carousel/carousel-next.svelte';
	import CarouselPrevious from '$lib/components/ui/carousel/carousel-previous.svelte';
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import { ChatBubble, EnterFullScreen, ExitFullScreen } from 'svelte-radix';
	import Author from './author.svelte';
	import Content from './content.svelte';
	import Comments from './comments.svelte';
	import * as Tooltip from '$lib/components/ui/tooltip/index.js';
	import CommentForm from './comment-form.svelte';
	import { commentsStore } from '$lib/store/comments-store';

	export let data;

	const { post, comments } = data;
	let toggle = false;

	$: commentsLength = $commentsStore === null ? 0 : $commentsStore.length;

	console.log(comments);

	$: commentsStore.set(comments);
</script>

<div class="h-screen flex flex-col lg:flex-row dark:bg-neutral-800 overflow-y-scroll">
	<div class="bg-black w-full h-full flex px-16 m-auto relative justify-center">
		{#if post.image_path}
			<Carousel.Root
				class="flex flex-col h-full my-auto justify-center"
				opts={{ loop: true, skipSnaps: true, watchDrag: false, dragThreshold: 0 }}
			>
				<Carousel.Content class="my=-auto h-full">
					{#each post.image_path as image, i (i)}
						<CarouselItem class="my-auto h-full">
							<div class="">
								<img loading="lazy" src={image} class="m-auto" alt={'' + i} />
							</div>
						</CarouselItem>
					{/each}
				</Carousel.Content>
				<CarouselPrevious />
				<CarouselNext />
			</Carousel.Root>
		{:else}
			<p>No images</p>
		{/if}

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

	<div class="w-full lg:w-[480px] lg:overflow-y-scroll flex flex-col {toggle && 'hidden'}">
		<Author postAuthor={post.user_information} created_at={post.created_at} />

		<Content content={post.content} />

		<Tooltip.Root>
			<Tooltip.Trigger class="justify-end h-8 w-full border-b flex pr-4">
				<ChatBubble class="flex items-center text-muted-foreground w-4" />
				<span class="flex items-center text-muted-foreground w-4 ml-2">{commentsLength}</span>
			</Tooltip.Trigger>
			<Tooltip.Content align="center" alignOffset={800} class="flex items-center self-center">
				<p>Comments {commentsLength}</p>
			</Tooltip.Content>
		</Tooltip.Root>

		<div class="h-full w-full lg:max-w-[480px] lg:overflow-y-scroll">
			<Comments />
		</div>

		<footer class="w-full p-4">
			<CommentForm data={data.form} post_id={data.post.id} />
		</footer>
	</div>
</div>
