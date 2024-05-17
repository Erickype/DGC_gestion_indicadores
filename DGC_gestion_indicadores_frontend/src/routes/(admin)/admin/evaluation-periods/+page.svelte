<script lang="ts">
	import type { PageData } from './$types';
	import AddModal from './AddModal.svelte';

	export let data: PageData

	const formData = data.addEvaluationPeriodForm
</script>

<svelte:head>
	<title>Periodos Evaluación</title>
</svelte:head>

<div class="m-auto flex w-full p-8">
	<AddModal {formData}/>
</div>