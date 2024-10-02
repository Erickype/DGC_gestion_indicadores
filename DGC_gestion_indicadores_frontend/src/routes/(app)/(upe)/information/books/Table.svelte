<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateBookOrChaptersProductionListSchema } from './schema';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';
	/* 	import UpdateForm from './UpdateForm.svelte';
	 */
	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/booksOrChaptersProductionLists/booksOrChaptersProductionLists';
	import type {
		BooksOrChaptersProductionListsByAcademicPeriodJoined,
		FilterBooksOrChaptersProductionListsByAcademicPeriodRequest,
		FilterBooksOrChaptersProductionListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/booksOrChaptersProductionLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type { Message } from '$lib/components/combobox/combobox';

	export let filterBooksOrChaptersProductionListsByAcademicPeriodResponse: FilterBooksOrChaptersProductionListsByAcademicPeriodResponse;
	let filterBooksOrChaptersProductionListsResponse: BooksOrChaptersProductionListsByAcademicPeriodJoined[] =
		filterBooksOrChaptersProductionListsByAcademicPeriodResponse.books_or_chapters_production_lists;

	export let comboMessages: Message[][] | undefined;
	comboMessages = [];
	export let formData: SuperValidated<
		Infer<UpdateBookOrChaptersProductionListSchema>,
		App.Superforms.Message
	>;
	let academicProduction: BooksOrChaptersProductionListsByAcademicPeriodJoined;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterBooksOrChaptersProductionListsByAcademicPeriodRequest: FilterBooksOrChaptersProductionListsByAcademicPeriodRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterBooksOrChaptersProductionListsByAcademicPeriodRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = [
		'publication_date',
		'is_chapter',
		'title',
		'detailed_field',
		'peer_reviewed'
	];

	const table = createTable(readable(filterBooksOrChaptersProductionListsResponse), {
		page: addPagination({
			initialPageSize: filterBooksOrChaptersProductionListsByAcademicPeriodResponse.page_size,
			initialPageIndex: filterBooksOrChaptersProductionListsByAcademicPeriodResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterBooksOrChaptersProductionListsByAcademicPeriodResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'publication_date',
			header: 'Fecha',
			cell: ({ value }) => {
				const date = new Date(value);
				const formattedDate = date.toLocaleDateString('es-EC', {
					year: 'numeric',
					month: 'numeric',
					day: 'numeric',
					timeZone: 'UTC'
				});
				return formattedDate;
			}
		}),
		table.column({
			accessor: 'is_chapter',
			header: 'Tipo publicación',
			cell: ({ value }) => {
				if (value) {
					return 'Capítulo libro';
				}
				return 'Libro';
			}
		}),
		table.column({
			accessor: 'title',
			header: 'Título'
		}),
		table.column({
			accessor: 'peer_reviewed',
			header: 'Revisado pares',
			cell: ({ value }) => {
				if (value) {
					return 'Sí';
				}
				return 'No';
			}
		}),
		table.column({
			accessor: ({ ID }) => ID,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${value}`,
					moreAction: true,
					moreActionURL: `/`
				});
				actions.on('delete-confirmation', handleDeleteConfirmation);
				actions.on('update-action', handleUpdateAction);
				return actions;
			}
		})
	]);

	const dispatch = createEventDispatcher();
	async function handleDeleteConfirmation(event: any) {}

	let updateFormOpen = false;
	function handleUpdateAction(event: any) {
		const detail: { status: boolean; id: string } = event.detail;
		if (detail.status) {
			academicProduction = filterBooksOrChaptersProductionListsResponse.find(
				(academicProduction) => academicProduction.ID.toString() === detail.id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteTeacher(id: string) {}

	function handleUpdated(event: any) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			dispatch('updated', {
				status: true
			});
		}
	}
	function handleOnFilterChanged() {
		dispatch('filterChanged', {
			filter: filterValue
		});
	}
	function handleOnDetailedFilter() {
		dispatch('detailedFilter');
	}
	function handleOnPageChanged() {
		filterBooksOrChaptersProductionListsByAcademicPeriodRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterBooksOrChaptersProductionListsByAcademicPeriodRequest.page_size = pageSize;
		filterBooksOrChaptersProductionListsByAcademicPeriodRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<!-- <UpdateModal
	modalTitle="Actualizar publicación académica"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={academicProduction}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass="h-[50vh]"
		{table}
		{columns}
		serverItemCount={filterBooksOrChaptersProductionListsByAcademicPeriodResponse.count}
		{filterFields}
		bind:popoverFilterDataMap
		bind:filter_value={filterValue}
		bind:page_index={pageIndex}
		bind:page_size={pageSize}
		on:filterChanged={handleOnFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
		on:pageChanged={handleOnPageChanged}
		on:pageSizeChanged={handleOnPageSizeChanged}
	/>
</div>
