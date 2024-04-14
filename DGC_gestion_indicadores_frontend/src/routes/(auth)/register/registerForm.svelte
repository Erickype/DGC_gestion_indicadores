<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import CircleAlert from 'lucide-svelte/icons/circle-alert';
	import { registerSchema, type RegisterSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { page } from '$app/stores';

	export let data: SuperValidated<Infer<RegisterSchema>>;

	const form = superForm(data, {
		validators: zodClient(registerSchema)
	});

	const { form: formData, message, enhance } = form;
</script>

<form action="?/register" method="POST" use:enhance>
	<Form.Field {form} name="username" class="my-2">
		<Form.Control let:attrs>
			<Form.Label>Usuario</Form.Label>
			<Input type="text" {...attrs} bind:value={$formData.username} />
		</Form.Control>
		<Form.Description>Nombre de nuevo usuario</Form.Description>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Field {form} name="email" class="my-2">
		<Form.Control let:attrs>
			<Form.Label>Email</Form.Label>
			<Input type="email" placeholder="ma@example.com" {...attrs} bind:value={$formData.email} />
		</Form.Control>
		<Form.Description>Email de nuevo usuario</Form.Description>
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
		Ya tiene una cuenta?
		<a href="/login" class="underline"> Ingresr </a>
	</div>
</form>
