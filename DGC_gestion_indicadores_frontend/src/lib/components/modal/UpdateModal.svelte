<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';

	import { createEventDispatcher } from 'svelte';
	
	import type { Message } from '../combobox/combobox';

	export let formData;
	export let formComponent;
	export let comboMessages: [Message[]] | undefined = undefined;
	export let updateEntity;
	export let modalTitle = 'Actualizar';
	export let modalDescription = 'Actualizar la información luego click en Guardar.';

	export let open: boolean;
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
			<Dialog.Title>{modalTitle}</Dialog.Title>
			<Dialog.Description>{modalDescription}</Dialog.Description>
		</Dialog.Header>
		<svelte:component
			this={formComponent}
			data={formData}
			{updateEntity}
			{comboMessages}
			on:updated={handleUpdated}
		></svelte:component>
	</Dialog.Content>
</Dialog.Root>
