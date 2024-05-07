<script lang="ts">
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import * as Form from '$lib/components/ui/form';

	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label/index.js';
	import Editor from '$lib/components/Editor.svelte';

	import { postSchema, type PostSchema } from '../post-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { FilterRuleName } from '@aws-sdk/client-s3';

	import { PUBLIC_LOCAL_PATH } from '$env/static/public';
	import { fromTheme } from 'tailwind-merge';

	export let data: SuperValidated<Infer<PostSchema>>;

	let editorContent = '';

	let files;

	const form = superForm(data, {
		validators: zodClient(postSchema),
		onSubmit: ({ formData }) => {
			formData.set('content', editorContent);

			const imageFormData = new FormData();

			const post_id = 'pseudoID';

			imageFormData.append('images', files);

			console.log($formData.privacy);


			imageFormData.append('path', `/post/${post_id}`);

			async function imageStore() {
				const fetchResp = await fetch(PUBLIC_LOCAL_PATH + `/api/v1/auth/images/${post_id}`, {
					method: 'POST',
					headers: {
						'Access-Control-Request-Method': 'POST'
					},
					credentials: 'include',
					body: imageFormData
				});
				const json = await fetchResp.json();

				console.log(json);
			}

			imageStore();
		}
	});

	const { form: formData, enhance } = form;

	function handleFileChange(event) {
		files = event.target.files;
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
		<div class="border border-neutral-800 p-2 rounded-lg">
			<Editor bind:editorContent />
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
</form>
