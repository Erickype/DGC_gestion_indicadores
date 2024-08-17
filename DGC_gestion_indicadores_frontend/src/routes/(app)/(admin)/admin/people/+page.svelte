<script lang="ts">
	import type { FilterPeopleRequest, FilterPeopleResponse } from '$lib/api/model/api/person';
	import type { CommonError } from '$lib/api/model/errors';

	import { Badge } from '$lib/components/ui/badge/index.js';
	import X from 'lucide-svelte/icons/x';

	import Alert from '$lib/components/alert/alert.svelte';
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	import { flip } from 'svelte/animate';
	import { dndzone } from 'svelte-dnd-action';
	import type { DragableFilterFields } from '$lib/utils';

	import PeopleTable from './Table.svelte';
	import { goto } from '$app/navigation';

	export let data: PageData;
	const addPersonFormData = data.addPersonForm;
	const updatePersonFormData = data.updatePersonForm;

	let filterPeopleRequest: FilterPeopleRequest = {
		identity: '',
		name: '',
		lastname: '',
		email: '',
		page_size: 5,
		page: 1
	};
	let filterPeopleResponsePromise: Promise<FilterPeopleResponse> = fetchFilterPeople();

	let itemsLabelsMap = new Map<string, string>([
		['identity', 'Identificación'],
		['name', 'Nombre'],
		['lastname', 'Apellido'],
		['email', 'Email']
	]);
	let items: DragableFilterFields[] = [];

	const flipDurationMs = 300;

	function fetchOnSuccess(event: any) {
		const data: { status: boolean } = event.detail;
		if (data.status) {
			fetchFilterPeople();
		}
	}

	async function fetchFilterPeople() {
		const url = `/api/people/filter`;
		const response = await fetch(url, {
			method: 'POST',
			body: JSON.stringify(filterPeopleRequest)
		});
		if (!response.ok) {
			const errorData = (await response.json()) as CommonError;
			if (response.status === 401) {
				throw goto('/');
			}
			throw errorData;
		}
		return (filterPeopleResponsePromise = response.json());
	}

	function handleDndConsider(e: any) {
		items = e.detail.items;
	}
	function handleDndFinalize(e: any) {
		items = e.detail.items;
	}

	function findItemWithGreatestId(items: DragableFilterFields[]): DragableFilterFields | undefined {
		return items.reduce(
			(maxItem, currentItem) => (currentItem.id > maxItem.id ? currentItem : maxItem),
			items[0]
		);
	}

	function handleOnFilterChanged(event: CustomEvent) {
		if (items.length >= itemsLabelsMap.size) return; // Prevent adding more items than available in the map
		const data: { filter: string } = event.detail;

		const filter = data.filter.trim();

		// Find the next unique key that has not been used in the items array
		let mapEntry: [string, string] | undefined;
		for (let [key, label] of itemsLabelsMap) {
			if (!items.find((item) => item.label === label)) {
				mapEntry = [key, label];
				break;
			}
		}

		if (!mapEntry) return; // Exit if no unique keys are available

		const [key, label] = mapEntry;

		// Create a new item with a unique id and push it to the items array
		const id = items.length <= 0 ? 1 : findItemWithGreatestId(items)!.id + 1;
		items.push({ id: id, key: key, name: filter, label: label });

		// Force the update to the reactive items array
		items = [...items];
		filterPeopleRequest = mapItemsToFilterRequest(
			items,
			filterPeopleRequest.page,
			filterPeopleRequest.page_size
		);
		fetchFilterPeople();
	}

	function mapItemsToFilterRequest(
		items: DragableFilterFields[],
		page: number,
		page_size: number
	): FilterPeopleRequest {
		let request: FilterPeopleRequest = { page: page, page_size: page_size };
		items.forEach((item) => {
			(request as any)[item.key] = item.name;
		});

		return request;
	}

	function handleRemoveFilter(id: number) {
		const index = items.findIndex((item) => item.id === id);
		items.splice(index, 1);

		items = items;
		filterPeopleRequest = mapItemsToFilterRequest(
			items,
			filterPeopleRequest.page,
			filterPeopleRequest.page_size
		);
		fetchFilterPeople();
	}

	function handlePaginationChanged() {
		fetchFilterPeople();
	}
</script>

<svelte:head>
	<title>Personas</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Personas</h2>
	<AddModal formData={addPersonFormData} on:created={fetchOnSuccess} />
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
	{#await filterPeopleResponsePromise}
		cargando...
	{:then filterPeopleResponse}
		<section
			class="flex gap-2 {filterPeopleResponse.people.length <= 0
				? 'pb-2'
				: ''} rounded-sm px-2 py-1"
			use:dndzone={{
				items,
				flipDurationMs,
				dropTargetStyle: { outline: 'hsl(var(--border) / var(--tw-border-opacity)) solid 2px' }
			}}
			on:consider={handleDndConsider}
			on:finalize={handleDndFinalize}
		>
			{#each items as item (item.id)}
				<div
					class="m-0 flex min-w-[100px] flex-col justify-between p-0"
					animate:flip={{ duration: flipDurationMs }}
				>
					<div class="min-w-[100px] text-xs font-extralight">{item.label}</div>
					<Badge class="flex min-w-[100px] justify-between" variant="secondary">
						{item.name}
						<button
							on:click={() => {
								handleRemoveFilter(item.id);
							}}
						>
							<X class="ml-2 h-4 w-4" />
						</button>
					</Badge>
				</div>
			{/each}
		</section>

		{#if filterPeopleResponse.people.length > 0}
			<PeopleTable
				bind:filterPeopleRequest
				formData={updatePersonFormData}
				{filterPeopleResponse}
				on:updated={fetchOnSuccess}
				on:deleted={fetchOnSuccess}
				on:filterChanged={handleOnFilterChanged}
				on:paginationChanged={handlePaginationChanged}
			></PeopleTable>
		{:else}
			<Alert title="Sin registros" description={'Ups, no hay personas registradas.'} />
		{/if}
	{:catch error}
		<Alert variant="destructive" description={`${error.Detail}: ${error.Message}`} />
	{/await}
</div>
