<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAcademicProductionSchema, type AddAcademicProductionSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import TeachersServerFormSelect from '$lib/components/filters/teachers/teachersServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { Teacher } from '$lib/api/model/api/teacher';

	export let data: SuperValidated<Infer<AddAcademicProductionSchema>, App.Superforms.Message>;
	/* 	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;
	const dedicationsComboData = comboMessages.at(1)!;
	const scaledGradesComboData = comboMessages.at(2)!;
	const contractTypesComboData = comboMessages.at(3)!; */

	const dispatch = createEventDispatcher();

	function TeacherCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addAcademicProductionSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const teacher = message.data as Teacher;
				TeacherCreated();
				return toast.success(`Docente creado: ${teacher.person_id}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

/*	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));*/
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<!-- 		<Form.Field {form} name="career" class="flex flex-col">
			<FormSelect
				formLabel="Carreras"
				comboData={careersComboData}
				bind:formDataID={formDataCareerID}
				scrollAreaHeight="h-72"
			/> -->
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
