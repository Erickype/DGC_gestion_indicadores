<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import { writable } from 'svelte/store';

	import { GenerateComboMessagesFromAcademicPeriods } from '$lib/api/controller/view/academicPeriod';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import {
		newFilterAcademicPeriodsRequest,
		generateInitialFilterValue,
		newPopoverFilterDataMap,
		fetchFilterAcademicPeriods,
		fetchOnFilterChanged,
		fetchOnDetailedFilter
	} from '$lib/components/filters/academicPeriods/academicPeriods';
	import type {
		FilterAcademicPeriodsRequest,
		FilterAcademicPeriodsResponse
	} from '$lib/api/model/view/academicPeriod';

	export let formDataAcademicPeriodID = writable();

	let openAcademicPeriods = false;
	let filterAcademicPeriodsRequest: FilterAcademicPeriodsRequest = newFilterAcademicPeriodsRequest(
		5,
		1
	);
	let academicPeriodsFilterValue: string = '';

	let filterAcademicPeriodsPromise: Promise<FilterAcademicPeriodsResponse> =
		fetchFilterAcademicPeriods(filterAcademicPeriodsRequest);
	let academicPeriodsPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterAcademicPeriodsPromise = fetchOnFilterChanged(
			academicPeriodsFilterValue.trim(),
			filterAcademicPeriodsRequest,
			academicPeriodsPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterAcademicPeriodsPromise = fetchOnDetailedFilter(
			filterAcademicPeriodsRequest,
			academicPeriodsPopoverFilterDataMap
		).then(({ request, response }) => {
			filterAcademicPeriodsRequest = request;
			academicPeriodsFilterValue = generateInitialFilterValue(filterAcademicPeriodsRequest)!;
			return response;
		});
	}
</script>

{#await Promise.all([filterAcademicPeriodsPromise])}
	<FormFieldSkeleton />
{:then [filterAcademicPeriodsResponse]}
	<ServerFormSelect
		bind:filterValue={academicPeriodsFilterValue}
		formLabel="Periodo académico"
		bind:popoverFilterDataMap={academicPeriodsPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromAcademicPeriods(
			filterAcademicPeriodsResponse.academic_periods
		)}
		bind:openCombo={openAcademicPeriods}
		bind:formDataID={formDataAcademicPeriodID}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
