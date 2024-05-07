<script lang="ts">
	import * as Command from '$lib/components/ui/command';
	import Input from '$lib/components/ui/input/input.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import * as Form from '$lib/components/ui/form';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm, superValidate } from 'sveltekit-superforms';
	import { eventSchema, type EventSchema } from '$lib/types/group-schema';
	import { CreateGroupEvent, type EventJson, type Group, type GroupEventJson, type GroupJson } from '$lib/client/api/group-requests';

	import { v4 as uuidv4 } from "uuid";
	import type { User } from '$lib/types/user';
	import { date } from 'zod';
	import { goto } from '$app/navigation';

	export let data: SuperValidated<Infer<EventSchema>>;
	export let currUser: User
	export let group: GroupJson

	const form = superForm(data, {
		validators: zodClient(eventSchema),
		onSubmit: async ({ formData, cancel, controller }) => {
			const { name, description } = $formData;

			const event: EventJson = {
				id: uuidv4(),
				user_id: currUser.id,
				group_id: group.id,
				name: name,
				description: description,
				date: new Date(Date.now()),
			};

			const resp = await CreateGroupEvent(event);
			if (!resp.ok) {
				console.log(resp)
				alert('Invalid event create stuff');
				controller.abort('Creating and event was unsuccessful');
				return;
			}
			goto(`/app/g/${group.name.replace(" ", "_")}`);

			cancel();
		},
		onError: (event) => {
			// console.log(event);
		}
	});

	const { form: formData, enhance, submit } = form;
</script>

<form method="POST" use:enhance>
	<Form.Field {form} name="name">
		<Form.Control let:attrs>
			<Form.Label>Title</Form.Label>
			<Input {...attrs} bind:value={$formData.name} placeholder="Event name" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="description">
		<Form.Control let:attrs>
			<Form.Label>Body</Form.Label>
			<Input {...attrs} bind:value={$formData.description} placeholder="Description" />

		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button type="submit" class="w-full">Create</Form.Button>
</form>
