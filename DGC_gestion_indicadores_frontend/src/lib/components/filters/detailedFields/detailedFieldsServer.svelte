<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import type {
		FilterDetailedFieldRequest,
		FilterDetailedFieldResponse,
		DetailedFieldJoined
	} from '$lib/api/model/api/knowledgeFields/detailedFields';
	import {
		fetchFilterDetailedFields,
		fetchGetDetailedFieldJoinedByDetailedFieldID,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterDetailedFieldsRequest,
		newPopoverFilterDataMap
	} from '$lib/components/filters/detailedFields/detailedFileds';
	import type { PopoverFilterDataMap } from '../../table/types';

	import { writable } from 'svelte/store';
	import { GenerateComboMessagesFromDetailedFieldJoined } from '$lib/api/controller/api/knowledgeFields/detailedFieldsFilter';

	export let formDataDetailedFieldID = writable();
	let detailedFieldJoinedPromise: Promise<DetailedFieldJoined>;

	let openTeachers = false;
	let filterDetailedFieldRequest: FilterDetailedFieldRequest = newFilterDetailedFieldsRequest(5, 1);
	let teachersFilterValue: string = '';

	detailedFieldJoinedPromise = fetchGetDetailedFieldJoinedByDetailedFieldID(
		$formDataDetailedFieldID as string
	);

	let filterDetailedFieldResponsePromise: Promise<FilterDetailedFieldResponse> =
		fetchFilterDetailedFields(filterDetailedFieldRequest);
	let teachersPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnFilterChanged() {
		filterDetailedFieldResponsePromise = fetchOnFilterChanged(
			teachersFilterValue.trim(),
			filterDetailedFieldRequest,
			teachersPopoverFilterDataMap
		);
	}

	/* async function handleOnDetailedFilter() {
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

{#await Promise.all([filterDetailedFieldResponsePromise, detailedFieldJoinedPromise])}
	<FormFieldSkeleton />
{:then [filterDetailedFieldResponse, detailedFieldJoinedResponse]}
	{#if detailedFieldJoinedResponse.detailed_field_id}
		{#if !filterDetailedFieldResponse.detailed_fields.some((detailedField) => detailedField.detailed_field_id === detailedFieldJoinedResponse.detailed_field_id)}
			<p class="hidden">
				{filterDetailedFieldResponse.detailed_fields.unshift(detailedFieldJoinedResponse)}
			</p>
		{/if}
	{/if}

	<ServerFormSelect
		bind:filterValue={teachersFilterValue}
		formLabel="Campo detallado"
		bind:popoverFilterDataMap={teachersPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromDetailedFieldJoined(
			filterDetailedFieldResponse.detailed_fields
		)}
		bind:openCombo={openTeachers}
		bind:formDataID={formDataDetailedFieldID}
		on:filterChanged={handleOnFilterChanged}
	/>
{/await}
