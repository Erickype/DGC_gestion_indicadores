<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import {
		updateBooksOrChaptersProductionListsAuthorCareersSchema,
		type UpdateBooksOrChaptersProductionListsAuthorCareersSchema
	} from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import X from 'lucide-svelte/icons/circle-x';

	import AuthorsServer from '$lib/components/filters/authors/authorsServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';

	import type { BooksOrChaptersProductionListsAuthorsCareersJoined } from '$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionListsAuthor';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { GenerateComboMessagesFromCareers } from '$lib/api/controller/api/career';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<
		Infer<UpdateBooksOrChaptersProductionListsAuthorCareersSchema>,
		App.Superforms.Message
	>;
	export let updateEntity: BooksOrChaptersProductionListsAuthorsCareersJoined;
	export let comboMessages: Message[][];
	let careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function UpdateProductionListAuthorCareersCreated() {
		dispatch('updated', { created: true });
	}

	const form = superForm(data, {
		validators: zodClient(updateBooksOrChaptersProductionListsAuthorCareersSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				UpdateProductionListAuthorCareersCreated();
				return toast.success(`Autor y sus carreras actulizadas`);
			}
			fillForm();
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});
	const { form: formData, enhance } = form;

	let selectedCareers: Message[] = [];

	fillForm();

	let formDataAuthorID = writable($formData.authorID);
	formDataAuthorID.subscribe((value) => ($formData.authorID = value));
	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));

	$: {
		if ($formData.career !== 0) {
			addItem($formData.career);
			$formData.career = 0;
		}
	}

	function fillForm() {
		selectedCareers = GenerateComboMessagesFromCareers(updateEntity.careers);
		$formData.authorID = updateEntity.author_id;
		$formData.careers = updateEntity.careers.map((career) => career.ID);
	}

	function addItem(id: number) {
		if (!selectedCareers.some((career) => career.value === id)) {
			$formData.careers = [...$formData.careers, id];
			selectedCareers = [
				...selectedCareers,
				careersComboData.find((career) => career.value === id)!
			];
		}
	}

	function removeItem(id: number) {
		$formData.careers = $formData.careers.filter((i) => i !== id);
		selectedCareers = selectedCareers.filter((i) => i.value !== id);
	}
</script>

<form action="?/updateBooksOrChaptersProductionListsAuthorCareers" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="booksOrChaptersProductionListID">
			<Form.Control let:attrs>
				<input hidden value={$formData.booksOrChaptersProductionListID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="authorID">
			<AuthorsServer {formDataAuthorID}></AuthorsServer>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="career" class="flex flex-col">
			<FormSelect
				formLabel="Carreras"
				bind:formDataID={formDataCareerID}
				comboData={careersComboData}
				scrollAreaHeight="h-72"
			/>
		</Form.Field>

		<Form.Fieldset {form} name="careers" class="space-y-0">
			<ScrollArea class="h-32 rounded-md border">
				{#each selectedCareers as item}
					<div class="flex flex-row items-start space-x-3 p-2">
						<Form.Control let:attrs>
							<Button class="m-0 h-min p-0" variant="ghost" on:click={() => removeItem(item.value)}>
								<X class="stroke-primary h-4 w-4"></X>
							</Button>
							<Form.Label class="font-normal">{item.label}</Form.Label>
							<input hidden name={attrs.name} value={item.value} />
						</Form.Control>
					</div>
				{/each}
				{#if selectedCareers.length === 0}
					<p class="flex h-32 items-center justify-center">
						Aquí se mostrarán las carreras seleccionadas
					</p>
				{/if}
			</ScrollArea>
			<Form.FieldErrors />
		</Form.Fieldset>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
