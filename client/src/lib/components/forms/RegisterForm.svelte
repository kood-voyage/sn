<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { signUpSchema, type SignUpSchema } from '../../../routes/(auth)/schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { v4 as uuidv4 } from 'uuid';

	import RedStar from '../../../routes/(auth)/signup/RedStar.svelte';
	import EyeOpen from 'svelte-radix/EyeOpen.svelte';
	import EyeClosed from 'svelte-radix/EyeClosed.svelte';

	import { RegisterUser } from '$lib/client/api/user-requests';

	import { Confetti } from 'svelte-confetti';

	import { goto } from '$app/navigation';
	import toast from 'svelte-french-toast';

	let confetti = false;

	function toogle() {
		isHide = !isHide;
	}

	export type UserModel = {
		id: string;
		username: string;
		email: string;
		password: string;
		timestamp?: string;
		date_of_birth: string;
		first_name: string;
		last_name: string;
		description: string;
		avatar: string;
		cover: string;
		privacy: string;
	};

	export let data: SuperValidated<Infer<SignUpSchema>>;

	let isHide: boolean = true;

	const form = superForm(data, {
		validators: zodClient(signUpSchema),
		onSubmit: async ({ formData, cancel }) => {
			const { username, email, dateOfBirth, password, firstName, lastName } = $formData;

			const user: UserModel = {
				id: uuidv4(),
				username,
				email,
				password,
				date_of_birth: dateOfBirth,
				first_name: firstName,
				last_name: lastName,
				description: `👋, I'm ${username}.`,
				avatar: `https://api.dicebear.com/7.x/bottts-neutral/svg?seed=${username}`,
				cover:
					'https://t3.ftcdn.net/jpg/03/10/17/76/360_F_310177612_ZF5ucdsR1SEm76F9ydhfLzlotishG1Ug.jpg',
				privacy: 'public'
			};

			const resp = await RegisterUser(user);

			if (resp?.ok) {
				confetti = true;
				setTimeout(() => {
					goto('/signin');
				}, 3000);
			}

			cancel();
		},
		onError: (event) => {
			console.log(event);
		}
	});

	const { form: formData, enhance } = form;
</script>

{#if confetti}
	<div class="fixed top-[-50px] left-0 h-screen w-screen flex justify-center overflow-hidden">
		<Confetti
			x={[-5, 8]}
			y={[0, 0.1]}
			delay={[0, 1400]}
			duration="2000"
			amount="300"
			fallDistance="100vh"
		/>

		<div class="h-full flex flex-col items-center p-20">
			<img src={'success.png'} alt="success" class="" />
		</div>
	</div>
{:else}
	<form method="POST" use:enhance>
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
{/if}
