<script lang="ts">
	import type { Faculty } from '$lib/api/model/api/faculty';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';

	let facultiesPromise: Promise<Faculty[]> = fetchFaculties();

	async function fetchFaculties() {
		const url = `/api/faculty`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			throw errorData;
		}
		return (facultiesPromise = response.json());
	}

	function handleCreated(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchFaculties();
		}
	}
</script>

<svelte:head>
	<title>Facultades</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Facultades</h2>
	<!-- 	<AddModal formData={addPersonFormData} on:created={handleCreated} />
 -->
</div>

<div class="mx-auto flex w-full place-content-center px-8">
	{#await facultiesPromise}
		cargando...
	{:then faculties}
		{#if faculties.length > 0}
			<!-- <PeopleTable
				formData={updatePersonFormData}
				{people}
				on:updated={handleUpdated}
				on:deleted={handleDeleted}
			></PeopleTable> -->
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay facultades registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
