<script lang="ts">
	import type { Table as HeadlessTable, Column } from 'svelte-headless-table';
	import { Render, Subscribe } from 'svelte-headless-table';

	import ArrowUpDown from 'lucide-svelte/icons/arrow-up-down';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Table from '$lib/components/ui/table';

	export let table: HeadlessTable<any, any>;
	export let columns: Column<any, any>[];
	export let filterFields: string[];
	export let tableHeightClass = 'h-[60vh]';
	export let itemCount: number;

	const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } =
		table.createViewModel(columns);

	const { hasNextPage, hasPreviousPage, pageIndex, pageSize } = pluginStates.page;

	const { filterValue } = pluginStates.filter;
</script>

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
<div class="flex rounded-md border {tableHeightClass} overflow-auto">
	<ScrollArea class="w-full rounded-md border">
		<Table.Root {...$tableAttrs}>
			<Table.Header>
				{#each $headerRows as headerRow}
					<Subscribe rowAttrs={headerRow.attrs()}>
						<Table.Row>
							{#each headerRow.cells as cell (cell.id)}
								<Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
									<Table.Head {...attrs}>
										{#if filterFields.find((field) => cell.id === field)}
											<Button variant="ghost" on:click={props.sort.toggle}>
												<Render of={cell.render()} />
												<ArrowUpDown class={'stroke-primary ml-2 h-4 w-4'} />
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
	</ScrollArea>
</div>
<div class="flex items-center justify-end space-x-4 py-4">
	<dir class="m-0 p-0">
		{`${$pageIndex * $pageSize + $pageRows.length}/${itemCount}`}
	</dir>
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
