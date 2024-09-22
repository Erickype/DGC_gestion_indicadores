<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import type {
		FilterDetailedFieldRequest,
		FilterDetailedFieldResponse,
		DetailedFilterJoined
	} from '$lib/api/model/api/knowledgeFields/detailedFields';
	import {
		fetchFilterDetailedFields,
		generateInitialFilterValue,
		newFilterDetailedFieldsRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/detailedFields/detailedFileds';
	import type { PopoverFilterDataMap } from '../../table/types';

	import { writable } from 'svelte/store';
	import { GenerateComboMessagesFromDetailedFilterJoined } from '$lib/api/controller/api/knowledgeFields/detailedFieldsFilter';

	export let formDataDetailedFieldID = writable();
	/* 	let teacherPerson: Promise<TeacherPerson>;
	 */
	let openTeachers = false;
	let filterDetailedFieldRequest: FilterDetailedFieldRequest = newFilterDetailedFieldsRequest(5, 1);
	let teachersFilterValue: string = '';

	/* 	teacherPerson = fetchTeacherPersonJoinedByTeacherID($formDataDetailedFieldID as string);
	 */
	let filterDetailedFieldResponsePromise: Promise<FilterDetailedFieldResponse> =
		fetchFilterDetailedFields(filterDetailedFieldRequest);
	let teachersPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	/* function handleOnTeachersFilterChanged() {
		filterDetailedFieldResponsePromise = fetchOnFilterChanged(
			teachersFilterValue.trim(),
			filterDetailedFieldRequest,
			teachersPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterDetailedFieldResponsePromise = fetchOnDetailedFilter(
			filterDetailedFieldRequest,
			teachersPopoverFilterDataMap
		).then(({ request, response }) => {
			filterDetailedFieldRequest = request;
			teachersFilterValue = generateInitialFilterValue(filterDetailedFieldRequest)!;
			return response;
		});
	} */
</script>

{#await Promise.all([filterDetailedFieldResponsePromise])}
	<FormFieldSkeleton />
{:then [filterDetailedFieldResponse]}
	<!-- {#if resolvedTeacherPerson.ID}
		{#if !filterDetailedFieldResponse.teachers.some((teacher) => teacher.ID === resolvedTeacherPerson.ID)}
			<p class="hidden">
				{filterDetailedFieldResponse.teachers.unshift(resolvedTeacherPerson)}
			</p>
		{/if}
	{/if} -->

	<ServerFormSelect
		bind:filterValue={teachersFilterValue}
		formLabel="Campo detallado"
		bind:popoverFilterDataMap={teachersPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromDetailedFilterJoined(
			filterDetailedFieldResponse.detailed_fields
		)}
		bind:openCombo={openTeachers}
		bind:formDataID={formDataDetailedFieldID}
	/>
{/await}
