<script lang="ts">
	// import Editor from 'tailwind-editor';
	// let html = '';

	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { postSchema, type PostSchema } from '../post-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { Label } from '$lib/components/ui/label/index.js';

	import * as Carousel from '$lib/components/ui/carousel/index.js';

	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import { createPostImagesStore } from '$lib/store/create-post-store';

	import Editorsn from '$lib/components/Editorsn.svelte';
	import { onMount } from 'svelte';

	export let data: SuperValidated<Infer<PostSchema>>;

	let images: File[] = [];
	let childValue;

	const form = superForm(data, {
		validators: zodClient(postSchema),
		onSubmit: ({ formData }) => {
			formData.set('content', childValue);
		}
	});

	const { form: formData, enhance } = form;

	function limitFiles(files: File[], maxFiles: number) {
		images = Array.from(files);

		console.log('hello');

		if (images.length > maxFiles) {
			alert('You can only select up to 3 images.');
			images = images.slice(0, maxFiles); // Limit the images array to the first 3 images
		} else {
			displayImagePreviews();
		}
	}

	// setInterval(() => {
	// 	console.log(divElement.childNodes);
	// }, 1000)

	function displayImagePreviews(): void {
		const updatedImages: File[] = [];

		images.forEach((image) => {
			const reader = new FileReader();

			reader.onload = (e) => {
				if (e.target == null) {
					return;
				}

				const result = e.target.result;
				updatedImages.push(result);
			};
			reader.readAsDataURL(image);
		});

		createPostImagesStore.set(updatedImages);
	}

	function handleChildValue(value) {
		childValue = value.detail.innerHTML;

		console.log(childValue);
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
	action="?/postSubmit"
	enctype="multipart/form-data"
	use:enhance
	class="w-full mt-10"
>
	<RadioGroup.Root value="public">
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="public" id="r1" />
			<Label for="r1">Public</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="private" id="r2" />
			<Label for="r2">Private</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="selected" id="r3" />
			<Label for="r3">Selected</Label>
		</div>
		<RadioGroup.Input name="privacy" />
	</RadioGroup.Root>

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

			<Editorsn {...attrs} placeholder="body" on:valueChange={handleChildValue} />
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
