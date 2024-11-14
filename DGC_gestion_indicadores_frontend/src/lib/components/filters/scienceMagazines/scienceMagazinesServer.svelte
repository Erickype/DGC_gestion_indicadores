<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import { writable } from 'svelte/store';

	import { GenerateComboMessagesFromScienceMagazines } from '$lib/api/controller/api/academicProduction/scienceMagazines/scienceMagazine';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		ScienceMagazineJoined,
		FilterScienceMagazinesRequest,
		FilterScienceMagazinesResponse
	} from '$lib/api/model/api/academicProduction/scienceMagazines/scienceMagazine';
	import {
		fetchGetScienceMagazineFilterJoinedByScienceMagazineID,
		fetchFilteScienceMagazines,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterScienceMagazinesRequest,
		newPopoverFilterDataMap
	} from './scienceMagazines';

	export let formDataScienceMagazineID = writable();
	let getScienceMagazinePromise: Promise<ScienceMagazineJoined>;

	let openScienceMagazines = false;
	let filterScienceMagazinesRequest: FilterScienceMagazinesRequest =
		newFilterScienceMagazinesRequest(5, 1);
	let postScienceMagazinesFilter: string = '';

	getScienceMagazinePromise = fetchGetScienceMagazineFilterJoinedByScienceMagazineID(
		$formDataScienceMagazineID as string
	);

	let filterScienceMagazinesResponsePromise: Promise<FilterScienceMagazinesResponse> =
		fetchFilteScienceMagazines(filterScienceMagazinesRequest);
	let postgraduateProgramsPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterScienceMagazinesResponsePromise = fetchOnFilterChanged(
			postScienceMagazinesFilter.trim(),
			filterScienceMagazinesRequest,
			postgraduateProgramsPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterScienceMagazinesResponsePromise = fetchOnDetailedFilter(
			filterScienceMagazinesRequest,
			postgraduateProgramsPopoverFilterDataMap
		).then(({ request, response }) => {
			filterScienceMagazinesRequest = request;
			postScienceMagazinesFilter = generateInitialFilterValue(filterScienceMagazinesRequest)!;
			return response;
		});
	}

	function setInitialValue(programID: number | undefined) {
		formDataScienceMagazineID.set(programID);
	}
</script>

{#await Promise.all([filterScienceMagazinesResponsePromise, getScienceMagazinePromise])}
	<FormFieldSkeleton />
{:then [filterScienceMagazinesResponse, getScienceMagazineByIDResponse]}
	{#if getScienceMagazineByIDResponse && getScienceMagazineByIDResponse.science_magazine_id}
		{#if !filterScienceMagazinesResponse.science_magazines.some((magazine) => magazine.science_magazine_id === getScienceMagazineByIDResponse.science_magazine_id)}
			<p class="hidden">
				{filterScienceMagazinesResponse.science_magazines.unshift(getScienceMagazineByIDResponse)}
			</p>
		{/if}
	{/if}

	<ServerFormSelect
		formSelectWidth="w-[50%]"
		bind:filterValue={postScienceMagazinesFilter}
		formLabel="Programa posgrado"
		bind:popoverFilterDataMap={postgraduateProgramsPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromScienceMagazines(
			filterScienceMagazinesResponse.science_magazines
		)}
		bind:openCombo={openScienceMagazines}
		bind:formDataID={formDataScienceMagazineID}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
