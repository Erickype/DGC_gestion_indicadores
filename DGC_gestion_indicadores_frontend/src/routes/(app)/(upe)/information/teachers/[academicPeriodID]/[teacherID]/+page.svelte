<script lang="ts">
	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';

	import AddTeachersListsDegreeForm from './AddTeachersListsDegreeForm.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';

	import { Button } from '$lib/components/ui/button/index';

	import MoveLeft from 'lucide-svelte/icons/move-left';

	import type { Message } from '$lib/components/combobox/combobox';

	export let data: PageServerData;
	const addTeachersListsDegreeForm = data.addTeachersListsDegreeForm;

	const comboMessages: Message[][] = [data.degreeLevelsData.messages];

	function returnToTeachers() {
		if (browser) {
			window.history.back();
		}
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			//fetchCareers();
		}
	}
</script>

<svelte:head>
	<title>Títulos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Button variant="ghost" on:click={returnToTeachers}>
			<MoveLeft class="h-4 w-4" />
		</Button>
		<h2 class="text-2xl font-bold">Títulos</h2>
	</div>
	<AddModal
		formComponent={AddTeachersListsDegreeForm}
		modalTitle="Crear título"
		formData={addTeachersListsDegreeForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>
