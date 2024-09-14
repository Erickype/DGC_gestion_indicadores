<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAuthorFromPersonSchema, type AddAuthorFromPersonSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Author } from '$lib/api/model/api/academicProduction/authors/author';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';

	export let data: SuperValidated<Infer<AddAuthorFromPersonSchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function AuthorFromPersonCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addAuthorFromPersonSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const _ = message.data as Author;
				AuthorFromPersonCreated();
				return toast.success(`Autor creado`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;
</script>

<form action="?/postAuthorFromPerson" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="identity">
			<Form.Control let:attrs>
				<Form.Label>Identificación</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.identity} placeholder="1020304050" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="name">
				<Form.Control let:attrs>
					<Form.Label>Nombre</Form.Label>
					<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Erick" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="lastname">
				<Form.Control let:attrs>
					<Form.Label>Apellido</Form.Label>
					<Input type="text" {...attrs} bind:value={$formData.lastname} placeholder="Carrasco" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field {form} name="email">
			<Form.Control let:attrs>
				<Form.Label>Correo</Form.Label>
				<Input
					type="email"
					{...attrs}
					bind:value={$formData.email}
					placeholder="ecarrasco@mail.com"
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- 	{#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
