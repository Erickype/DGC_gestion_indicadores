<script lang="ts">
	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';

	import { Button } from '$lib/components/ui/button/index';

	import MoveLeft from 'lucide-svelte/icons/move-left';

	import AddModal from '$lib/components/modal/AddModal.svelte';

	export let data: PageServerData;

	function fetchOnSuccess(event: CustomEvent) {
	}

	function returnToAcademicProduction() {
		if (browser) {
			window.history.back();
		}
	}
</script>

<svelte:head>
	<title>Producción academica autores</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Button variant="ghost" on:click={returnToAcademicProduction}>
			<MoveLeft class="h-4 w-4" />
		</Button>
		<h2 class="text-2xl font-bold">Producción Académica Autores</h2>
	</div>
	<!-- <AddModal
		formComponent={AddDegreeAndTeachersListsDegreeForm}
		modalTitle="Añadir autor"
		formData={addDegreeAndTeachersListsDegreeForm}
		on:created={fetchOnSuccess}
	/> -->
</div>
