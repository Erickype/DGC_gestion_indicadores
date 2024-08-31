<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addTeacherSchema, type AddTeacherSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';
	import { toast } from 'svelte-sonner';

	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';
	import * as Form from '$lib/components/ui/form';

	import type { FilterPeopleRequest, FilterPeopleResponse } from '$lib/api/model/api/person';
	import { GenerateComboMessagesFromPeople } from '$lib/api/controller/api/person';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { Teacher } from '$lib/api/model/api/teacher';
	import {
		fetchFilterPeople,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterPeopleRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/people';
	import {
		manageToastFromErrorMessageOnAddForm,
		manageToastFromInvalidAddForm
	} from '$lib/utils.js';

	export let data: SuperValidated<Infer<AddTeacherSchema>, App.Superforms.Message>;

	let openPeople = false;

	let filterPeopleRequest: FilterPeopleRequest = newFilterPeopleRequest(5, 1);
	let filterPeopleResponsePromise: Promise<FilterPeopleResponse> =
		fetchFilterPeople(filterPeopleRequest);
	let filterValue: string = '';
	let popoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	const dispatch = createEventDispatcher();

	function TeacherCreated() {
		dispatch('message', {
			created: true
		});
	}

	const form = superForm(data, {
		validators: zodClient(addTeacherSchema),
		taintedMessage: null,
		onUpdated: ({ form: f }) => {
			const message = f.message!;
			if (!message) {
				formDataPersonID.subscribe((value) => ($formData.person_id = value));
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const teacher = message.data as Teacher;
				TeacherCreated();
				return toast.success(`Profesor creado: ${teacher.ID}`);
			}
			formDataPersonID.subscribe((value) => ($formData.person_id = value));
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let formDataPersonID = writable($formData.person_id);
	formDataPersonID.subscribe((value) => ($formData.person_id = value));

	function handleOnFilterChanged() {
		filterPeopleResponsePromise = fetchOnFilterChanged(
			filterValue.trim(),
			filterPeopleRequest,
			popoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterPeopleResponsePromise = fetchOnDetailedFilter(
			filterPeopleRequest,
			popoverFilterDataMap
		).then(({ request, response }) => {
			filterPeopleRequest = request;
			filterValue = generateInitialFilterValue(filterPeopleRequest)!;
			return response;
		});
	}
</script>

<form action="?/addTeacher" use:enhance>
	<div class="flex flex-col gap-2">
		<Form.Field class="flex flex-col" {form} name="person_id">
			{#await filterPeopleResponsePromise}
				<FormFieldSkeleton />
			{:then filterPeopleResponse}
				<ServerFormSelect
					bind:filterValue
					formLabel="Persona"
					bind:popoverFilterDataMap
					comboData={GenerateComboMessagesFromPeople(filterPeopleResponse.people)}
					bind:openCombo={openPeople}
					bind:formDataID={formDataPersonID}
					on:filterChanged={handleOnFilterChanged}
					on:detailedFilter={handleOnDetailedFilter}
				/>
			{/await}
			<Form.FieldErrors />
		</Form.Field>
	</div>
	<Form.Button class="my-2 w-full">Guardar</Form.Button>
	<!-- {#if browser}
		<SuperDebug data={$formData} />
	{/if} -->
</form>
