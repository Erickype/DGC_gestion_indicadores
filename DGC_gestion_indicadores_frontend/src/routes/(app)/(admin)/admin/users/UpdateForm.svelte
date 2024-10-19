<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateUserSchema, type UpdateUserSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { User } from '$lib/api/model/admin/user';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	import { writable } from 'svelte/store';

	export let data: SuperValidated<Infer<UpdateUserSchema>>;
	export let comboMessages: Message[][];
	export let updateEntity!: User;
	const rolesCombeData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function UserRoleUpdated() {
		dispatch('updated', {
			status: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateUserSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				UserRoleUpdated();
				return toast.success(`Rol de usuario actualizado`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	let formRoleID = writable($formData.role_id);
	formRoleID.subscribe((value) => ($formData.role_id = value));

	function fillForm() {
		$formData.ID = updateEntity.ID;
		$formData.username = updateEntity.username;
		$formData.email = updateEntity.email;
		$formData.role_id = updateEntity.role_id;
	}
</script>

<form action="?/updateUser" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="username">
				<Form.Control let:attrs>
					<Form.Label>Nombre</Form.Label>
					<Input
						disabled
						type="text"
						{...attrs}
						bind:value={$formData.username}
						placeholder="NombreUsuario"
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="email">
				<Form.Control let:attrs>
					<Form.Label>Abreviación</Form.Label>
					<Input
						disabled
						type="text"
						{...attrs}
						bind:value={$formData.email}
						placeholder="ejemplo@ejemplo.com"
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field {form} name="role_id">
			<FormSelect formLabel="Roles" comboData={rolesCombeData} bind:formDataID={formRoleID} />
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
