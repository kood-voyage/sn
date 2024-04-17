<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { signUpSchema, type SignUpSchema } from '../schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import RedStar from './red-star.svelte';
	import EyeOpen from 'svelte-radix/EyeOpen.svelte';
	import EyeClosed from 'svelte-radix/EyeClosed.svelte';


	function toogle() {
		isHide = !isHide;
	}

	export let data: SuperValidated<Infer<SignUpSchema>>;


	let isHide: boolean = true;

	const form = superForm(data, {
		validators: zodClient(signUpSchema)
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" action="?/signup" use:enhance>
	<Form.Field {form} name="username">
		<Form.Control let:attrs>
			<Form.Label>Username <RedStar /></Form.Label>
			<Input {...attrs} bind:value={$formData.username} placeholder="username" />
		</Form.Control>
		<Form.Description class="rounded bg-secondary p-1 text-xs text-sky-500"
			>This is your public display name.</Form.Description
		>
		<Form.FieldErrors />
	</Form.Field>

	<Form.Field {form} name="email">
		<Form.Control let:attrs>
			<Form.Label>Email <RedStar /></Form.Label>
			<Input {...attrs} bind:value={$formData.email} placeholder="e-mail" />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<div class="grid grid-cols-2 gap-8">
		<Form.Field {form} name="firstName">
			<Form.Control let:attrs>
				<Form.Label>First Name <RedStar /></Form.Label>
				<Input {...attrs} bind:value={$formData.firstName} placeholder="first name" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>

		<Form.Field {form} name="lastName">
			<Form.Control let:attrs>
				<Form.Label>Last Name <RedStar /></Form.Label>
				<Input {...attrs} bind:value={$formData.lastName} placeholder="last name" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>

	<Form.Field {form} name="dateOfBirth">
		<Form.Control let:attrs>
			<Form.Label>Date Of Birth <RedStar /></Form.Label>
			<Input {...attrs} type="date" bind:value={$formData.dateOfBirth} />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>

	<div>
		<Form.Field {form} name="password">
			<Form.Control let:attrs>
				<Form.Label>Password <RedStar /></Form.Label>
				<div class="relative">
					<Input
						{...attrs}
						type={isHide ? 'password' : 'text'}
						bind:value={$formData.password}
						placeholder="********"
					/>

					<button
						class="absolute right-0 bottom-[6px] mr-4 opacity-50"
						on:click|preventDefault={toogle}
					>
						{#if isHide}
							<EyeClosed />
						{:else}
							<EyeOpen />
						{/if}
					</button>
				</div>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>

		<div>
			<Form.Field {form} name="repeatPassword">
				<Form.Control let:attrs>
					<Form.Label>Repeat Password <RedStar /></Form.Label>

					<div class="relative">
						<Input
							{...attrs}
							type={isHide ? 'password' : 'text'}
							bind:value={$formData.repeatPassword}
							placeholder="********"
						/>

						<button
							class="absolute right-0 bottom-[6px] mr-4 opacity-50"
							on:click|preventDefault={toogle}
						>
							{#if isHide}
								<EyeClosed />
							{:else}
								<EyeOpen />
							{/if}
						</button>
					</div>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>

			<!-- icon button -->
		</div>
	</div>

	<div class="my-8"></div>

	<Form.Button class="w-full h-16 dark:bg-green-600 hover:dark:bg-green-500">SIGN-UP</Form.Button>
</form>
