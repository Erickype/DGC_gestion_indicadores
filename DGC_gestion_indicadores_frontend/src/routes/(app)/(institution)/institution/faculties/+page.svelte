<script lang="ts">
	import type { Faculty } from '$lib/api/model/api/faculty';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';

	import Icon from "lucide-svelte/icons/rows-3";

	import AddModal from '$lib/components/modal/AddModal.svelte';
	import FacultiesTable from './Table.svelte';
	import AddForm from './AddForm.svelte';
	import { goto } from '$app/navigation';

	export let data: PageData;
	const addFacultyForm = data.addFacultyForm;
	const updateFacultyFormData = data.updateFacultyForm;

	let facultiesPromise: Promise<Faculty[]> = fetchFaculties();

	async function fetchFaculties() {
		const url = `/api/faculty`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData
		}
		return (facultiesPromise = response.json());
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchFaculties();
		}
	}
</script>

<svelte:head>
	<title>Facultades</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Facultades</h2>
	</div>
	<AddModal
		formComponent={AddForm}
		modalTitle="Crear facultad"
		formData={addFacultyForm}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await facultiesPromise}
		cargando...
	{:then faculties}
		{#if faculties.length > 0}
			<FacultiesTable
				formData={updateFacultyFormData}
				{faculties}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></FacultiesTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay facultades registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
