<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import type {
		FilterTeachersRequest,
		FilterTeachersResponse,
		TeacherPerson
	} from '$lib/api/model/api/teacher';
	import {
		newFilterTeachersRequest,
		fetchFilterTeachers,
		fetchOnFilterChanged,
		fetchOnDetailedFilter,
		generateInitialFilterValue,
		newPopoverFilterDataMap,
		fetchTeacherPersonJoinedByTeacherID
	} from '$lib/components/filters/teachers/teachers';
	import { GenerateComboMessagesFromTeachers } from '$lib/api/controller/api/teacher';
	import type { PopoverFilterDataMap } from '../../table/types';

	import { writable } from 'svelte/store';

	export let formDataTeacherID = writable();
	let teacherPerson: Promise<TeacherPerson>;

	let openTeachers = false;
	let filterTeachersRequest: FilterTeachersRequest = newFilterTeachersRequest(5, 1);
	let teachersFilterValue: string = '';

	teacherPerson = fetchTeacherPersonJoinedByTeacherID($formDataTeacherID as string);

	let filterTeachersResponsePromise: Promise<FilterTeachersResponse> =
		fetchFilterTeachers(filterTeachersRequest);
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

{#await Promise.all([filterTeachersResponsePromise, teacherPerson])}
	<FormFieldSkeleton />
{:then [filterTeachersResponse, resolvedTeacherPerson]}
	{#if resolvedTeacherPerson.ID}
		{#if !filterTeachersResponse.teachers.some((teacher) => teacher.ID === resolvedTeacherPerson.ID)}
			<p class="hidden">
				{filterTeachersResponse.teachers.unshift(resolvedTeacherPerson)}
			</p>
		{/if}
	{/if}

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
