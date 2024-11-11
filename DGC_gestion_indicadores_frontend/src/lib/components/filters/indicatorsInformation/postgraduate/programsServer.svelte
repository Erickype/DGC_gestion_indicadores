<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import { writable } from 'svelte/store';

	import { GenerateComboMessagesFromAuthors } from '$lib/api/controller/api/indicatorsInformation/postgraduate';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import type {
		PostgraduateProgram,
		FilterPostgraduateProgramsRequest,
		FilterPostGraduateProgramsResponse
	} from '$lib/api/model/api/indicatorsInformation/postgraduate';
	import {
		fetchGetGetPostgraduateProgramByID,
		fetchFilterPostgraduatePrograms,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterPostgraduateProgramsRequest,
		newPopoverFilterDataMap
	} from './programs';

	export let formDataPostgraduateProgramID = writable();
	let postgraduateProgramPromise: Promise<PostgraduateProgram>;

	let openPostgraduatePrograms = false;
	let filterPostgraduateProgramsRequest: FilterPostgraduateProgramsRequest =
		newFilterPostgraduateProgramsRequest(5, 1);
	filterPostgraduateProgramsRequest.is_active = true;
	let postgraduateProgramsFilter: string = '';

	postgraduateProgramPromise = fetchGetGetPostgraduateProgramByID(
		$formDataPostgraduateProgramID as string
	);

	let filterPostGraduateProgramsResponsePromise: Promise<FilterPostGraduateProgramsResponse> =
		fetchFilterPostgraduatePrograms(filterPostgraduateProgramsRequest);
	let postgraduateProgramsPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterPostGraduateProgramsResponsePromise = fetchOnFilterChanged(
			postgraduateProgramsFilter.trim(),
			filterPostgraduateProgramsRequest,
			postgraduateProgramsPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterPostGraduateProgramsResponsePromise = fetchOnDetailedFilter(
			filterPostgraduateProgramsRequest,
			postgraduateProgramsPopoverFilterDataMap
		).then(({ request, response }) => {
			filterPostgraduateProgramsRequest = request;
			filterPostgraduateProgramsRequest.is_active = true;
			postgraduateProgramsFilter = generateInitialFilterValue(filterPostgraduateProgramsRequest)!;
			return response;
		});
	}

	function setInitialValue(programID: number | undefined) {
		formDataPostgraduateProgramID.set(programID);
	}
</script>

{#await Promise.all([filterPostGraduateProgramsResponsePromise, postgraduateProgramPromise])}
	<FormFieldSkeleton />
{:then [filterPostGraduateProgramsResponse, postgraduateByIDResponse]}
	{#if postgraduateByIDResponse && postgraduateByIDResponse.ID}
		{#if !filterPostGraduateProgramsResponse.postgraduate_programs.some((postgraduate) => postgraduate.ID === postgraduateByIDResponse.ID)}
			<p class="hidden">
				{filterPostGraduateProgramsResponse.postgraduate_programs.unshift(postgraduateByIDResponse)}
			</p>
		{/if}
	{:else}
		<p class="hidden">
			{setInitialValue(filterPostGraduateProgramsResponse.postgraduate_programs.at(0)?.ID)}
		</p>
	{/if}

	<ServerFormSelect
		formSelectWidth="w-[50%]"
		bind:filterValue={postgraduateProgramsFilter}
		formLabel="Programa posgrado"
		bind:popoverFilterDataMap={postgraduateProgramsPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromAuthors(
			filterPostGraduateProgramsResponse.postgraduate_programs
		)}
		bind:openCombo={openPostgraduatePrograms}
		bind:formDataID={formDataPostgraduateProgramID}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
