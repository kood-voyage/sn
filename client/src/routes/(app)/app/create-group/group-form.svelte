<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { groupSchema, type GroupSchema } from '../group-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { Label } from '$lib/components/ui/label/index.js';

	import * as Carousel from '$lib/components/ui/carousel/index.js';

	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import { Table } from 'svelte-radix';
	import { handleImageCopression } from '$lib/client/image-compression';
	import { handleSubmit } from './logic';
	import { applyAction, deserialize } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import type { ActionResult } from '@sveltejs/kit';

	export let data: SuperValidated<Infer<GroupSchema>, any>;
	$: test = '';
	$: console.log(test);

	$: image = '';

	const form = superForm(data, {
		validators: zodClient(groupSchema),
		onSubmit: (input) => {
			handleSubmit(input.formData);
			input.cancel();
		}
	});

	const { form: formData, enhance } = form;

	async function displayImagePreviews(file: File) {
		console.log(file instanceof File);
		const reader = new FileReader();
		reader.onloadend = (e) => (image = reader.result as string);
		reader.readAsDataURL(file);
	}

	function handleChange(event: Event) {
		const target = event.target as HTMLInputElement;
		if (target && target.files && target.files.length > 0) {
			displayImagePreviews(target.files[0]);
		} else {
			image = '';
		}
	}
</script>

<form method="POST" id="groupForm" use:enhance enctype="multipart/form-data" class="w-full mt-5">
	<!-- Input field for uploading multiple images -->

	<Form.Field {form} name="image">
		<Form.Control let:attrs>
			<Form.Label>Group Image</Form.Label>
			<Input
				type="file"
				accept="image/gif, image/jpeg, image/png, image/webp, image/svg+xml"
				on:change={handleChange}
				{...attrs}
			/>
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<!-- Img Preview -->

	{#if image != ''}
		<img
			id="previewCover"
			src={image}
			alt="previewCover"
			class="w-full my-2 h-16 object-cover rounded"
		/>
	{/if}
	<Form.Field {form} name="title">
		<Form.Control let:attrs>
			<Form.Label>Name</Form.Label>
			<Input {...attrs} bind:value={$formData.title} placeholder="name" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="content">
		<Form.Control let:attrs>
			<Form.Label>About You</Form.Label>

			<Textarea {...attrs} bind:value={$formData.content} placeholder="description" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<RadioGroup.Root class="my-3" value="public">
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="public" id="r1" />
			<Label for="r1">Public</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="private" id="r2" />
			<Label for="r2">Private</Label>
		</div>

		<RadioGroup.Input name="privacy" />
	</RadioGroup.Root>

	<Form.Button class="w-full">Submit</Form.Button>
</form>
