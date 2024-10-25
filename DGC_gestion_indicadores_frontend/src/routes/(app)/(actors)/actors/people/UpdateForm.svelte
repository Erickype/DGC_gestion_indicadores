<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { roles, updatePersonSchema, type Role, type UpdatePersonSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	import * as Select from '$lib/components/ui/select';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { PersonWithRoles, UpdatePersonWithRolesRequest } from '$lib/api/model/api/person';
	import type { Message } from '$lib/components/combobox/combobox';
	import { toast } from 'svelte-sonner';

	export let data: SuperValidated<Infer<UpdatePersonSchema>, App.Superforms.Message>;
	export let updateEntity!: PersonWithRoles;
	export let comboMessages: Message[][];
	comboMessages = [];

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
				const person = message.data as UpdatePersonWithRolesRequest;
				PersonUpdated();
				return toast.success(`Persona y roles actualizados: ${person.person.identity}`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	$: selectedColors = $formData.roles.map((c) => ({ label: roles[c], value: c }));

	function fillForm() {
		$formData.ID = updateEntity.ID;
		$formData.identity = updateEntity.identity;
		$formData.name = updateEntity.name;
		$formData.lastname = updateEntity.lastname;
		$formData.email = updateEntity.email;
		updateEntity.roles!.forEach((role) => {
			$formData.roles = [...$formData.roles, role as Role];
		});
	}
</script>

<form action="?/updatePersonWithRoles" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="identity">
				<Form.Control let:attrs>
					<Form.Label>Identificación</Form.Label>
					<Input type="text" {...attrs} bind:value={$formData.identity} placeholder="1020304050" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="roles">
				<Form.Control let:attrs>
					<Form.Label>Roles</Form.Label>
					<Select.Root
						multiple
						selected={selectedColors}
						onSelectedChange={(s) => {
							if (s) {
								$formData.roles = s.map((c) => c.value);
							} else {
								$formData.roles = [];
							}
						}}
					>
						{#each $formData.roles as role}
							<input name={attrs.name} hidden value={role} />
						{/each}
						<Select.Trigger {...attrs}>
							<Select.Value placeholder="Seleccionar rol" />
						</Select.Trigger>
						<Select.Content>
							{#each Object.entries(roles) as [value, label]}
								<Select.Item {value} {label} />
							{/each}
						</Select.Content>
					</Select.Root>
					<Form.FieldErrors />
				</Form.Control>
			</Form.Field>
		</div>
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
