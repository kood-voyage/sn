<script>
	import * as Form from '$lib/components/ui/form';
	import Label from '$lib/components/ui/label/label.svelte';

	import * as RadioGroup from '$lib/components/ui/radio-group';
</script>

<form>
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
			<RadioGroup.Item value="not_going" id="r2" />
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
