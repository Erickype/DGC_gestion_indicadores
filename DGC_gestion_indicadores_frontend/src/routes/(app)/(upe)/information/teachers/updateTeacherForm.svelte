<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateTeacherSchema, type UpdateTeacherSchema } from './scheme';
	import { updateTeacherAction } from '../../../../../stores';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import Button from '$lib/components/ui/button/button.svelte';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { browser } from '$app/environment';
	import { createEventDispatcher } from 'svelte';
	import { writable } from 'svelte/store';

	import type { TeachersListByAcademicPeriodJoined } from '$lib/api/model/api/indicatorsInformation/teachersLists';
	import TeachersServerFormSelect from '$lib/components/filters/teachers/teachersServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import { manageToastFromInvalidAddForm } from '$lib/utils';

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
			if ($message!.success) {
				updateTeacherAction.set({
					status: false,
					teacherID: -1
				});
				toast.success('Registro de docente actualizado.');
				dispatchTeacherUpdated();
			} else {
				toast.error('Fallo actualizando docente.');
			}
		}
	});

	const dispatch = createEventDispatcher();

	function dispatchTeacherUpdated() {
		dispatch('teacher-updated', {
			status: true
		});
	}

	const { form: formData, message, enhance } = form;

	$formData.teacher = updateEntity.teacher_id;
	$formData.career = updateEntity.career_id;
	$formData.dedication = updateEntity.dedication_id;
	$formData.scaledGrade = updateEntity.scaled_grade_id;
	$formData.contractType = updateEntity.contract_type_id;

	let formDataTeacherID = writable($formData.teacher);
	formDataTeacherID.subscribe((value) => ($formData.teacher = value));
	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));
	let formDataDedicationID = writable($formData.dedication);
	formDataDedicationID.subscribe((value) => ($formData.dedication = value));
	let formDataScaledGradeID = writable($formData.scaledGrade);
	formDataScaledGradeID.subscribe((value) => ($formData.scaledGrade = value));
	let formDataContractTypeID = writable($formData.contractType);
	formDataContractTypeID.subscribe((value) => ($formData.contractType = value));
</script>

<form action="?/updateTeacher" method="post" use:enhance>
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
		<Form.Field {form} name="scaledGrade" class="flex flex-col">
			<FormSelect
				formLabel="Grado escalafonado"
				comboData={scaledGradesComboData}
				bind:formDataID={formDataScaledGradeID}
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
	</div>
	<div class="flex items-center justify-center gap-2">
		<Form.Button>Actualizar</Form.Button>
		<Button
			variant="secondary"
			on:click={() => {
				updateTeacherAction.set({
					status: false,
					teacherID: -1
				});
			}}>Cancelar</Button
		>
	</div>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
