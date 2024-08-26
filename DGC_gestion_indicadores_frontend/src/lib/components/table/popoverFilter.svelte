<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import Filter from 'lucide-svelte/icons/filter';
	import type { PopoverFilterDataMap } from './types';
	import { createEventDispatcher } from 'svelte';

	export let popoverFilterDataMap: PopoverFilterDataMap;

	const dispatch = createEventDispatcher();

	function handleOnClickDetailedFilter() {
		dispatch('detailedFilter');
	}

	function handleKeyDown(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			event.preventDefault();
			handleOnClickDetailedFilter();
		}
	}
</script>

<Popover.Root portal={null}>
	<Popover.Trigger asChild let:builder>
		<Button builders={[builder]} variant="outline">
			<Filter class="h-4 w-4" />
		</Button>
	</Popover.Trigger>
	<Popover.Content class="w-80">
		<div class="grid gap-4">
			<div class="space-y-2">
				<h4 class="font-medium leading-none">Filtro detallado</h4>
				<p class="text-muted-foreground text-sm">Ingrese los criterios a buscar</p>
			</div>
			<div class="flex flex-col gap-1">
				<div class="grid gap-2">
					{#each popoverFilterDataMap.entries() as [key, popoverFilterData]}
						<div class="grid grid-cols-3 items-center gap-4">
							<Label for={key}>{popoverFilterData.label}</Label>
							<Input
								id={key}
								bind:value={popoverFilterData.value}
								class="col-span-2 h-8"
								on:keydown={handleKeyDown}
							/>
						</div>
					{/each}
				</div>
				<Button type="button" variant="secondary" size="sm" on:click={handleOnClickDetailedFilter}>
					Filtrar
					<Filter class="ml-2 h-4 w-4" />
				</Button>
			</div>
		</div>
	</Popover.Content>
</Popover.Root>
