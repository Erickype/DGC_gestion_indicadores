<script lang="ts">
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;

	async function fetchPeople() {
		console.log('People data fetch');
	}

	function handleCreated(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
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
