<script lang="ts">
	import type { Table as HeadlessTable, Column } from 'svelte-headless-table';
	import { Render, Subscribe } from 'svelte-headless-table';

	import ArrowUpDown from 'lucide-svelte/icons/arrow-up-down';
	import Search from 'lucide-svelte/icons/search';
	import X from 'lucide-svelte/icons/x';

	import PopoverFilter from './popoverFilter.svelte';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Table from '$lib/components/ui/table';
	import { createEventDispatcher } from 'svelte';

	import type { PopoverFilterDataMap } from './types';

	export let table: HeadlessTable<any, any>;
	export let columns: Column<any, any>[];
	export let filterFields: string[];
	export let tableHeightClass = 'h-[60vh]';
	export let serverItemCount: number;

	const { headerRows, pageRows, tableAttrs, tableBodyAttrs, pluginStates } =
		table.createViewModel(columns);

	const { hasNextPage, hasPreviousPage, pageIndex, pageSize } = pluginStates.page;

	const { filterValue } = pluginStates.filter;

	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();

	export let page_index = $pageIndex;
	export let filter_value = $filterValue;
	export let page_size = $pageSize;

	$: {
		filter_value = $filterValue;
	}
	$: {
		page_index = $pageIndex;
		handlePageChanged();
	}
	$: {
		page_size = $pageSize;
	}

	let isFocused = false;

	function initSearchBar(input: HTMLInputElement) {
		isFocused = true;
		input.focus();
		input.select();
	}

	function handleFocus() {
		isFocused = true;
	}

	function handleBlur() {
		isFocused = false;
	}

	function handleDeleteFilter() {
		if (filter_value !== '') {
			filter_value = '';
			return handleFilterChanged();
		}
	}

	let typingTimeout: number;
	function handleInput() {
		clearTimeout(typingTimeout);
		typingTimeout = window.setTimeout(() => {
			handleFilterChanged();
		}, 500);
	}

	const dispatch = createEventDispatcher();
	function handleFilterChanged() {
		return dispatch('filterChanged');
	}
	function handleOnDetailedFilter() {
		return dispatch('detailedFilter');
	}
	function handlePageChanged() {
		return dispatch('pageChanged');
	}
	function handlePageSizeChanged() {
		return dispatch('pageSizeChanged');
	}
</script>

<div class="flex items-center justify-between gap-2 py-4">
	<div
		class="flex w-3/4 items-center rounded-sm border px-3 pr-2 {isFocused
			? 'border-bg ring-ring ring-offset-2-primary ring-offset-background ring-2 ring-offset-2'
			: 'border-bg'}"
	>
		<Search class={'mr-2 h-4 w-4'} />
		<input
			use:initSearchBar
			class="placeholder:text-muted-foreground flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-none disabled:cursor-not-allowed disabled:opacity-50"
			placeholder="Filtrar..."
			type="text"
			bind:value={$filterValue}
			on:input={handleInput}
			on:focus={handleFocus}
			on:blur={handleBlur}
		/>
		<Button class="h-min p-1" variant="ghost" on:click={handleDeleteFilter}>
			<X class={'h-4 w-4'} />
		</Button>
	</div>

	<PopoverFilter bind:popoverFilterDataMap on:detailedFilter={handleOnDetailedFilter}
	></PopoverFilter>

	<Input
		type="number"
		placeholder="Tamaño página"
		min="1"
		max="100"
		class="w-1/4"
		bind:value={$pageSize}
		on:change={handlePageSizeChanged}
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
		{`${$pageIndex * $pageSize + $pageRows.length}/${serverItemCount}`}
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
