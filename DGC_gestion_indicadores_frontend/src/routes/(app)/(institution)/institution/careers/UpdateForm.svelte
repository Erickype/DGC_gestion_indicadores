<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateCarrerSchema, type UpdateCarrerSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { createEventDispatcher, onMount } from 'svelte';

	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Faculty } from '$lib/api/model/api/faculty';
	import type { Career } from '$lib/api/model/api/career';
	import {
		fetchFaculties,
		GenerateComboMessagesFromFaculties
	} from '$lib/components/filters/faculties';

	export let data: SuperValidated<Infer<UpdateCarrerSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	comboMessages = [];
	export let updateEntity: Career;

	const facultiesPromise: Promise<Faculty[]> = fetchFaculties();

	const dispatch = createEventDispatcher();

	function CareerUpdated() {
		dispatch('updated', {
			status: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateCarrerSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const career = message.data as Career;
				CareerUpdated();
				return toast.success(`Carrera actualizada: ${career.abbreviation}`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});
	const { form: formData, enhance } = form;

	function fillForm() {
		$formData.ID = updateEntity!.ID;
		$formData.facultyID = updateEntity!.faculty_id;
		$formData.name = updateEntity!.name;
		$formData.abbreviation = updateEntity!.abbreviation;
		$formData.description = updateEntity!.description!;
		$formDataFacultyID = $formData.facultyID;
	}

	let formDataFacultyID = writable($formData.facultyID);
	formDataFacultyID.subscribe((value) => ($formData.facultyID = value));

	onMount(async () => {
		fillForm();
	});
</script>

<form action="?/updateCareer" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
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
				<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Facultad de ..." />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="abbreviation">
			<Form.Control let:attrs>
				<Form.Label>Abreviación</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.abbreviation} placeholder="FISEI" />
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
