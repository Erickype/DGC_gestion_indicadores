<script lang="ts">
	import type { PageServerData } from './$types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';

	import { Button } from '$lib/components/ui/button/index';

	import MoveLeft from 'lucide-svelte/icons/move-left';

	import TableSkeleton from '$lib/components/skeleton/table.svelte';
	import type { Message } from '$lib/components/combobox/combobox';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import type { CommonError } from '$lib/api/model/errors';
	import Alert from '$lib/components/alert/alert.svelte';

	import type { BooksOrChaptersProductionListsAuthorsCareersJoined } from '$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionListsAuthor';
	import AuthorsCareersTable from './Table.svelte';
	import AddForm from './AddForm.svelte';

	export let data: PageServerData;

	const comboMessages: Message[][] = [data.careersData.messages];

	let addBooksOrChaptersProductionListsAuthorCareersForm =
		data.addBooksOrChaptersProductionListsAuthorCareersForm;
	let updateBooksOrChaptersProductionListsAuthorCareersForm =
		data.updateBooksOrChaptersProductionListsAuthorCareersForm;

	let authorsCareersPromise: Promise<BooksOrChaptersProductionListsAuthorsCareersJoined[]> =
		fetchGetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID();

	addBooksOrChaptersProductionListsAuthorCareersForm.data.booksOrChaptersProductionListID =
		data.booksOrChapersProductionListID;
	updateBooksOrChaptersProductionListsAuthorCareersForm.data.booksOrChaptersProductionListID =
		data.booksOrChapersProductionListID;

	async function fetchGetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID() {
		const url = `/api/indicatorsInformation/booksOrChaptersProductionListsAuthors/JoinedByBooksOrChaptersProductionListID/${data.booksOrChapersProductionListID}`;
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
		return (authorsCareersPromise = response.json());
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchGetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID();
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

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await authorsCareersPromise}
		<TableSkeleton tableHeightClass="h-[65vh]" />
	{:then authorsCareers}
		{#if authorsCareers.length > 0}
			<AuthorsCareersTable
				{authorsCareers}
				{comboMessages}
				formData={updateBooksOrChaptersProductionListsAuthorCareersForm}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
			></AuthorsCareersTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay autores registrados.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={error.message} />
	{/await}
</div>
