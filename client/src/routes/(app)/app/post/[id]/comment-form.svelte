<script lang="ts">
	import Editorsn from '$lib/components/Editorsn.svelte';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { commentSchema } from '../../comment-schema';
	import { superForm } from 'sveltekit-superforms';

	export let data;
	export let post_id;

	let childValue = '';

	const form = superForm(data, {
		validators: zodClient(commentSchema),
		onSubmit: ({ formData }) => {
			formData.set('content', childValue);
			formData.set('post_id', post_id);
			formData.set('content', childValue);
		}
	});

	const { form: formData, enhance } = form;

	function handleChildValue(value: any) {
		childValue = value.detail.innerHTML;
	}

</script>

<form
	method="POST"
	action="?/commentSubmit"
	enctype="multipart/form-data"
	use:enhance
	class="w-full"
>
	<Form.Field {form} name="content">
		<Form.Control let:attrs>
			<Editorsn {...attrs} on:valueChange={handleChildValue} />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button class="w-full">Submit</Form.Button>
</form>

<!-- Img Preview -->

<!-- Input field for uploading multiple images -->

<!-- <Form.Field {form} name="images">
		<Form.Control let:attrs>
			<Form.Label>Images (up to 3)</Form.Label>
			<Input
				type="file"
				required
				accept="image/gif, image/jpeg, image/png, image/webp"
				multiple
				on:change={(e) => limitFiles(e.target.files, 3)}
				{...attrs}
			/>
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field> -->
