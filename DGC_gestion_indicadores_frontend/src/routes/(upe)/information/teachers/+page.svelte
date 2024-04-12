<script lang="ts">
	import AcademicPeriodAutocomplete from '$lib/components/autocomplete/academicPeriodAutocomplete.svelte';

	let selectedPeriod: number = 0;

	$: if (selectedPeriod > 0) {
		// todo: load data based on this selected period
	}
</script>

<svelte:head>
	<title>Docentes</title>
</svelte:head>

<div id="headerSection" class="px-4 py-2 items-center">
	<div class="flex-row h-full justify-around items-center card bg-base-300 rounded-box">
		<h1 class="text-xl font-bold">Docentes</h1>
		<AcademicPeriodAutocomplete id="period" name="period" bind:selected={selectedPeriod} />
	</div>
</div>
<div id="formSection" class="overflow-auto flex-grow px-4 py-2 items-center">
	<div class="flex h-full justify-center items-center bg-base-300 rounded-box"></div>
</div>
<div id="contentSection" class="flex-grow px-4 py-2 items-center">
	<div class="overflow-x-auto h-full bg-base-300 px-6 py-2 rounded-box"></div>
</div>

<style>
	#headerSection {
		height: 10%;
	}
	#formSection {
		height: 35%;
	}
	#contentSection {
		height: 55%;
	}
</style>
