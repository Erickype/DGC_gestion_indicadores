<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addPersonSchema, type AddPersonSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import type { CommonError } from '$lib/api/model/errors';
	import { Input } from '$lib/components/ui/input';
	import * as Form from '$lib/components/ui/form';
	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	export let data: SuperValidated<Infer<AddPersonSchema>, App.Superforms.Message>;

	const dispatch = createEventDispatcher();

	function EvaluationPeriodCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addPersonSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (message.success) {
				EvaluationPeriodCreated();
				return toast.success(`Persona creada.`);
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

	const { form: formData, enhance } = form;
</script>

<form action="?/addPerson" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="identity">
			<Form.Control let:attrs>
				<Form.Label>Identificación</Form.Label>
				<Input type="text" {...attrs} bind:value={$formData.identity} placeholder="1020304050" />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<div class="grid grid-cols-2 justify-between gap-4">
			<Form.Field {form} name="name">
				<Form.Control let:attrs>
					<Form.Label>Nombre</Form.Label>
					<Input type="text" {...attrs} bind:value={$formData.name} placeholder="Erick" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
			<Form.Field {form} name="lastname">
				<Form.Control let:attrs>
					<Form.Label>Apellido</Form.Label>
					<Input type="text" {...attrs} bind:value={$formData.lastname} placeholder="Carrasco" />
				</Form.Control>
				<Form.FieldErrors />
			</Form.Field>
		</div>
		<Form.Field {form} name="email">
			<Form.Control let:attrs>
				<Form.Label>Correo</Form.Label>
				<Input
					type="email"
					{...attrs}
					bind:value={$formData.email}
					placeholder="ecarrasco@mail.com"
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
