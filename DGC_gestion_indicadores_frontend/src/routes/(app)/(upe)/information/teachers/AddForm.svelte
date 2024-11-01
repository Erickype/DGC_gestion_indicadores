<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addTeacherSchema, type AddTeacherSchema } from './scheme';
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

	export let data: SuperValidated<Infer<AddTeacherSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;
	const dedicationsComboData = comboMessages.at(1)!;
	const scaledGradesComboData = comboMessages.at(2)!;
	const contractTypesComboData = comboMessages.at(3)!;

	const dispatch = createEventDispatcher();

	function TeacherCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addTeacherSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				TeacherCreated();
				return toast.success(`Docente creado`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataTeacherID = writable($formData.teacher);
	let formDataCareerID = writable($formData.career);
	let formDataDedicationID = writable($formData.dedication);
	let formDataScaledGradeID = writable($formData.scaledGrade);
	let formDataContractTypeID = writable($formData.contractType);

	fillForm();

	function fillForm() {
		formDataTeacherID.subscribe((value) => ($formData.teacher = value));
		formDataCareerID.subscribe((value) => ($formData.career = value));
		formDataDedicationID.subscribe((value) => ($formData.dedication = value));
		formDataScaledGradeID.subscribe((value) => ($formData.scaledGrade = value));
		formDataContractTypeID.subscribe((value) => ($formData.contractType = value));
	}
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="teacher" class="flex flex-col">
			<TeachersServerFormSelect bind:formDataTeacherID />
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="career" class="flex flex-col">
			<FormSelect
				formLabel="Carreras"
				comboData={careersComboData}
				bind:formDataID={formDataCareerID}
				scrollAreaHeight="h-72"
			/>
		</Form.Field>
		<Form.Field {form} name="dedication" class="flex flex-col">
			<FormSelect
				formLabel="Dedicación"
				comboData={dedicationsComboData}
				bind:formDataID={formDataDedicationID}
			/>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="contractType" class="flex flex-col">
			<FormSelect
				formLabel="Tipo contrato"
				comboData={contractTypesComboData}
				bind:formDataID={formDataContractTypeID}
			/>
			<Form.FieldErrors />
		</Form.Field>
		{#if $formData.contractType && contractTypesComboData.find((contractType) => contractType.value === $formData.contractType)?.label === 'Nombramiento'}
			<Form.Field {form} name="scaledGrade" class="flex flex-col">
				<FormSelect
					formLabel="Grado escalafonario"
					comboData={scaledGradesComboData}
					bind:formDataID={formDataScaledGradeID}
				/>
				<Form.FieldErrors />
			</Form.Field>
		{/if}
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
