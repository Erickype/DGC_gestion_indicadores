<script lang="ts">
	import ChevronRight from 'lucide-svelte/icons/chevron-right';
	import CircleX from 'lucide-svelte/icons/circle-x';
	import Pencil from 'lucide-svelte/icons/pencil';

	import Dialog from '$lib/components/alert/dialog.svelte';
	import { Button } from '$lib/components/ui/button';
	import { createEventDispatcher } from 'svelte';

	export let id: string;
	export let moreAction = false;
	export let moreActionURL: string | null = null;

	let dialogOpen = false;

	const dispatch = createEventDispatcher();

	function sendDeleteConfirmation() {
		dispatch('delete-confirmation', {
			status: true,
			id: id
		});
	}

	function sendUpdateAction() {
		dispatch('update-action', {
			status: true,
			id: id
		});
	}
</script>

<Dialog bind:dialogOpen on:dialog-continue={sendDeleteConfirmation}></Dialog>

<div class="flex flex-auto">
	<Button variant="ghost" size="icon" on:click={sendUpdateAction}>
		<Pencil class="h-4 w-4" />
	</Button>
	{#if moreAction && moreActionURL}
		<Button variant="ghost" size="icon" href={moreActionURL + id}>
			<ChevronRight class="h-4 w-4" />
		</Button>
	{/if}
	<Button
		variant="ghost"
		size="icon"
		on:click={() => {
			dialogOpen = !dialogOpen;
		}}
	>
		<CircleX class="stroke-primary h-4 w-4" />
	</Button>
</div>
