<script lang="ts">
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import * as Form from '$lib/components/ui/form';

	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label/index.js';
	import Editor from '$lib/components/Editor.svelte';

	import { postSchema, type PostSchema } from '../post-schema';
	import SuperDebug, { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';

	import { v4 as uuidv4 } from 'uuid';
	import { Title } from '$lib/components/ui/dialog';
	import { browser } from '$app/environment';

	export let data: SuperValidated<Infer<PostSchema>>;

	let files;

	const form = superForm(data, {
		validators: zodClient(postSchema),
		onSubmit: ({ formData }) => {
			const imageFormData = new FormData();

			const post_id = uuidv4();

			for (const image of files) {
				imageFormData.append('images', image);
			}

			imageFormData.append('path', `post/${post_id}`);

			async function imageStore(formData) {
				const fetchResp = await fetch(PUBLIC_LOCAL_PATH + `/api/v1/auth/images/${post_id}`, {
					method: 'POST',
					headers: {
						'Access-Control-Request-Method': 'POST'
					},
					credentials: 'include',
					body: formData
				});
				const json = await fetchResp.json();

				console.log(json);
			}

			imageStore(imageFormData);

			async function createPost() {
				const json = {
					id: post_id,
					title: $formData.title,
					content: $formData.content,
					privacy: $formData.privacy,
					commnity_id: ''
				};

				const resp = await fetch(PUBLIC_LOCAL_PATH + '/api/v1/auth/posts/create', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'Access-Control-Request-Method': 'POST'
					},
					credentials: 'include',
					body: JSON.stringify(json)
				});

				console.log(resp)

				console.log(JSON.stringify(json))
			}

			createPost();
		},

		onError: (event) => {
			console.log('Hello');
			console.log(event);
		}
	});

	const { form: formData, enhance } = form;

	function handleFileChange(event) {
		files = event.target.files;
		$formData.images = files;
	}

	function generateImagePreviews(files) {
		return Array.from(files).map((file) => {
			const objectURL = URL.createObjectURL(file);
			return {
				name: file.name,
				size: file.size,
				preview: objectURL
			};
		});
	}

	$: imagePreviews = files ? generateImagePreviews(files) : [];

	$: console.log($formData);
</script>

{#if imagePreviews}
	<p>Image Preview</p>
	<Carousel.Root class="w-full max-w-xs m-auto">
		<Carousel.Content>
			{#each imagePreviews as file}
				<Carousel.Item>
					<img src={file.preview} alt="Preview" />
				</Carousel.Item>
			{/each}
		</Carousel.Content>
		<Carousel.Previous />
		<Carousel.Next />
	</Carousel.Root>
{/if}

<form method="POST" enctype="multipart/form-data" use:enhance>
	<RadioGroup.Root bind:value={$formData.privacy}>
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
		<div class="border border-neutral-800 p-2 rounded-lg">
			<Editor bind:editorContent={$formData.content} />
		</div>
	</Form.Field>

	<Form.Field {form} name="images">
		<Form.Control let:attrs>
			<Input
				required
				accept="image/gif, image/jpeg, image/png, image/webp"
				on:change={handleFileChange}
				type="file"
				multiple
				{...attrs}
			/>
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button class="w-full">Submit</Form.Button>

	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
