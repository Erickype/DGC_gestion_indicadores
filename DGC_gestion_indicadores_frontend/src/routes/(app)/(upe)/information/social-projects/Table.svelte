<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	/* 	import type { UpdateBookOrChaptersProductionListSchema } from './schema';
	 */
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/tablePaginated.svelte';
	/* 	import UpdateForm from './UpdateForm.svelte';
	 */
	import { generateInitialFilterValue } from '$lib/components/filters/indicatorsInformation/socialProjectLists/socialProjectLists';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		SocialProjectListByAcademicPeriodJoined,
		FilterSocialProjectListsByAcademicPeriodRequest,
		FilterSocialProjectListsByAcademicPeriodResponse
	} from '$lib/api/model/api/indicatorsInformation/socialProjectLists';
	import type { Message } from '$lib/components/combobox/combobox';

	export let filterSocialProjectListsByAcademicPeriodResponse: FilterSocialProjectListsByAcademicPeriodResponse;
	let filterSocialProjectListsResponse: SocialProjectListByAcademicPeriodJoined[] =
		filterSocialProjectListsByAcademicPeriodResponse.social_project_lists;

	export let comboMessages: Message[][] | undefined;
	comboMessages = [];
	/* 	export let formData: SuperValidated<
		Infer<UpdateBookOrChaptersProductionListSchema>,
		App.Superforms.Message
	>; */
	let socialProject: SocialProjectListByAcademicPeriodJoined;

	let filterValue = '';
	let pageIndex: number = 0;
	let pageSize: number = 0;
	export let filterSocialProjectListsByAcademicPeriodRequest: FilterSocialProjectListsByAcademicPeriodRequest;
	let initialFilterValue: string | undefined = generateInitialFilterValue(
		filterSocialProjectListsByAcademicPeriodRequest
	);

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	const filterFields = [
		'publication_date',
		'is_chapter',
		'title',
		'detailed_field',
		'peer_reviewed'
	];

	const table = createTable(readable(filterSocialProjectListsResponse), {
		page: addPagination({
			initialPageSize: filterSocialProjectListsByAcademicPeriodResponse.page_size,
			initialPageIndex: filterSocialProjectListsByAcademicPeriodResponse.page - 1,
			serverSide: true,
			serverItemCount: readable(filterSocialProjectListsByAcademicPeriodResponse.count)
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			initialFilterValue: initialFilterValue?.trim(),
			serverSide: true
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'career',
			header: 'Carrera'
		}),
		table.column({
			accessor: 'name',
			header: 'Título'
		}),
		table.column({
			accessor: ({ ID }) => ID,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, {
					id: `${value}`
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
			const segments = detail.id.split('/');
			const career_id = parseInt(segments.at(segments.length - 1)!);
			socialProject = filterSocialProjectListsResponse.find(
				(socialProject) => socialProject.career_id === career_id
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
		filterSocialProjectListsByAcademicPeriodRequest.page = pageIndex + 1;
		dispatch('paginationChanged');
	}
	function handleOnPageSizeChanged() {
		filterSocialProjectListsByAcademicPeriodRequest.page_size = pageSize;
		filterSocialProjectListsByAcademicPeriodRequest.page = 1;
		dispatch('paginationChanged');
	}
</script>

<!-- <UpdateModal
	modalTitle="Actualizar libro o capítulo libro"
	{formData}
	formComponent={UpdateForm}
	{comboMessages}
	bind:updateEntity={socialProject}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/> -->

<div class="w-full">
	<Table
		tableHeightClass="h-[48vh]"
		{table}
		{columns}
		serverItemCount={filterSocialProjectListsByAcademicPeriodResponse.count}
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
