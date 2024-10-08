<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updatePersonSchema, type UpdatePersonSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Person } from '$lib/api/model/api/person';
	import { toast } from 'svelte-sonner';

	export let data: SuperValidated<Infer<UpdatePersonSchema>>;
	export let updateEntity: Person;

	const dispatch = createEventDispatcher();

	function PersonUpdated() {
		dispatch('updated', {
			status: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updatePersonSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const person = message.data as Person;
				PersonUpdated();
				return toast.success(`Persona actualizada: ${person.identity}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	$formData.ID = updateEntity!.ID;
	$formData.identity = updateEntity!.identity;
	$formData.name = updateEntity!.name;
	$formData.lastname = updateEntity!.lastname;
	$formData.email = updateEntity!.email;
</script>

<form action="?/updatePerson" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
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
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
