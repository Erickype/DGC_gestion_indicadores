<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog';
	import Plus from 'lucide-svelte/icons/plus';

	import { createEventDispatcher } from 'svelte';

	export let formData;
	export let formComponent;
	export let modalTitle = 'Crear';
	export let modalDescription = 'Resgistrar la información luego click en Guardar.';

	let open: boolean;
	const dispatch = createEventDispatcher();

	function handleCreated(event: any) {
		const data: { created: boolean } = event.detail;
		if (data.created) {
			open = false;
		}
		dispatch('created', {
			status: true
		});
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger>
		<Button variant="default" class="ml-auto">
			Nuevo <Plus class="ml-2 h-4 w-4" />
		</Button>
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>{modalTitle}</Dialog.Title>
			<Dialog.Description>{modalDescription}</Dialog.Description>
		</Dialog.Header>
		<svelte:component this={formComponent} data={formData} on:message={handleCreated}
		></svelte:component>
	</Dialog.Content>
</Dialog.Root>
