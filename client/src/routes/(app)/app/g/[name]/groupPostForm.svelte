<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { postSchema, type PostSchema } from '$lib/types/post-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import * as Carousel from '$lib/components/ui/carousel/index.js';

	import { createPostImagesStore } from '$lib/store/create-post-store';
	import { z } from 'zod';

	// I made schema here, i think it's more appropriate
	export const groupPostSchema = z.object({
		groupId: z.string(),
		title: z.string(),
		content: z.string().min(8),
		images: z
			.array(z.any()) // Temporarily treat images as an array of any type
			.max(3, { message: 'You can only upload up to 3 images.' })
			.optional()
	});

	export let data;
	export let groupId: string;
	console.log('THIS IS THE DATA', groupId);

	let images = [];

	const form = superForm(data, {
		validators: zodClient(groupPostSchema)

		// all of this looks to coomplicated for me :(
		// onSubmit: async (input) => {
		// 	console.log('THIS IS ANOTHER TEST');
		// 	const image = input.formData.get('image') as File;

		// 	input.formData.set('groupId', groupId); /// i dont't know about this but i want to store current groupId to a groupPostSchema

		// 	// const imgResp = await handleImageCopression(image);
		// 	// if (!imgResp.ok) {
		// 	// 	input.cancel();
		// 	// 	return;
		// 	// }
		// 	// const file = imgResp.file as File;

		// 	input.formData.set('image', image);
		// 	// console.log(`compressedFile size ${file.size / 1024 / 1024} MB`);
		// }
	});

	const { form: formData, enhance } = form;

	function limitFiles(files, maxFiles) {
		images = Array.from(files);

		if (images.length > maxFiles) {
			alert('You can only select up to 3 images.');
			images = images.slice(0, maxFiles); // Limit the images array to the first 3 images
		} else {
			displayImagePreviews();
		}
	}

	function displayImagePreviews() {
		const updatedImages = [];

		images.forEach((image) => {
			const reader = new FileReader();
			reader.onload = (e) => {
				updatedImages.push(e.target.result);
			};
			reader.readAsDataURL(image);
		});

		createPostImagesStore.set(updatedImages);
	}
</script>

{#if $createPostImagesStore.length > 0}
	<Carousel.Root class="w-full max-w-xs m-auto">
		<Carousel.Content>
			{#each $createPostImagesStore as $image}
				<Carousel.Item>
					<img src={$image} alt="preview" />
				</Carousel.Item>
			{/each}
		</Carousel.Content>
		<Carousel.Previous />
		<Carousel.Next />
	</Carousel.Root>
{/if}

<form
	method="POST"
	action="?/groupPostSubmit"
	enctype="multipart/form-data"
	use:enhance
	class="w-full mt-10"
>
	<Form.Field {form} name="title">
		<Form.Control let:attrs>
			<Form.Label>Title</Form.Label>
			<Input {...attrs} bind:value={$formData.title} placeholder="title" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="content">
		<Form.Control let:attrs>
			<Form.Label>Body</Form.Label>

			<Textarea {...attrs} bind:value={$formData.content} placeholder="body" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<!-- Img Preview -->

	<!-- Input field for uploading multiple images -->

	<Form.Field {form} name="images">
		<Form.Control let:attrs>
			<Form.Label>Images (up to 3)</Form.Label>
			<Input
				type="file"
				accept="image/gif, image/jpeg, image/png, image/webp"
				multiple
				on:change={(e) => limitFiles(e.target.files, 3)}
				{...attrs}
			/>
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button class="w-full">Submit</Form.Button>
</form>
