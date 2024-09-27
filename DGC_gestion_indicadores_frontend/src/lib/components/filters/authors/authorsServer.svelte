<script lang="ts">
	import ServerFormSelect from '$lib/components/combobox/serverFormSelect.svelte';
	import FormFieldSkeleton from '$lib/components/skeleton/formField.svelte';

	import { writable } from 'svelte/store';

	import type {
		AuthorPerson,
		FilterAuthorsRequest,
		FilterAuthorsResponse
	} from '$lib/api/model/api/academicProduction/authors/authorsFilter';
	import {
		fetchAuthorPersonJoinedByAuthorID,
		fetchFilterAuthors,
		fetchOnDetailedFilter,
		fetchOnFilterChanged,
		generateInitialFilterValue,
		newFilterAuthorsRequest,
		newPopoverFilterDataMap
	} from './authors';
	import type { PopoverFilterDataMap } from '$lib/components/table/types';
	import { GenerateComboMessagesFromAuthors } from '$lib/api/controller/api/academicProduction/authors/author';

	export let formDataAuthorID = writable();
	let authorPerson: Promise<AuthorPerson>;

	let openAuthors = false;
	let filterAuthorsRequest: FilterAuthorsRequest = newFilterAuthorsRequest(5, 1);
	let authorsFilterValue: string = '';

	authorPerson = fetchAuthorPersonJoinedByAuthorID($formDataAuthorID as string);

	let filterAuthorsResponsePromise: Promise<FilterAuthorsResponse> =
		fetchFilterAuthors(filterAuthorsRequest);
	let teachersPopoverFilterDataMap: PopoverFilterDataMap = newPopoverFilterDataMap();

	function handleOnTeachersFilterChanged() {
		filterAuthorsResponsePromise = fetchOnFilterChanged(
			authorsFilterValue.trim(),
			filterAuthorsRequest,
			teachersPopoverFilterDataMap
		);
	}

	async function handleOnDetailedFilter() {
		filterAuthorsResponsePromise = fetchOnDetailedFilter(
			filterAuthorsRequest,
			teachersPopoverFilterDataMap
		).then(({ request, response }) => {
			filterAuthorsRequest = request;
			authorsFilterValue = generateInitialFilterValue(filterAuthorsRequest)!;
			return response;
		});
	}
</script>

{#await Promise.all([filterAuthorsResponsePromise, authorPerson])}
	<FormFieldSkeleton />
{:then [filterTeachersResponse, resolvedTeacherPerson]}
	{#if resolvedTeacherPerson.ID}
		{#if !filterTeachersResponse.authors.some((teacher) => teacher.ID === resolvedTeacherPerson.ID)}
			<p class="hidden">
				{filterTeachersResponse.authors.unshift(resolvedTeacherPerson)}
			</p>
		{/if}
	{/if}

	<ServerFormSelect
		bind:filterValue={authorsFilterValue}
		formLabel="Autor"
		bind:popoverFilterDataMap={teachersPopoverFilterDataMap}
		comboData={GenerateComboMessagesFromAuthors(filterTeachersResponse.authors)}
		bind:openCombo={openAuthors}
		bind:formDataID={formDataAuthorID}
		on:filterChanged={handleOnTeachersFilterChanged}
		on:detailedFilter={handleOnDetailedFilter}
	/>
{/await}
