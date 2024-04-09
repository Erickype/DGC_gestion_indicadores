<script lang="ts">
	import type { Message } from './select';
	import Search from '$lib/icons/search.svelte';
	import { createEventDispatcher } from 'svelte';

	export let messages: Message[] = [];
	export let selected: number = 0;

	const dispatch = createEventDispatcher();

	function itemSelected(event: Event) {
		const button = event.target as HTMLButtonElement;
        const value = parseInt(button.value);
        selected = value;
        dispatch('selected', selected);		
	}
	$: console.log(selected);
	
</script>

<div class="flex flex-col w-1/2">
	<div class="dropdown w-full">
		<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
		<label tabindex="0" class="input input-bordered flex items-center gap-2">
			<input type="text" class="grow" placeholder="Search" />
			<Search customClass="w-4 h-4"></Search>
		</label>
		<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
		<ul tabindex="0" class="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-full">
			{#each messages as message}
				<!-- svelte-ignore a11y-missing-attribute -->
				<li value={message.id}>
					<button value={message.id} type="button" on:click={itemSelected}>{message.name}</button>
				</li>
			{/each}
		</ul>
	</div>
</div>