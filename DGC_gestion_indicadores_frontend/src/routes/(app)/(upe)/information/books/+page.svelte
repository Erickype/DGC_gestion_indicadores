<script lang="ts">
	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import AddForm from './AddForm.svelte';

	import Icon from 'lucide-svelte/icons/book-open-text';

	import type { PageServerData } from './$types';

	import AddModal from '$lib/components/modal/AddModal.svelte';

	export let data: PageServerData;
	const addBookOrChaptersProductionForm = data.addBookOrChaptersProductionForm;
	const updateBookOrChaptersProductionForm = data.updateBookOrChaptersProductionForm;

	const academicPeriodsData = data.academicPeriodsData;

	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(0)!.ID;

	$: {
		addBookOrChaptersProductionForm.data.academic_period_id = selectedAcademicPeriod;
	}

	function fetchOnAcademicPeriodChange() {}

	function fetchOnSuccess(event: CustomEvent) {}
</script>

<svelte:head>
	<title>Libros y Capítulos</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<div class="flex items-center gap-1">
		<Icon class="h-8 w-8" />
		<h2 class="text-2xl font-bold">Libros y Capítulos</h2>
	</div>
</div>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={fetchOnAcademicPeriodChange}
	></AcademicPeriodCombo>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear libro o capítulo"
		formData={addBookOrChaptersProductionForm}
		on:created={fetchOnSuccess}
	/>
</div>

<div class="mx-auto flex w-full flex-col place-content-center px-8">
</div>