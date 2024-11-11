<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addPostgraduateCohortListSchema, type AddPostgraduateCohortListSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';

	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Message } from '$lib/components/combobox/combobox';
	import type {
		GetPostgraduateProgramMissingCohortYearsByProgramIDResponse,
		PostgraduateCohortList
	} from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import type { CommonError } from '$lib/api/model/errors';
	import { goto } from '$app/navigation';
	import { GenerateComboMessagesFromMissingYears } from '$lib/api/controller/api/indicatorsInformation/postgraduate';
	import { writable } from 'svelte/store';

	export let data: SuperValidated<Infer<AddPostgraduateCohortListSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	comboMessages = [];

	const dispatch = createEventDispatcher();

	function GradeRateListCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addPostgraduateCohortListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as PostgraduateCohortList;
				GradeRateListCreated();
				return toast.success(`Valores de tasas de cohorte creados: ${response.name}`);
			}
			$formData.name = f.data.name;
			$formData.year = f.data.year;
			$formData.graduated_students = f.data.graduated_students;
			$formData.total_students = f.data.total_students;
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataYear = writable($formData.year);
	formDataYear.subscribe((value) => ($formData.year = value));

	let missingYearsPromise: Promise<GetPostgraduateProgramMissingCohortYearsByProgramIDResponse> =
		fetchGetPostgraduateProgramMissingCohortYearsByProgramID();

	async function fetchGetPostgraduateProgramMissingCohortYearsByProgramID(): Promise<GetPostgraduateProgramMissingCohortYearsByProgramIDResponse> {
		const url = `/api/indicatorsInformation/postgraduate/program/cohort/missingYears/${$formData.postgraduate_program_id}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return response.json();
	}
</script>

<form action="?/postPostgraduateCohortList" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="postgraduate_program_id">
			<Form.Control let:attrs>
				<input hidden value={$formData.postgraduate_program_id} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="name">
				<Form.Control let:attrs>
					<Form.Label>Nombre</Form.Label>
					<Input {...attrs} bind:value={$formData.name} type="text" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="year" class="flex flex-col">
				{#await missingYearsPromise}
					<FormFieldSkeleton />
				{:then missingYears}
					<FormSelect
						formLabel="Año"
						formSelectWidth="w-[45%]"
						comboData={GenerateComboMessagesFromMissingYears(missingYears.years)}
						bind:formDataID={formDataYear}
					/>
				{/await}
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="graduated_students">
				<Form.Control let:attrs>
					<Form.Label>Estudiantes graduados</Form.Label>
					<Input {...attrs} bind:value={$formData.graduated_students} type="number" min="0" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="total_students">
				<Form.Control let:attrs>
					<Form.Label>Total estudiantes</Form.Label>
					<Input {...attrs} bind:value={$formData.total_students} type="number" min="0" />
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
