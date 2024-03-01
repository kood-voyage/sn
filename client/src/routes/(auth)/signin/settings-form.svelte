<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { signInSchema, type SignInSchema } from '../schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	export let data: SuperValidated<Infer<SignInSchema>>;

	const form = superForm(data, {
		validators: zodClient(signInSchema)
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" action="?/signin" use:enhance>
	<Form.Field {form} name="login">
		<Form.Control let:attrs>
			<Form.Label>Login</Form.Label>
			<Input {...attrs} bind:value={$formData.login} placeholder="Email/Username" />
		</Form.Control>

		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="password">
		<Form.Control let:attrs>
			<Form.Label>Password</Form.Label>

			<Input {...attrs} type="password" bind:value={$formData.password} placeholder="Password" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Button class="w-full">Submit</Form.Button>
</form>
