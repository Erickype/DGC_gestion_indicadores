<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { updateSocialProjectListSchema, type UpdateSocialProjectListSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import type { SocialProjectList } from '$lib/api/model/api/indicatorsInformation/socialProjectLists';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';
	import type { Message } from '$lib/components/combobox/combobox';

	export let updateEntity!: SocialProjectList;
	export let data: SuperValidated<Infer<UpdateSocialProjectListSchema>, App.Superforms.Message>;
	export let comboMessages: Message[][];
	const careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function SocialProjectListCreated() {
		dispatch('updated', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(updateSocialProjectListSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const response = message.data as SocialProjectList;
				SocialProjectListCreated();
				return toast.success(`Proyecto de vinculación actualizado:${response.name}`);
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
		$formData.ID = updateEntity.ID!;
		$formData.career_id = updateEntity.career_id;
		$formData.name = updateEntity.name;
	}
</script>

<form action="?/putSocialProjectList" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="ID">
			<Form.Control let:attrs>
				<input hidden value={$formData.ID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
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
		<Form.Field {form} name="name">
			<Form.Control let:attrs>
				<Form.Label>Título</Form.Label>
				<Textarea
					{...attrs}
					placeholder="Saberes emergentes de las artes urbanas..."
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
