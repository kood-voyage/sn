<script lang="ts">
	import {
		AttendGroupEvent,
		type GroupEventJson,
	} from '$lib/client/api/group-requests';
	import * as Form from '$lib/components/ui/form';
	import Label from '$lib/components/ui/label/label.svelte';

	import * as RadioGroup from '$lib/components/ui/radio-group';
	import { createEventDispatcher } from 'svelte';

	export let eventInfo: GroupEventJson;
	const dispatch = createEventDispatcher();

	async function attendEvent(event: SubmitEvent) {
		event.preventDefault();

		console.log("THIS IS HTE EVENT FKINC INGOF", eventInfo)
		const formData = new FormData(event.target);

		const selectedValue = formData.get('event_selection');

		if (selectedValue) {
			const resp = await AttendGroupEvent(eventInfo.id, selectedValue.toString());
			if (!resp.ok) {
				console.log(resp);
				alert('Something went wrong');
				return;
			}
		}
		dispatch('submit', { detail: 'Created reaction for event' });
		
		console.log('Selected value:', selectedValue);
	}
	// const { form: formData, enhance, submit } = form;
</script>

<form on:submit|stopPropagation={attendEvent}>
	<RadioGroup.Root class="my-3" value="reactform">
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="going" id="r1" />
			<Label for="r1">Going</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="interested" id="r2" />
			<Label for="r2">Interested</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="notgoing" id="r2" />
			<Label for="r3">Not going</Label>
		</div>
		<div class="flex items-center space-x-2">
			<RadioGroup.Item value="maybe" id="r2" />
			<Label for="r4">Maybe</Label>
		</div>
		<RadioGroup.Input name="event_selection" />
	</RadioGroup.Root>
	<Form.Button type="submit" class="w-full">Submit</Form.Button>
</form>
