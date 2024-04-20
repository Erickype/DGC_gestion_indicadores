<script lang="ts">
	import ArrowUpDown from 'lucide-svelte/icons/arrow-up-down';

	import type { Teacher } from '$lib/api/model/api/teacher';
	import DataTableActions from './teachers-data-table-actions.svelte';

	import { addPagination, addSortBy, addTableFilter } from 'svelte-headless-table/plugins';
	import { createTable, Render, Subscribe, createRender } from 'svelte-headless-table';
	import { readable } from 'svelte/store';

	import { Button } from '$lib/components/ui/button';
	import * as Table from '$lib/components/ui/table';
	import { Input } from '$lib/components/ui/input/index.js';

	export let teachers: Teacher[];

	const table = createTable(readable(teachers), {
		page: addPagination({
			initialPageSize: 4
		}),
		sort: addSortBy(),
		filter: addTableFilter({
			fn: ({ filterValue, value }) => value.toLowerCase().includes(filterValue.toLowerCase())
		})
	});

	const columns = table.createColumns([
		table.column({
			accessor: 'ID',
			header: 'ID'
		}),
		table.column({
			accessor: 'person_id',
			header: 'Nombre'
		}),
		table.column({
			accessor: 'career_id',
			header: 'Carrera'
		}),
		table.column({
			accessor: 'dedication_id',
			header: 'Dedicación'
		}),
		table.column({
			accessor: 'scaled_grade_id',
			header: 'Grado Esc.'
		}),
		table.column({
			accessor: ({ ID }) => ID,
			header: '',
			cell: ({ value }) => {
				return createRender(DataTableActions, { id: value.toString() });
			}
		})
	]);

	const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } =
		table.createViewModel(columns);

	const { hasNextPage, hasPreviousPage, pageIndex, pageSize } = pluginStates.page;

	const { filterValue } = pluginStates.filter;
</script>

<div>
	<div class="flex items-center gap-2 py-4">
		<Input class="w-3/4" placeholder="Filtrar..." type="text" bind:value={$filterValue} />
		<Input
			type="number"
			placeholder="Tamaño página"
			min="1"
			max="10"
			class="w-1/4"
			bind:value={$pageSize}
		/>
	</div>
	<div class="rounded-md border">
		<Table.Root {...$tableAttrs}>
			<Table.Header>
				{#each $headerRows as headerRow}
					<Subscribe rowAttrs={headerRow.attrs()}>
						<Table.Row>
							{#each headerRow.cells as cell (cell.id)}
								<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
									<Table.Head {...attrs}>
										{#if cell.id === 'person_id'}
											<Button variant="ghost" on:click={props.sort.toggle}>
												<Render of={cell.render()} />
												<ArrowUpDown class={'ml-2 h-4 w-4'} />
											</Button>
										{:else}
											<Render of={cell.render()} />
										{/if}
									</Table.Head>
								</Subscribe>
							{/each}
						</Table.Row>
					</Subscribe>
				{/each}
			</Table.Header>
			<Table.Body {...$tableBodyAttrs}>
				{#each $pageRows as row (row.id)}
					<Subscribe rowAttrs={row.attrs()} let:rowAttrs>
						<Table.Row {...rowAttrs}>
							{#each row.cells as cell (cell.id)}
								<Subscribe attrs={cell.attrs()} let:attrs>
									<Table.Cell {...attrs}>
										<Render of={cell.render()} />
									</Table.Cell>
								</Subscribe>
							{/each}
						</Table.Row>
					</Subscribe>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
	<div class="flex items-center justify-end space-x-4 py-4">
		<Button
			variant="outline"
			size="sm"
			on:click={() => ($pageIndex = $pageIndex - 1)}
			disabled={!$hasPreviousPage}>Anterior</Button
		>
		<Button
			variant="outline"
			size="sm"
			disabled={!$hasNextPage}
			on:click={() => ($pageIndex = $pageIndex + 1)}>Siguiente</Button
		>
	</div>
</div>
