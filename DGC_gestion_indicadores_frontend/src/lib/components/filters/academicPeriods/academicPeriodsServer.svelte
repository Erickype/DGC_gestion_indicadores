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
		fetchFilterAcademicPeriods
	} from '$lib/components/filters/academicPeriods/academicPeriods';
	import type {
		FilterAcademicPeriodsRequest,
		FilterAcademicPeriodsResponse
	} from '$lib/api/model/view/academicPeriod';

	export let formDataAcademicPeriodID = writable();

	let openAcademicPeriods = false;
	let filterTeachersRequest: FilterAcademicPeriodsRequest = newFilterAcademicPeriodsRequest(5, 1);
	let academicPeriodsFilterValue: string = '';

	let filterAcademicPeriodsPromise: Promise<FilterAcademicPeriodsResponse> =
		fetchFilterAcademicPeriods(filterTeachersRequest);
	let academicPeriodsPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		/* filterTeachersResponsePromise = fetchOnFilterChanged(
			teachersFilterValue.trim(),
			filterTeachersRequest,
			teachersPopoverFilterDataMap
		); */
	}

	async function handleOnDetailedFilter() {
		/* filterTeachersResponsePromise = fetchOnDetailedFilter(
			filterTeachersRequest,
			teachersPopoverFilterDataMap
		).then(({ request, response }) => {
			filterTeachersRequest = request;
			teachersFilterValue = generateInitialFilterValue(filterTeachersRequest)!;
			return response;
		}); */
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
