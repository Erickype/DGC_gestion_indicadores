<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import { writable } from 'svelte/store';

	import { GenerateComboMessagesCohorts } from '$lib/api/controller/api/indicatorsInformation/postgraduate';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		FilterCohortListsRequest,
		FilterCohortListsResponse,
		CohortList
	} from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import {
		fetchGetCohortByYear,
		fetchFilterCohorts,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterCohortListsRequest,
		newPopoverFilterDataMap
	} from './cohorts';

	export let formDataCohortYear = writable();
	let getCohortByYearPromise: Promise<CohortList>;

	let openPostgraduatePrograms = false;
	let filterCohortListsRequest: FilterCohortListsRequest = newFilterCohortListsRequest(5, 1);
	let postCohortsFilter: string = '';

	getCohortByYearPromise = fetchGetCohortByYear($formDataCohortYear as string);

	let filterCohortListsResponsePromise: Promise<FilterCohortListsResponse> =
		fetchFilterCohorts(filterCohortListsRequest);
	let postCohortsPopoverFilterMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterCohortListsResponsePromise = fetchOnFilterChanged(
			postCohortsFilter.trim(),
			filterCohortListsRequest,
			postCohortsPopoverFilterMap
		);
		getCohortByYearPromise = fetchGetCohortByYear($formDataCohortYear as string);
	}

	async function handleOnDetailedFilter() {
		filterCohortListsResponsePromise = fetchOnDetailedFilter(
			filterCohortListsRequest,
			postCohortsPopoverFilterMap
		).then(({ request, response }) => {
			filterCohortListsRequest = request;
			postCohortsFilter = generateInitialFilterValue(filterCohortListsRequest)!;
			return response;
		});
		getCohortByYearPromise = fetchGetCohortByYear($formDataCohortYear as string);
	}

	function setInitialValue(year: number | undefined) {
		formDataCohortYear.set(year);
	}
</script>

{#await Promise.all([filterCohortListsResponsePromise, getCohortByYearPromise])}
	<FormFieldSkeleton />
{:then [filterCohortListsResponse, cohortByYearResponse]}
	{#if cohortByYearResponse && cohortByYearResponse.year}
		{#if !filterCohortListsResponse.cohort_lists.some((cohort) => cohort.year === cohortByYearResponse.year)}
			<p class="hidden">
				{filterCohortListsResponse.cohort_lists.unshift(cohortByYearResponse)}
			</p>
		{/if}
	{:else}
		<p class="hidden">
			{setInitialValue(filterCohortListsResponse.cohort_lists.at(0)?.year)}
		</p>
	{/if}

	<ServerFormSelect
		formSelectWidth="w-[25%]"
		bind:filterValue={postCohortsFilter}
		formLabel="Cohorte"
		bind:popoverFilterDataMap={postCohortsPopoverFilterMap}
		comboData={GenerateComboMessagesCohorts(filterCohortListsResponse.cohort_lists)}
		bind:openCombo={openPostgraduatePrograms}
		bind:formDataID={formDataCohortYear}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
