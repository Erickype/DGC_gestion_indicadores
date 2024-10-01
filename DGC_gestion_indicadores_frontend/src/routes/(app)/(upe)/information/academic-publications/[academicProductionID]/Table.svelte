<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { createEventDispatcher } from 'svelte';
	import { readable } from 'svelte/store';

	import type { UpdateAcademicProductionListsAuthorCareersSchema } from './schema';
	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/table.svelte';
	import UpdateForm from './UpdateForm.svelte';

	import type { AcademicProductionListsAuthorsCareersJoined } from '$lib/api/model/api/indicatorsInformation/academicProductionListsAuthor';
	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { Message } from '$lib/components/combobox/combobox';

	export let comboMessages: Message[][] | undefined = undefined;
	export let authorsCareers: AcademicProductionListsAuthorsCareersJoined[];
	export let formData: SuperValidated<
		Infer<UpdateAcademicProductionListsAuthorCareersSchema>,
		App.Superforms.Message
	>;
	let authorCareer: AcademicProductionListsAuthorsCareersJoined;

	const filterFields = ['author', 'careers'];

	const table = createTable(readable(authorsCareers), {
		page: addPagination({
			initialPageSize: 30
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			fn: ({ filterValue, value }) => value.toLowerCase().includes(filterValue.toLowerCase())
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'author',
			header: 'Autor'
		}),
		table.column({
			accessor: 'careers',
			header: 'Carreras',
			cell: ({ value }) => {
				let label = '';
				value.map((career) => {
					label += career.name + ', ';
				});
				label = label.slice(0, label.length - 2);
				return label;
			}
		}),
		table.column({
			accessor: ({ author_id }) => author_id,
			header: '',
			cell: ({ value }) => {
				const actions = createRender(DataTableActions, { id: value.toString() });
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
			authorCareer = authorsCareers.find(
				(authorCareer) => authorCareer.author_id.toString() === detail.id
			)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteCareer(id: string) {}

	function handleUpdated(event: any) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			dispatch('updated', {
				status: true
			});
		}
	}
</script>

<UpdateModal
	modalTitle="Actualizar información de autor"
	{formData}
	{comboMessages}
	formComponent={UpdateForm}
	bind:updateEntity={authorCareer}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table {table} {columns} {filterFields} itemCount={authorsCareers.length} />
</div>
