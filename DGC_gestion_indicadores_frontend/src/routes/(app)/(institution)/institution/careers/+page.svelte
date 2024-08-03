<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Career } from '$lib/api/model/api/career';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { PageServerData } from './$types';
	import AddForm from './AddForm.svelte';

	export let data: PageServerData;
	const addCareerForm = data.addCarrerForm;
	const facultiesComboData = data.facultiesData.messages

	let careersPromise: Promise<Career[]> = fetchCareers();

	async function fetchCareers() {
		const url = `/api/career`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return (careersPromise = response.json());
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchCareers();
		}
	}
</script>

<svelte:head>
	<title>Carreras</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Carreras</h2>
	<AddModal
		formComponent={AddForm}
		modalTitle="Crear facultad"
		formData={addCareerForm}
		comboMessages={[facultiesComboData]}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await careersPromise}
		cargando...
	{:then careers}
		{#if careers.length > 0}
			<!-- <FacultiesTable
				formData={updateFacultyFormData}
				{faculties}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></FacultiesTable> -->
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay carreras registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
