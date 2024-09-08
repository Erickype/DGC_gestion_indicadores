<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import { addTeachersListsDegreeSchema, type AddTeachersListsDegreeSchema } from './schema';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';

	import type { Faculty } from '$lib/api/model/api/faculty';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';

	export let data: SuperValidated<Infer<AddTeachersListsDegreeSchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function FacultyCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addTeachersListsDegreeSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const faculty = message.data as Faculty;
				FacultyCreated();
				return toast.success(`Título creado: ${faculty.abbreviation}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;
</script>

<form action="?/addTeachersListsDegree" use:enhance>
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
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Diploma de ..." />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- 	{#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
