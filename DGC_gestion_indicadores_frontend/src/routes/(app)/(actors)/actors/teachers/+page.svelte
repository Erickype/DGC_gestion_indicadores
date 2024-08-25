<script lang="ts">
	import type { FilterPeopleRequest, FilterPeopleResponse } from '$lib/api/model/api/person';
	import type { CommonError } from '$lib/api/model/errors';

	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	import PeopleTable from './Table.svelte';
	import { goto } from '$app/navigation';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import { toast } from 'svelte-sonner';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;
	const updatePersonFormData = data.updatePersonForm;

	let filterPeopleRequest: FilterPeopleRequest = {
		identity: '',
		name: '',
		lastname: '',
		email: '',
		page_size: 5,
		page: 1
	};
	let filterPeopleResponsePromise: Promise<FilterPeopleResponse> = fetchFilterPeople();

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	popoverFilterDataMap.set('identity', { label: 'Identificación', value: '' });
	popoverFilterDataMap.set('name', { label: 'Nombre', value: '' });
	popoverFilterDataMap.set('lastname', { label: 'Apellido', value: '' });
	popoverFilterDataMap.set('email', { label: 'Email', value: '' });

	function fetchOnSuccess(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchFilterPeople();
		}
	}

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

	function handleOnFilterChanged(event: CustomEvent) {
		const data: { filter: string } = event.detail;

		const filter = data.filter.trim();
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

	function handlePaginationChanged() {
		fetchFilterPeople();
	}
</script>

<svelte:head>
	<title>Personas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Personas</h2>
	<AddModal formData={addPersonFormData} on:created={fetchOnSuccess} />
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterPeopleResponsePromise}
		cargando...
	{:then filterPeopleResponse}
		{#if filterPeopleResponse.people.length > 0}
			<PeopleTable
				bind:filterPeopleRequest
				formData={updatePersonFormData}
				{filterPeopleResponse}
				bind:popoverFilterDataMap
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:detailedFilter={handleOnDetailedFilter}
				on:paginationChanged={handlePaginationChanged}
			></PeopleTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay personas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
