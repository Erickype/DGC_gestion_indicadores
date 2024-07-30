<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { browser } from '$app/environment';

	import type { CommonError } from '$lib/api/model/errors';

	import { createEventDispatcher } from 'svelte';
	
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import { updateFacultySchema, type UpdateFacultySchema } from './schema';
	import type { Faculty } from '$lib/api/model/api/faculty';

	export let data: SuperValidated<Infer<UpdateFacultySchema>>;
	export let updateEntity: Faculty;

	const dispatch = createEventDispatcher();

	function FacultyUpdated() {
		dispatch('updated', {
			status: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateFacultySchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (message.success) {
				FacultyUpdated();
				return toast.success(`Facultad actualizada.`);
			}
			if (message.error as CommonError) {
				const commonError = message!.error as CommonError;
				fillForm();
				return toast.error(`${commonError.message}: ${commonError.detail}`);
			} else {
				const error = message.error as string;
				fillForm();
				return toast.error(error);
			}
		}
	});

	function fillForm() {
		$formData.ID = updateEntity!.ID;
		$formData.name = updateEntity!.name;
		$formData.abbreviation = updateEntity!.abbreviation;
		$formData.description = updateEntity!.description!;
	}

	const { form: formData, enhance } = form;

	$: {
		$formData.ID = updateEntity!.ID;
		$formData.name = updateEntity!.name;
		$formData.abbreviation = updateEntity!.abbreviation;
		$formData.description = updateEntity!.description!;
	}
</script>

<form action="?/updateFaculty" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
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
