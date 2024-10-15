<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addGradeRateListSchema, type AddGradeRateListSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Input } from '$lib/components/ui/input/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import type { SocialProjectList } from '$lib/api/model/api/indicatorsInformation/socialProjectLists';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<Infer<AddGradeRateListSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function SocialProjectListCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addGradeRateListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as SocialProjectList;
				SocialProjectListCreated();
				return toast.success(`Proyecto de vinculación creado:${response.name}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataCareerID = writable($formData.career_id);
	formDataCareerID.subscribe((value) => ($formData.career_id = value));
</script>

<form action="?/postSocialProjectList" use:enhance>
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
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
