<script lang="ts">
	import type { AcademicPeriod } from '$lib/api/model/view/academicPeriod';
	import * as Dialog from '$lib/components/ui/dialog';

	import UpdateForm from './UpdateForm.svelte';
	import { createEventDispatcher } from 'svelte';

	export let formData;

	export let open: boolean;
	export let academicPeriod: AcademicPeriod;
	const dispatch = createEventDispatcher();

	function handleUpdated(event: any) {
		const data: { updated: boolean } = event.detail;
		if (data.updated) {
			open = false;
		}
		dispatch('updated', {
			status: true
		});
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Actualizar periodo académico</Dialog.Title>
			<Dialog.Description>Actualizar la información luego click en Guardar.</Dialog.Description>
		</Dialog.Header>
		<UpdateForm data={formData} {academicPeriod} on:updated={handleUpdated} />
	</Dialog.Content>
</Dialog.Root>
