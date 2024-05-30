<script lang="ts">
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import * as Form from '$lib/components/ui/form';

	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label/index.js';
	import Editor from '$lib/components/Editor.svelte';

	import { postSchema, type PostSchema } from '../../../routes/(app)/app/post-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';

	import { v4 as uuidv4 } from 'uuid';
	import { handleImageCopression } from '$lib/client/image-compression';
	import { goto } from '$app/navigation';
	import { imageStore } from '$lib/client/api/image-requests';
	import { toast } from 'svelte-sonner';

	export let data: SuperValidated<Infer<PostSchema>>;
	export let community_id: string;

	let files: File[];

	const form = superForm(data, {
		validators: zodClient(postSchema),

		onSubmit: async ({ controller }) => {
			controller.abort();
			const imageFormData = new FormData();
			const post_id = uuidv4();

			const valid = (await validateForm()).valid;

			if (!valid) {
				return;
			}

			for (const image of files) {
				imageFormData.append('images', (await handleImageCopression(image)).file as File);
			}

			imageFormData.append('path', `post/${post_id}`);

			async function createPost() {
				const json = {
					id: post_id,
					title: $formData.title,
					content: $formData.content,
					// community_id:
					privacy: $formData.privacy
				};

				if (community_id !== '') {
					json.community_id = community_id;
				}

				const resp = await fetch(PUBLIC_LOCAL_PATH + '/api/v1/auth/posts/create', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'Access-Control-Request-Method': 'POST'
					},
					credentials: 'include',
					body: JSON.stringify(json)
				});
			}


			await createPost();
			await imageStore(imageFormData, post_id);

			toast.success('Post created!');
			goto('/app');
		},

		onResult: ({ result }) => {
			console.log(result.status);
		},

		onUpdate: async ({ form: f }) => {
			if (!f.valid) {
				console.log('Form has errors. Cannot submit.');
				return;
			}
		},

		onError: (event) => {
			console.log(event);
		}
	});

	const { form: formData, enhance, validateForm } = form;

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
</script>

{#if imagePreviews.length !== 0}
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
	<Form.Field {form} name="privacy">
		<Form.Control let:attrs
			><RadioGroup.Root bind:value={$formData.privacy} {...attrs}>
				<div class="flex items-center space-x-2">
					<RadioGroup.Item value="public" id="r1" />
					<Label for="r1">Public</Label>
				</div>
				<div class="flex items-center space-x-2">
					<RadioGroup.Item value="private" id="r2" />
					<Label for="r2">Private</Label>
				</div>

				<!-- {#if community_id === ''}
					<div class="flex items-center space-x-2">
						<RadioGroup.Item value="selected" id="r3" />
						<Label for="r3">Selected</Label>
					</div>
				{/if} -->
				<RadioGroup.Input name="privacy" />
			</RadioGroup.Root></Form.Control
		>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="title">
		<Form.Control let:attrs>
			<Form.Label>Title</Form.Label>
			<Input {...attrs} bind:value={$formData.title} placeholder="title" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="content">
		<Form.Control let:attrs>
			<div class="border border-neutral-800 p-2 rounded-lg">
				<Editor bind:editorContent={$formData.content} />
			</div></Form.Control
		>
		<Form.FieldErrors />
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
