<script>
	import AcademicPeriodSelect from "$lib/components/autocomplete/academicPeriodAutocomplete.svelte";
</script>

<svelte:head>
	<title>Dashboard</title>
</svelte:head>

<div class="flex items-center justify-center mt-10">
	<h1 class="text-xl">Dashboard</h1>
	<AcademicPeriodSelect id="s1" name="s1" selected={0}></AcademicPeriodSelect>
	<AcademicPeriodSelect id="s2" name="s2" selected={0}></AcademicPeriodSelect>
</div>