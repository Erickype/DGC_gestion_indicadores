<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import type { FilterTeachersRequest, FilterTeachersResponse } from '$lib/api/model/api/teacher';
	import {
		newFilterTeachersRequest,
		fetchFilterTeachers,
		fetchOnFilterChanged,
		fetchOnDetailedFilter,
		generateInitialFilterValue,
		newPopoverFilterDataMap
	} from '$lib/components/filters/teachers/teachers';
	import { GenerateComboMessagesFromTeachers } from '$lib/api/controller/api/teacher';
	import type { PopoverFilterDataMap } from '../../table/types';

	import { writable } from 'svelte/store';

	export let formDataTeacherID = writable();

	let openTeachers = false;
	let filterTeachersRequest: FilterTeachersRequest = newFilterTeachersRequest(5, 1);
	let filterTeachersResponsePromise: Promise<FilterTeachersResponse> =
		fetchFilterTeachers(filterTeachersRequest);
	let teachersFilterValue: string = '';
	let teachersPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterTeachersResponsePromise = fetchOnFilterChanged(
			teachersFilterValue.trim(),
			filterTeachersRequest,
			teachersPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterTeachersResponsePromise = fetchOnDetailedFilter(
			filterTeachersRequest,
			teachersPopoverFilterDataMap
		).then(({ request, response }) => {
			filterTeachersRequest = request;
			teachersFilterValue = generateInitialFilterValue(filterTeachersRequest)!;
			return response;
		});
	}
</script>

{#await filterTeachersResponsePromise}
	<FormFieldSkeleton />
{:then filterTeachersResponse}
	<ServerFormSelect
		bind:filterValue={teachersFilterValue}
		formLabel="Profesor"
		bind:popoverFilterDataMap={teachersPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromTeachers(filterTeachersResponse.teachers)}
		bind:openCombo={openTeachers}
		bind:formDataID={formDataTeacherID}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
