<script lang="ts">
	import SuperDebug, { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';
	import { addTeacherSchema, type AddTeacherSchema } from './schema';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import type { CommonError } from '$lib/api/model/errors';

	import * as Form from '$lib/components/ui/form';
	import {
		manageToastFromErrorMessageOnAddForm,
		manageToastFromInvalidAddForm
	} from '$lib/utils.js';

	import { createEventDispatcher } from 'svelte';
	import { browser } from '$app/environment';
	import { toast } from 'svelte-sonner';

	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import type { Teacher } from '$lib/api/model/api/teacher';
	import { goto } from '$app/navigation';
	import type { FilterPeopleRequest, FilterPeopleResponse } from '$lib/api/model/api/person';
	import Alert from '$lib/components/alert/alert.svelte';
	import { GenerateComboMessagesFromPeople } from '$lib/api/controller/api/person';
	import { writable } from 'svelte/store';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';

	export let data: SuperValidated<Infer<AddTeacherSchema>, App.Superforms.Message>;

	let filterPeopleRequest: FilterPeopleRequest = {
		identity: '',
		name: '',
		lastname: '',
		email: '',
		page_size: 5,
		page: 1
	};
	let filterPeopleResponsePromise: Promise<FilterPeopleResponse> = fetchFilterPeople();

	async function fetchFilterPeople() {
		const url = `/api/people/filter`;
		const response = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(filterPeopleRequest)
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return (filterPeopleResponsePromise = response.json());
	}

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
				return manageToastFromInvalidAddForm();
			}
			if (message.success) {
				const teacher = message.data as Teacher;
				TeacherCreated();
				return toast.success(`Profesor creado: ${teacher.ID}`);
			}
			return manageToastFromErrorMessageOnAddForm(message);
		}
	});

	const { form: formData, enhance } = form;

	let openPeople = false;

	let formDataPersonID = writable($formData.person_id);
	formDataPersonID.subscribe((value) => ($formData.person_id = value));

	let filterValue: string = '';

	let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	popoverFilterDataMap.set('identity', { label: 'Identificación', value: '' });
	popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
	popoverFilterDataMap.set('lastname', { label: 'Apellido', value: '' });
	popoverFilterDataMap.set('email', { label: 'Email', value: '' });

	function handleOnFilterChanged() {
		const filter = filterValue.trim();
		popoverFilterDataMap.forEach((_, key) => {
			(filterPeopleRequest as any)[key] = filter;
		});

		fetchFilterPeople().then((response: FilterPeopleResponse) => {
			if (response.count === 0) {
				toast.warning(`No hay datos para el filtro: ${filter}`);
				popoverFilterDataMap.forEach((_, key) => {
					(filterPeopleRequest as any)[key] = '';
				});
				fetchFilterPeople();
			}
		});
	}

	function handleOnDetailedFilter() {
		let request: FilterPeopleRequest = {
			page: filterPeopleRequest.page,
			page_size: filterPeopleRequest.page_size
		};
		popoverFilterDataMap.forEach((item, key) => {
			if (item.value !== '') {
				(request as any)[key] = item.value;
			}
		});
		filterPeopleRequest = request;

		fetchFilterPeople().then((response: FilterPeopleResponse) => {
			if (response.count === 0) {
				let message = 'No hay datos para el filtro\n';
				popoverFilterDataMap.forEach((item, _) => {
					if (item.value !== '') {
						message += `${item.label}: ${item.value}; `;
					}
				});
				message = message.slice(0, message.length - 2);
				toast.warning(message);
				popoverFilterDataMap.forEach((_, key) => {
					(filterPeopleRequest as any)[key] = '';
				});
				fetchFilterPeople();
			}
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
