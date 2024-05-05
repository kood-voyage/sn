<script lang="ts">
	import * as Command from '$lib/components/ui/command';
	import Input from '$lib/components/ui/input/input.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import * as Form from '$lib/components/ui/form';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { groupEventSchema } from '$lib/types/group-schema';

	export let data;
	const form = superForm(data, {
		validators: zodClient(groupEventSchema),
		onSubmit: async ({ action, submitter }) => {
			console.log('RAN >>>', action);
			console.log('RAN >>>', submitter);
            console.log("THIS IS THEDATA", data)
		}
	});

	const { form: formData, enhance, submit } = form;
</script>

<form
	method="post"
	action="?/groupEventCreate"
	use:enhance
    >
	<Form.Field {form} name="title">
		<Form.Control let:attrs>
			<Form.Label>Title</Form.Label>
			<Input {...attrs} placeholder="title" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="content">
		<Form.Control let:attrs>
			<Form.Label>Body</Form.Label>

			<Textarea {...attrs} placeholder="body" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button type="submit" class="w-full">Submit</Form.Button>
</form>
