<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateTeacherSchema, type UpdateTeacherSchema } from './scheme';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import type { TeachersListByAcademicPeriodJoined } from '$lib/api/model/api/indicatorsInformation/teachersLists';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let updateEntity!: TeachersListByAcademicPeriodJoined;
	export let data: SuperValidated<Infer<UpdateTeacherSchema>>;
	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;
	const dedicationsComboData = comboMessages.at(1)!;
	const scaledGradesComboData = comboMessages.at(2)!;
	const contractTypesComboData = comboMessages.at(3)!;

	const form = superForm(data, {
		validators: zodClient(updateTeacherSchema),
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				TeacherUpdated();
				return toast.success(`Registro de profesor actualizado.`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const dispatch = createEventDispatcher();

	function TeacherUpdated() {
		dispatch('updated', {
			status: true
		});
	}

	const { form: formData, enhance } = form;

	fillForm();

	let formDataTeacherID = writable($formData.teacher);
	formDataTeacherID.subscribe((value) => ($formData.teacher = value));
	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));
	let formDataDedicationID = writable($formData.dedication);
	formDataDedicationID.subscribe((value) => ($formData.dedication = value));
	let formDataContractTypeID = writable($formData.contractType);
	formDataContractTypeID.subscribe((value) => {
		$formData.contractType = value;
		if (value !== 1) {
			$formData.scaledGrade = null;
		} else {
			$formData.scaledGrade = updateEntity.scaled_grade_id;
		}
	});

	let formDataScaledGradeID = writable($formData.scaledGrade);
	formDataScaledGradeID.subscribe((value) => ($formData.scaledGrade = value));
	formData.subscribe((value) => {
		$formDataScaledGradeID = value.scaledGrade;
	});

	function fillForm() {
		$formData.teacher = updateEntity.teacher_id;
		$formData.career = updateEntity.career_id;
		$formData.dedication = updateEntity.dedication_id;
		if (updateEntity.scaled_grade_id === 0) {
			$formData.scaledGrade = null;
		} else {
			$formData.scaledGrade = updateEntity.scaled_grade_id;
		}
		$formData.contractType = updateEntity.contract_type_id;
	}
</script>

<form action="?/updateTeacher" method="post" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriod">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriod} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="teacher" class="flex flex-col">
			<Form.Control let:attrs>
				<label
					class="data-[fs-error]:text-destructive text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					for="teacher"
				>
					Profesor
				</label>
				<p
					class="ring-offset-background focus-visible:ring-ring border-input bg-background hover:bg-accent hover:text-accent-foreground inline-flex h-10 w-full items-center justify-between whitespace-nowrap rounded-md border px-4 py-2 text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"
				>
					{updateEntity.teacher_identity + ' ' + updateEntity.teacher}
				</p>
				<input hidden value={$formData.teacher} name={attrs.name} />
			</Form.Control>
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
