<script lang="ts">
	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import DataTableActions from '$lib/components/table/tableActions.svelte';
	import UpdateModal from '$lib/components/modal/UpdateModal.svelte';
	import Table from '$lib/components/table/table.svelte';
	import { toast } from 'svelte-sonner';

	import type { Infer, SuperValidated } from 'sveltekit-superforms';
	import type { UpdateUserSchema } from './schema';
	import { createEventDispatcher } from 'svelte';

	import type { User } from '$lib/api/model/admin/user';
	import UpdateForm from './UpdateForm.svelte';

	export let users: User[];
	export let formData: SuperValidated<Infer<UpdateUserSchema>>;
	let user: User;

	const filterFields = ['username', 'email', 'role_id'];

	const table = createTable(readable(users), {
		page: addPagination({
			initialPageSize: 10
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			fn: ({ filterValue, value }) => value.toLowerCase().includes(filterValue.toLowerCase())
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'username',
			header: 'Usuario'
		}),
		table.column({
			accessor: 'email',
			header: 'Correo'
		}),
		table.column({
			accessor: 'role_id',
			header: 'Role',
			cell: ({ value }) => {
				const roleID = value;
				let role = '';
				switch (roleID) {
					case 1:
						role = 'Administrador';
						break;
					case 2:
						role = 'Usuario UPE';
						break;
					case 3:
						role = 'Autoridad';
						break;
				}
				return role;
			}
		}),
		table.column({
			accessor: ({ ID }) => ID,
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
		const detail: { status: boolean; id: number } = event.detail;
		if (detail.status) {
			user = users.find((period) => period.ID === detail.id)!;
			updateFormOpen = true;
		} else {
			updateFormOpen = false;
		}
	}

	async function deleteUsers(id: string) {}

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
	modalTitle="Actualizar rol de usuario"
	{formData}
	formComponent={UpdateForm}
	bind:updateEntity={user}
	bind:open={updateFormOpen}
	on:updated={handleUpdated}
/>

<div class="w-full">
	<Table {table} {columns} {filterFields} itemCount={users.length} />
</div>
