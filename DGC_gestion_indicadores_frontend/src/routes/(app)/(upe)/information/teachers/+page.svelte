<script lang="ts">
	import type { PageServerData } from './$types';
	import { goto } from '$app/navigation';

	import { toast } from 'svelte-sonner';

	import AcademicPeriodCombo from '$lib/components/combobox/academicPeriodCombo.svelte';
	import AddModal from '$lib/components/modal/AddModal.svelte';
	import AddForm from './AddForm.svelte';

	import type { Message } from '$lib/components/combobox/combobox';

	export let data: PageServerData;
	const addTeacherForm = data.addTeacherForm;

	const academicPeriodsData = data.academicPeriodsData;

	const comboMessages: Message[][] = [
		data.peopleData.messages,
		data.careersData.messages,
		data.dedicationsData.messages,
		data.scaledGradesData.messages
	];

	let teachersPromise = data.teachersByAcademicPeriod;

	let selectedAcademicPeriod: number = academicPeriodsData.periods.at(
		academicPeriodsData.periods.length - 1
	)!.ID;

	$: addTeacherForm.data.academicPeriod = selectedAcademicPeriod;

	async function fetchTeachers() {
		const url = `/api/teacher/byAcademicPeriodID/${selectedAcademicPeriod}`;
		const response = await fetch(url, {
			method: 'GET'
		});
		if (response.ok) {
			return (teachersPromise = response.json());
		}

		if (response.status === 401) {
			toast.warning('No está autenticado.');
			return goto('/login');
		}
	}

	function fetchOnSuccess(event: CustomEvent) {
		const detail: { status: boolean } = event.detail;
		if (detail.status) {
			fetchTeachers();
		}
	}
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div class="mx-auto flex w-full place-content-center justify-between px-8">
	<h2 class="text-2xl font-bold">Docentes</h2>
</div>

<div class="flex items-center justify-between px-8">
	<AcademicPeriodCombo
		messages={academicPeriodsData.messages}
		bind:selectedValue={selectedAcademicPeriod}
		on:message={fetchTeachers}
	></AcademicPeriodCombo>

	<AddModal
		formComponent={AddForm}
		modalTitle="Crear docente"
		formData={addTeacherForm}
		{comboMessages}
		on:created={fetchOnSuccess}
	/>
</div>
