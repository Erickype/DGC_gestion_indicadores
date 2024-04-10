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

		const dropdown = button.closest('.dropdown');
		dropdown!.removeAttribute('open');
	}

	function openSelect(event: Event) {
		const input = event.target as HTMLInputElement;
		const dropdown = input.closest('.dropdown');
		dropdown!.setAttribute('open', 'open');
	}
</script>

<details class="dropdown w-1/3">
	<summary class="list-none">
		<label class="input input-bordered flex items-center gap-2">
			<input type="text" class="grow" placeholder="Search" on:focus={openSelect} />
			<Search customClass="w-4 h-4"></Search>
		</label>
	</summary>
	<ul class="p-2 shadow menu dropdown-content z-[1] bg-base-100 rounded-box w-full">
		{#each messages as message}
			<li value={message.id}>
				<button value={message.id} type="button" on:click={itemSelected}>{message.name}</button>
			</li>
		{/each}
	</ul>
</details>
