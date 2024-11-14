<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addScienceMagazineSchema, type AddScienceMagazineSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import type { ScienceMagazine } from '$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<Infer<AddScienceMagazineSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const academicDatabasesComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function ScienceMagazineCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addScienceMagazineSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as ScienceMagazine;
				ScienceMagazineCreated();
				return toast.success(`Revista científica creada: ${response.name}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataAcademicDatabaseID = writable($formData.academic_database_id);
	formDataAcademicDatabaseID.subscribe((value) => ($formData.academic_database_id = value));
</script>

<form action="?/postScienceMagazine" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academic_database_id" class="flex flex-col">
			<FormSelect
				formLabel="Base de datos"
				comboData={academicDatabasesComboData}
				bind:formDataID={formDataAcademicDatabaseID}
			/>
		</Form.Field>
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Revista en ..."
					class="resize-none"
					bind:value={$formData.name}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="abbreviation">
			<Form.Control let:attrs>
				<Form.Label>Abbreviación</Form.Label>
				<Input type="text" {...attrs} placeholder="Rev..." bind:value={$formData.abbreviation} />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Descripción</Form.Label>
				<Input
					{...attrs}
					type="text"
					placeholder="Revista dedicada a..."
					bind:value={$formData.description}
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button class="my-2 w-full">Guardar</Form.Button>
		<!-- {#if browser}
			<SuperDebug data={$formData} />
		{/if} -->
	</div>
</form>
