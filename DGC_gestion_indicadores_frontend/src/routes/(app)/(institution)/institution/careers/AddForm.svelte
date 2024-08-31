<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addCarrerSchema, type AddCarrerSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';
	import type { Faculty } from '$lib/api/model/api/faculty';
	import type { Career } from '$lib/api/model/api/career';
	import {
		fetchFaculties,
		GenerateComboMessagesFromFaculties
	} from '$lib/components/filters/faculties';
	import { writable } from 'svelte/store';

	export let data: SuperValidated<Infer<AddCarrerSchema>, App.Superforms.Message>;

	const facultiesPromise: Promise<Faculty[]> = fetchFaculties();

	const dispatch = createEventDispatcher();

	function CareerCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addCarrerSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const career = message.data as Career;
				CareerCreated();
				return toast.success(`Carrera creada: ${career.abbreviation}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});
	const { form: formData, enhance } = form;

	let formDataFacultyID = writable($formData.facultyID);
	formDataFacultyID.subscribe((value) => ($formData.facultyID = value));
</script>

<form action="?/addCareer" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field class="flex flex-col" {form} name="facultyID">
			{#await facultiesPromise}
				<FormFieldSkeleton />
			{:then faculties}
				<FormSelect
					formLabel="Facultad"
					comboData={GenerateComboMessagesFromFaculties(faculties)}
					bind:formDataID={formDataFacultyID}
				/>
			{/await}
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Carrera de ..." />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="abbreviation">
			<Form.Control let:attrs>
				<Form.Label>Abreviación</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.abbreviation}
					placeholder="Ing. en ..."
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Descripción</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.description}
					placeholder="Dedicada a ..."
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
