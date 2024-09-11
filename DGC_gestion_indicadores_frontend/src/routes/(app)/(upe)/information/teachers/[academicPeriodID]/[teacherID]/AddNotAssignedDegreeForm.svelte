<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addDegreeNotAssignedSchema, type AddDegreeNotAssignedSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<Infer<AddDegreeNotAssignedSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const notAssignedDegrees = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function DegreeNotAssignedAssigned() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addDegreeNotAssignedSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				DegreeNotAssignedAssigned();
				return toast.success(`Título asignado.`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;
	
	let formDataTeachersDegreeID = writable($formData.teachersDegreeID);
	formDataTeachersDegreeID.subscribe((value) => ($formData.teachersDegreeID = value));
</script>

<form action="?/postTeachersListsDegree" use:enhance>
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
		<Form.Field {form} name="teachersDegreeID" class="flex flex-col">
			<FormSelect
				formLabel="Título"
				comboData={notAssignedDegrees}
				bind:formDataID={formDataTeachersDegreeID}
			/>
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Agregar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
