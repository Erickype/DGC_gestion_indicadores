<script lang="ts">
	import type { Message } from './select';
	import Search from '$lib/icons/search.svelte';
	import { createEventDispatcher } from 'svelte';

	export let messages: Message[] = [];
	export let selected: number = 0;
	export let width: string = 'w-1/3';
	let messagesFilter: Message[] = [];

	const dispatch = createEventDispatcher();

	$: if (messages.length > 0) {
		messagesFilter = messages;
	}

	function itemSelected(event: Event) {
		const button = event.target as HTMLButtonElement;
		const value = parseInt(button.value);
		selected = value;
		dispatch('selected', selected);

		const input = document.getElementById('input') as HTMLInputElement;
		const text = button.innerText;
		input.value = text;

		const dropdown = button.closest('.dropdown');
		dropdown!.removeAttribute('open');
	}

	function openSelect(event: Event) {
		const input = event.target as HTMLInputElement;
		selected = 0;
		input.value = '';
		const dropdown = input.closest('.dropdown');
		dropdown!.setAttribute('open', 'open');
	}

	function filterMessages(event: Event) {
		const input = event.target as HTMLInputElement;

		const searchText = input.value.toLowerCase();
		messagesFilter = messages.filter((message) => message.name.toLowerCase().includes(searchText));

		const dropdown = input.closest('.dropdown');
		dropdown!.setAttribute('open', 'open');
	}

	function reset(event: Event) {
		const input = event.target as HTMLInputElement;
		if ((input.value = '')) {
			messagesFilter = messages;
		}
	}
</script>

<details class="dropdown {width}">
	<summary class="list-none">
		<label class="input input-bordered input-sm flex items-center gap-2">
			<input
				id="input"
				type="search"
				class="grow"
				placeholder="Buscar"
				on:reset={reset}
				on:focus={openSelect}
				on:input={filterMessages}
			/>
			<Search customClass="w-4 h-4"></Search>
		</label>
	</summary>
	<ul id="items" class="p-2 shadow menu dropdown-content z-[1] bg-base-100 rounded-box w-full">
		{#each messagesFilter as message}
			<li tabindex="-1" value={message.id}>
				<button value={message.id} type="button" on:click={itemSelected}>{message.name}</button>
			</li>
		{/each}
	</ul>
</details>
