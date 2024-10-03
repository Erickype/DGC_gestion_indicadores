<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import {
		addBooksOrChaptersProductionListsAuthorCareersSchema,
		type AddBooksOrChaptersProductionListsAuthorCareersSchema
	} from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { goto } from '$app/navigation';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import X from 'lucide-svelte/icons/circle-x';

	import AuthorsServer from '$lib/components/filters/authors/authorsServer.svelte';
	import FormSelect from '$lib/components/combobox/formSelect.svelte';

	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { GenerateComboMessagesFromCareers } from '$lib/api/controller/api/career';
	import type { Message } from '$lib/components/combobox/combobox';
	import type { CommonError } from '$lib/api/model/errors';
	import type { Career } from '$lib/api/model/api/career';

	export let data: SuperValidated<
		Infer<AddBooksOrChaptersProductionListsAuthorCareersSchema>,
		App.Superforms.Message
	>;
	export let comboMessages: Message[][];
	let careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function AcademicProductionListAuthorCareersCreated() {
		dispatch('message', { created: true });
	}

	const form = superForm(data, {
		validators: zodClient(addBooksOrChaptersProductionListsAuthorCareersSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message;
			if (!message) {
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				AcademicProductionListAuthorCareersCreated();
				return toast.success(`Autor y sus carreras creadas`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});
	const { form: formData, enhance } = form;

	let formDataAuthorID = writable($formData.authorID);
	formDataAuthorID.subscribe((value) => ($formData.authorID = value));
	let formDataCareerID = writable($formData.career);
	formDataCareerID.subscribe((value) => ($formData.career = value));

	let selectedCareers: Message[] = [];
	let selectedCareersPromise: Promise<Career[]> =
		fetchAcademicProductionListsAuthorPreviousCareersByAuthorID(0);

	$: {
		if ($formData.career !== 0) {
			addItem($formData.career);
			$formData.career = 0;
		}
	}
	formDataAuthorID.subscribe((value) => {
		selectedCareersPromise = fetchAcademicProductionListsAuthorPreviousCareersByAuthorID(value);
	});

	async function fetchAcademicProductionListsAuthorPreviousCareersByAuthorID(
		authorID: number
	): Promise<Career[]> {
		const url = `/api/indicatorsInformation/academicProductionListsAuthor/previousCareers/${authorID}`;
		const response = await fetch(url, { method: 'GET' });
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return response.json();
	}

	function setCareersComboData(careers: Career[]) {
		if (careers.length > 0) {
			selectedCareers = GenerateComboMessagesFromCareers(careers);
			$formData.careers = selectedCareers.map((career) => career.value);
		} else {
			selectedCareers = [];
			$formData.careers = [];
		}
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

<form action="?/postAcademicProductionListsAuthorCareers" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicProductionID">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicProductionID} name={attrs.name} />
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
				{#await selectedCareersPromise}
					cargando...
				{:then selectedCareersResponse}
					<div class="hidden">
						{setCareersComboData(selectedCareersResponse)}
					</div>
				{/await}
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
