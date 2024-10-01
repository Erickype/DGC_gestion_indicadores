<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addAcademicProductionListsAuthorSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { toast } from 'svelte-sonner';

	import X from 'lucide-svelte/icons/circle-x';

	import FormSelect from '$lib/components/combobox/formSelect.svelte';

	import type { AcademicProductionListsAuthorsCareersJoined } from '$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor';
	import { manageToastFromErrorMessageOnAddForm, manageToastFromInvalidAddForm } from '$lib/utils';
	import { GenerateComboMessagesFromCareers } from '$lib/api/controller/api/career';
	import type { Message } from '$lib/components/combobox/combobox';

	export let data: SuperValidated<
		Infer<typeof addAcademicProductionListsAuthorSchema>,
		App.Superforms.Message
	>;
	export let updateEntity: AcademicProductionListsAuthorsCareersJoined;
	export let comboMessages: Message[][];

	let careersComboData = comboMessages.at(0)!;

	const dispatch = createEventDispatcher();

	function AcademicProductionListAuthorCareersCreated() {
		dispatch('updated', { created: true });
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
				AcademicProductionListAuthorCareersCreated();
				return toast.success(`Autor y sus carreras actualizadas`);
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

<form action="?/updateAcademicProductionListsAuthorCareers" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field {form} name="academicProductionID">
			<Form.Control let:attrs>
				<input hidden value={$formData.academicProductionID} name={attrs.name} />
			</Form.Control>
		</Form.Field>
		<Form.Field {form} name="authorID" class="flex flex-col">
			<Form.Control let:attrs>
				<label
					class="data-[fs-error]:text-destructive text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
					for="teacher"
				>
					Autor
				</label>
				<p
					class="ring-offset-background focus-visible:ring-ring border-input bg-background hover:bg-accent hover:text-accent-foreground inline-flex h-10 w-full items-center justify-between whitespace-nowrap rounded-md border px-4 py-2 text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"
				>
					{updateEntity.author}
				</p>
				<input hidden value={$formData.authorID} name={attrs.name} />
			</Form.Control>
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
	{#if browser}
		<SuperDebug data={$formData} />
	{/if}
</form>
