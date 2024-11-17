<script lang="ts">
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Form from '$lib/components/ui/form';
	import { cn } from '$lib/utils.js';

	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import { createEventDispatcher, tick } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import Check from 'lucide-svelte/icons/check';
	import type { Message } from './combobox';
	import { writable } from 'svelte/store';

	import PopoverFilter from '$lib/components/table/popoverFilter.svelte';

	import type { PopoverFilterDataMap } from '../table/types';
	import Search from 'lucide-svelte/icons/search';
	import X from 'lucide-svelte/icons/x';

	export let comboData: Message[];
	export let openCombo = false;
	export let formDataID = writable();
	export let emptyLabel: string = 'Seleccionar';
	export let popoverFilterDataMap: PopoverFilterDataMap = new Map();
	export let filterValue: string;
	export let formLabel: string = 'Object';
	export let formSelectWidth = 'w-[90%]';
	export let hasLabel = true

	let isFocused = false;

	function handleFocus() {
		isFocused = true;
	}

	function handleBlur() {
		isFocused = false;
	}

	const dispatch = createEventDispatcher();

	function handleFilterChanged() {
		return dispatch('filterChanged');
	}
	function handleOnDetailedFilter() {
		return dispatch('detailedFilter');
	}

	let typingTimeout: number;
	function handleInput() {
		clearTimeout(typingTimeout);
		typingTimeout = window.setTimeout(() => {
			handleFilterChanged();
		}, 500);
	}

	function handleDeleteFilter() {
		if (filterValue !== '') {
			filterValue = '';
			return handleFilterChanged();
		}
	}

	function closeAndFocusTrigger(triggerId: string) {
		openCombo = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<Popover.Root bind:open={openCombo} let:ids>
	<Form.Control let:attrs>
		{#if hasLabel}
			<Form.Label>{formLabel}</Form.Label>
		{/if}
		<Popover.Trigger
			class={cn(
				buttonVariants({ variant: 'outline' }),
				'w-full justify-between',
				!$formDataID && 'text-muted-foreground'
			)}
			role="combobox"
			{...attrs}
		>
			<ScrollArea class="w-auto text-pretty py-4 text-left" orientation="horizontal">
				{comboData.find((f) => f.value === $formDataID)?.label ?? emptyLabel}
			</ScrollArea>
			<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Popover.Trigger>
		<input hidden value={$formDataID} name={attrs.name} />
	</Form.Control>
	<Popover.Content class="{formSelectWidth} p-0">
		<Command.Root loop>
			<div class="flex items-center gap-1 p-1">
				<div
					class="flex w-full items-center rounded-sm border px-3 pr-2 {isFocused
						? 'border-bg ring-ring ring-offset-2-primary ring-offset-background ring-2 ring-offset-2'
						: 'border-bg'}"
				>
					<Search class={'mr-2 h-4 w-4'} />
					<input
						class="placeholder:text-muted-foreground flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-none disabled:cursor-not-allowed disabled:opacity-50"
						placeholder="Filtrar..."
						type="text"
						bind:value={filterValue}
						on:input={handleInput}
						on:focus={handleFocus}
						on:blur={handleBlur}
					/>
					<Button class="h-min p-1" variant="ghost" on:click={handleDeleteFilter}>
						<X class={'h-4 w-4'} />
					</Button>
				</div>
				<PopoverFilter bind:popoverFilterDataMap on:detailedFilter={handleOnDetailedFilter}
				></PopoverFilter>
			</div>
			<Command.Empty>Sin coincidencias!</Command.Empty>
			<Command.Group>
				{#each comboData as comboMessage}
					<Command.Item
						value={comboMessage.label}
						onSelect={() => {
							$formDataID = comboMessage.value;
							closeAndFocusTrigger(ids.trigger);
						}}
					>
						{comboMessage.label}
						<Check
							class={cn(
								'ml-auto h-4 w-4',
								comboMessage.value !== $formDataID && 'text-transparent'
							)}
						/>
					</Command.Item>
				{/each}
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
