<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import CircleAlert from 'lucide-svelte/icons/circle-alert';
	import { loginSchema, type LoginSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { page } from '$app/stores';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { toast } from 'svelte-sonner';
	import type { LoginResponse } from '$lib/api/model/auth/login';

	export let data: SuperValidated<Infer<LoginSchema>>;

	const form = superForm(data, {
		validators: zodClient(loginSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			console.log(message);
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const loginResponse = message.data as LoginResponse;
				return toast.success(`Ingreso exitoso, bienvenido ${loginResponse.username}!`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, message, enhance } = form;
</script>

<form action="?/login" method="POST" use:enhance>
	<Form.Field {form} name="username" class="my-2">
		<Form.Control let:attrs>
			<Form.Label>Usuario</Form.Label>
			<Input type="text" {...attrs} bind:value={$formData.username} />
		</Form.Control>
		<Form.Description>Nombre de usuario registrado</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="password" class="my-2">
		<Form.Control let:attrs>
			<Form.Label>Contraseña</Form.Label>
			<Input type="password" {...attrs} bind:value={$formData.password} />
		</Form.Control>
		<Form.Description>Contraseña del usuario</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	{#if $message && $page.status >= 400}
		<Alert.Root class="my-2" variant="destructive">
			<CircleAlert class="h-4 w-4" />
			<Alert.Title>Error</Alert.Title>
			<Alert.Description>Credenciales Incorrectas. Vuelva a intentarlo.</Alert.Description>
		</Alert.Root>
	{/if}
	<Form.Button class="my-2 w-full">Enviar</Form.Button>
	<div class="mt-4 text-center text-sm">
		No tiene una cuenta?
		<a href="/register" class="underline"> Registrarse </a>
	</div>
</form>
