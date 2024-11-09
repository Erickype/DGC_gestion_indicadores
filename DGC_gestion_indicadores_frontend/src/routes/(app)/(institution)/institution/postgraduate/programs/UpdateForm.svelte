<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updatePostgraduateProgramSchema, type UpdatePostgraduateProgramSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import { Switch } from '$lib/components/ui/switch/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { PostgraduateProgram } from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import type { Message } from '$lib/components/combobox/combobox';

	export let updateEntity!: PostgraduateProgram;
	export let data: SuperValidated<Infer<UpdatePostgraduateProgramSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	comboMessages = [];

	const dispatch = createEventDispatcher();

	function PostgraduateProgramUpdated() {
		dispatch('updated', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updatePostgraduateProgramSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as PostgraduateProgram;
				PostgraduateProgramUpdated();
				return toast.success(`Programa posgrado actualizado: ${response.name}`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	function fillForm() {
		$formData.id = updateEntity.ID!;
		$formData.name = updateEntity.name;
		$formData.start_year = updateEntity.start_year;
		$formData.end_year = updateEntity.end_year;
		$formData.is_active = updateEntity.is_active;
	}
</script>

<form action="?/updatePostgraduateProgram" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="id">
			<Form.Control let:attrs>
				<input hidden value={$formData.id} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Maestría en ..."
					class="resize-none"
					bind:value={$formData.name}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="start_year">
				<Form.Control let:attrs>
					<Form.Label>Año inicio</Form.Label>
					<Input
						type="number"
						min="2020"
						{...attrs}
						placeholder="2022"
						bind:value={$formData.start_year}
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="end_year">
				<Form.Control let:attrs>
					<Form.Label>Año fin</Form.Label>
					<Input
						type="number"
						min="2020"
						{...attrs}
						placeholder="2025"
						bind:value={$formData.end_year}
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field
			{form}
			name="is_active"
			class="flex flex-row items-center justify-between rounded-lg border p-4"
		>
			<Form.Control let:attrs>
				<div class="space-y-0.5">
					<Form.Label>Estado</Form.Label>
					<Form.Description>{$formData.is_active ? 'Activo' : 'Inactivo'}</Form.Description>
				</div>
				<Switch {...attrs} aria-readonly includeInput bind:checked={$formData.is_active} />
			</Form.Control>
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
