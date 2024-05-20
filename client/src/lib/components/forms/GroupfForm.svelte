<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { groupSchema, type GroupSchema } from '$lib/types/group-schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { Label } from '$lib/components/ui/label/index.js';

	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import { handleImageCopression } from '$lib/client/image-compression';

	export let data: SuperValidated<Infer<GroupSchema>, any>;

	import { v4 as uuidv4 } from 'uuid';
	import { PUBLIC_LOCAL_PATH } from '$env/static/public';
	import RedStar from '../../../routes/(auth)/signup/RedStar.svelte';
	import { toast } from 'svelte-sonner';

	$: image = '';

	const form = superForm(data, {
		validators: zodClient(groupSchema),
		onSubmit: async ({ controller }) => {
			controller.abort();

			const imageFormData = new FormData();

			const community_id = uuidv4();

			const imgResp = await handleImageCopression($formData.image);
			if (!imgResp.ok) {
				return;
			}

			const file = imgResp.file as File;

			imageFormData.append('images', file);

			imageFormData.append('path', `group/${community_id}`);

			async function createGroup() {
				const json = {
					id: community_id,
					name: $formData.name,
					description: $formData.description,
					privacy: $formData.privacy
				};

				const resp = await fetch(PUBLIC_LOCAL_PATH + '/api/v1/auth/group/create', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'Access-Control-Request-Method': 'POST'
					},
					credentials: 'include',
					body: JSON.stringify(json)
				});

				return resp;
			}

			async function imageStore(formData) {
				const fetchResp = await fetch(
					PUBLIC_LOCAL_PATH + `/api/v1/auth/images/${community_id}/default`,
					{
						method: 'POST',
						headers: {
							'Access-Control-Request-Method': 'POST'
						},
						credentials: 'include',
						body: formData
					}
				);
				await fetchResp.json();
			}

			const resp = await createGroup();
			console.log(await resp.json());

			await imageStore(imageFormData);

			toast.success(`Group ${$formData.name} has been created`);
		}
	});

	const { form: formData, enhance } = form;

	async function displayImagePreviews(file: File) {
		$formData.image = file;
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

		$formData.image = target.files[0];
	}
</script>

<form method="POST" id="groupForm" use:enhance enctype="multipart/form-data" class="w-full mt-5">
	<!-- Input field for uploading multiple images -->

	<Form.Field {form} name="image">
		<Form.Control let:attrs>
			<Form.Label>Cover image<RedStar /></Form.Label>
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
	<Form.Field {form} name="name">
		<Form.Control let:attrs>
			<Form.Label>Name <RedStar /></Form.Label>
			<Input {...attrs} bind:value={$formData.name} placeholder="name" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="description">
		<Form.Control let:attrs>
			<Form.Label>Description</Form.Label>

			<Textarea {...attrs} bind:value={$formData.description} placeholder="description" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="privacy">
		<Form.Control let:attrs>
			<Form.Label>Privacy<RedStar /></Form.Label>
			<RadioGroup.Root bind:value={$formData.privacy} {...attrs}>
				<div class="flex items-center space-x-2">
					<RadioGroup.Item value="public" id="r1" />
					<Label for="r1">Public</Label>
				</div>
				<div class="flex items-center space-x-2">
					<RadioGroup.Item value="private" id="r2" />
					<Label for="r2">Private</Label>
				</div>

				<RadioGroup.Input name="privacy" />
			</RadioGroup.Root></Form.Control
		>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button class="w-full">Submit</Form.Button>
</form>
