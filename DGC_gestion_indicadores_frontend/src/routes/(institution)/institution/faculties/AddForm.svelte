<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import { addFacultySchema, type AddFacultySchema } from './schema';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	
	import type { CommonError } from '$lib/api/model/errors';
	import type { Faculty } from '$lib/api/model/api/faculty';

	export let data: SuperValidated<Infer<AddFacultySchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function EvaluationPeriodCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addFacultySchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (message.success) {
				const faculty = message.data as Faculty
				EvaluationPeriodCreated();
				return toast.success(`Facultad creada: ${faculty.abbreviation}`);
			}
			if (message.error as CommonError) {
				const commonError = message!.error as CommonError;
				return toast.error(`${commonError.message}: ${commonError.detail}`);
			} else {
				const error = message.error as string;
				return toast.error(error);
			}
		}
	});

	const { form: formData, message, enhance } = form;
</script>

<form action="?/addFaculty" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Nombre</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Facultad de ..." />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="abbreviation">
			<Form.Control let:attrs>
				<Form.Label>Abreviación</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.abbreviation} placeholder="FISEI" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="description">
			<Form.Control let:attrs>
				<Form.Label>Descripción</Form.Label>
				<Input
					type="text"
					{...attrs}
					bind:value={$formData.description}
					placeholder="Dedicada a ..."
				/>
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
<!-- 	{#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
