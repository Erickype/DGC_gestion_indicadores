<script lang="ts">
	import AcademicPeriodAutocomplete from '$lib/components/autocomplete/academicPeriodAutocomplete.svelte';

	let selectedPeriod: number = 0

	$:if(selectedPeriod > 0){
		// todo: load data based on this selected period
	}
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div class="flex items-center justify-center">
	<h1 class="text-xl font-bold">Docentes</h1>
	<AcademicPeriodAutocomplete id="period" name="period" bind:selected={selectedPeriod}></AcademicPeriodAutocomplete>
	{selectedPeriod}
</div>
