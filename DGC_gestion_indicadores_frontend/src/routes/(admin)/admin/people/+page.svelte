<script lang="ts">
	import type { Person } from '$lib/api/model/api/person';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	import PeopleTable from './Table.svelte';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;
	const updatePersonFormData = data.updatePersonForm;

	const peoplePromise: Promise<Person[]> = fetchPeople();

	async function fetchPeople() {
		const url = `/api/person`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			throw errorData;
		}
		return response.json();
	}

	function handleCreated(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchPeople();
		}
	}

	function handleDeleted(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchPeople();
		}
	}

	function handleUpdated(event: any) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchPeople();
		}
	}
</script>

<svelte:head>
	<title>Personas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Personas</h2>
	<AddModal formData={addPersonFormData} on:created={handleCreated} />
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await peoplePromise}
		cargando...
	{:then people}
		{#if people.length > 0}
			<PeopleTable
				formData={updatePersonFormData}
				{people}
				on:updated={handleUpdated}
				on:deleted={handleDeleted}
			></PeopleTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay personas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
