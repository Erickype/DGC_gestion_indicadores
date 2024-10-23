<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addPersonSchema, roles, type AddPersonSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { PostPersonWithRolesRequest } from '$lib/api/model/api/person';
	import * as Select from '$lib/components/ui/select';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	export let data: SuperValidated<Infer<AddPersonSchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function PersonWithRolesCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addPersonSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const person = message.data as PostPersonWithRolesRequest;
				PersonWithRolesCreated();
				return toast.success(`Persona y roles creados: ${person.person.identity}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	$formData.roles.push('teacher');

	$: selectedColors = $formData.roles.map((c) => ({ label: roles[c], value: c }));
</script>

<form action="?/addPersonWithRoles" use:enhance>
	<div class="flex flex-col gap-2">
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
