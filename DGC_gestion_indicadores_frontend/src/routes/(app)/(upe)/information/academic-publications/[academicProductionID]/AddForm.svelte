<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import {
		addAcademicProductionListsAuthorSchema,
		type AddAcademicProductionListsAuthorSchema
	} from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import AuthorsServer from '$lib/components/filters/authors/authorsServer.svelte';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<
		Infer<AddAcademicProductionListsAuthorSchema>,
		App.Superforms.Message
	>;
	export let comboMessages: Message[][];

	const dispatch = createEventDispatcher();

	function AcademicProductionListAuthorCareersCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addAcademicProductionListsAuthorSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				/* 				const academicProduction = message.data as AcademicProductionList;
				 */ AcademicProductionListAuthorCareersCreated();
				/* 				return toast.success(`Publicación creada: ${academicProduction.publication_name}`);
				 */
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});
	const { form: formData, enhance } = form;

	let formDataAuthorID = writable($formData.authorID);
	formDataAuthorID.subscribe((value) => ($formData.authorID = value));
</script>

<form action="?/postAcademicProductionList" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicProductionID">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicProductionID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="authorID">
			<AuthorsServer {formDataAuthorID}></AuthorsServer>
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
