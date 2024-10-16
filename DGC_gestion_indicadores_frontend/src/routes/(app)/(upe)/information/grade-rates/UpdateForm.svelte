<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateGradeRateListSchema, type UpdateGradeRateListSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { GradeRateListJoined } from '$lib/api/model/api/indicatorsInformation/gradeRateLists';

	export let data: SuperValidated<Infer<UpdateGradeRateListSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	export let updateEntity: GradeRateListJoined;
	const careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function GradeRateListCreated() {
		dispatch('updated', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateGradeRateListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				GradeRateListCreated();
				return toast.success(`Valores de tasas de grado creados`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	fillForm();

	let formDataCareerID = writable($formData.career_id);
	formDataCareerID.subscribe((value) => ($formData.career_id = value));

	function fillForm() {
		$formData.career_id = updateEntity.career_id;
		$formData.count_graduated_students = updateEntity.count_graduated_students;
		$formData.count_court_students = updateEntity.count_court_students;
		$formData.count_admitted_matriculated_students =
			updateEntity.count_admitted_matriculated_students;
		$formData.count_admitted_students = updateEntity.count_admitted_students;
	}
</script>

<form action="?/updateGradeRateList" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academic_period_id">
			<Form.Control let:attrs>
				<input hidden value={$formData.academic_period_id} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="career_id">
			<FormSelect
				formLabel="Carreras"
				comboData={careersComboData}
				bind:formDataID={formDataCareerID}
				scrollAreaHeight="h-72"
			/>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="count_graduated_students">
				<Form.Control let:attrs>
					<Form.Label>Estudiantes graduados</Form.Label>
					<Input {...attrs} bind:value={$formData.count_graduated_students} type="number" min="0" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="count_court_students">
				<Form.Control let:attrs>
					<Form.Label>Estudiantes en cohorte</Form.Label>
					<Input {...attrs} bind:value={$formData.count_court_students} type="number" min="0" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="count_admitted_matriculated_students">
				<Form.Control let:attrs>
					<Form.Label>Estudiantes matriculados</Form.Label>
					<Input
						{...attrs}
						bind:value={$formData.count_admitted_matriculated_students}
						type="number"
						min="0"
					/>
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="count_admitted_students">
				<Form.Control let:attrs>
					<Form.Label>Estudiantes admitidos</Form.Label>
					<Input {...attrs} bind:value={$formData.count_admitted_students} type="number" min="0" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
