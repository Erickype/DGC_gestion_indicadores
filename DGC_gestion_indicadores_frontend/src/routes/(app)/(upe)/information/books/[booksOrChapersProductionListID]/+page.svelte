<script lang="ts">
	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';

	import { Button } from '$lib/components/ui/button/index';

	import MoveLeft from 'lucide-svelte/icons/move-left';

	import type { Message } from '$lib/components/combobox/combobox';
	import AddModal from '$lib/components/modal/AddModal.svelte';

	import AddForm from './AddForm.svelte';

	export let data: PageServerData;

	const comboMessages: Message[][] = [data.careersData.messages];

	let addBooksOrChaptersProductionListsAuthorCareersForm =
		data.addBooksOrChaptersProductionListsAuthorCareersForm;

	addBooksOrChaptersProductionListsAuthorCareersForm.data.booksOrChaptersProductionListID =
		data.booksOrChapersProductionListID;

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			/* fetchGetAcademicProductionListsAuthorsCareersJoined(); */
		}
	}

	function returnToBooksOrChaptersProductionLists() {
		if (browser) {
			window.history.back();
		}
	}
</script>

<svelte:head>
	<title>Libros o capítulos autores</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Button variant="ghost" on:click={returnToBooksOrChaptersProductionLists}>
			<MoveLeft class="h-4 w-4" />
		</Button>
		<h2 class="text-2xl font-bold">Libros o capítulos Autores</h2>
	</div>
	<AddModal
		formComponent={AddForm}
		modalTitle="Añadir autor"
		formData={addBooksOrChaptersProductionListsAuthorCareersForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>
