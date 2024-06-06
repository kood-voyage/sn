<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { signInSchema, type SignInSchema } from '../../../routes/(auth)/schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	export let data: SuperValidated<Infer<SignInSchema>>;

	import type { SignIn } from '../../../routes/(auth)/signin/type';
	import { LoginUser } from '$lib/client/api/user-requests';
	import { goto } from '$app/navigation';
	import toast from 'svelte-french-toast';

	const form = superForm(data, {
		validators: zodClient(signInSchema),
		onSubmit: async ({ formData, cancel, controller }) => {
			const { login, password } = $formData;

			if ((await validate('password')) != undefined) {
				cancel();
				return;
			}
			const credentials: SignIn = {
				login,
				password
			};

			const resp = await LoginUser(credentials);

			if (!resp.ok) {
				toast.error('Username or password incorrect!');
				controller.abort('User logging unsuccessful');
				return;
			}

			if (resp.ok) {
				goto('/app');
			}

			cancel();
		}
	});
	const { form: formData, enhance, validate } = form;
</script>

<form method="POST" use:enhance>
	<img src={'power.png'} alt="login" class="p-20" />
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
