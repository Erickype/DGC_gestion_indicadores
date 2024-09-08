<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import {
		addDegreeAndTeachersListsDegreeSchema,
		type AddDegreeAndTeachersListsDegreeSchema
	} from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<
		Infer<AddDegreeAndTeachersListsDegreeSchema>,
		App.Superforms.Message
	>;
	export let comboMessages: Message[][];
	const degreeLevelsComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function FacultyCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addDegreeAndTeachersListsDegreeSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				FacultyCreated();
				return toast.success(`Título creado.`);
			}
			$formData.degreeLevelID = $formDataDegreeLevelID
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataDegreeLevelID = writable($formData.degreeLevelID);
	formDataDegreeLevelID.subscribe((value) => ($formData.degreeLevelID = value));
</script>

<form action="?/addDegreeAndTeachersListsDegree" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicPeriodID">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicPeriodID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="teacherID">
			<Form.Control let:attrs>
				<input hidden value={$formData.teacherID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="degreeLevelID" class="flex flex-col">
			<FormSelect
				formLabel="Nivel título"
				comboData={degreeLevelsComboData}
				bind:formDataID={formDataDegreeLevelID}
			/>
		</Form.Field>
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Diploma/Master/Doctorado en ..."
					class="resize-none"
					bind:value={$formData.name}
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
