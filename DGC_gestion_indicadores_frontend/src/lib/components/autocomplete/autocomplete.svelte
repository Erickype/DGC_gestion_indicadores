<script lang="ts">
	import type { Message } from './select';
	import Search from '$lib/icons/search.svelte';
	import { createEventDispatcher } from 'svelte';

	export let id: string;
	export let name: string;
	export let messages: Message[] = [];
	export let selected: number = 0;
	export let width: string = 'w-1/3';
	let messagesFilter: Message[] = [];

	const dispatch = createEventDispatcher();

	let input: HTMLInputElement;
	let dropdown: HTMLDetailsElement;

	let ul: HTMLUListElement;
	let items: NodeListOf<HTMLButtonElement>;
	let currentItem = 0;

	$: if (messages.length > 0) {
		messagesFilter = messages;
	}

	$: if (input && input.value === '') {
		reset();
	}

	function manageItemSelected(event: KeyboardEvent) {
		// Check if the current component's input is focused
		if (document.activeElement === input) {
			switch (event.key) {
				case 'ArrowUp':
					event.preventDefault();
					items[currentItem].classList.remove('focus');
					currentItem = currentItem > 0 ? --currentItem : 0;
					items[currentItem].classList.add('focus');
					break;
				case 'ArrowDown':
					event.preventDefault();
					items[currentItem].classList.remove('focus');
					currentItem = currentItem < items.length - 1 ? ++currentItem : items.length - 1;
					items[currentItem].classList.add('focus');
					break;
				case 'Enter':
					const button = items[currentItem] as HTMLButtonElement;
					selected = parseInt(button.value);
					dispatch('selected', selected);
					const text = button.innerText;
					input.value = text;
					dropdown.removeAttribute('open');
					break;
			}
		}
	}

	function itemSelected(event: Event) {
		const button = event.target as HTMLButtonElement;
		const value = parseInt(button.value);
		selected = value;
		dispatch('selected', selected);

		const text = button.innerText;
		input.value = text;

		dropdown.removeAttribute('open');
		keepItemsOpen = true;
	}

	function openSelect() {
		selected = 0;
		currentItem = 0;
		input.value = '';
		dropdown.setAttribute('open', 'open');

		setTimeout(() => {
			items = ul.querySelectorAll('.item');

			if (items.length > 0) {
				items[currentItem].classList.add('focus');
			}
		}, 0);
	}

	function filterMessages() {
		const searchText = input.value.toLowerCase();
		messagesFilter = messages.filter((message) => message.name.toLowerCase().includes(searchText));

		setTimeout(() => {
			items = ul.querySelectorAll('.item');

			if (items.length > 0) {
				currentItem = 0;
				items.forEach((item) => item.classList.remove('focus'));
				items[currentItem].classList.add('focus');
			}
			dropdown.setAttribute('open', 'true');
		}, 0);
	}

	function reset() {
		messagesFilter = messages;
	}

	let keepItemsOpen = false;
	function closeSelect() {
		if (!keepItemsOpen) {
			currentItem = 0;
			messagesFilter = messages;
			dropdown.removeAttribute('open');
		}
	}

	function closeItemsList() {
		if (keepItemsOpen) {
			currentItem = 0;
			messagesFilter = messages;
			dropdown.removeAttribute('open');
			keepItemsOpen = false;
		}
	}
</script>

<svelte:document on:keydown={manageItemSelected} on:mouseup={closeItemsList} />

<details bind:this={dropdown} class="dropdown {width}">
	<summary class="list-none">
		<label class="input input-bordered input-sm flex items-center gap-2">
			<input
				bind:this={input}
				{id}
				{name}
				type="search"
				class="grow"
				placeholder="Buscar"
				autocomplete="off"
				on:blur={closeSelect}
				on:reset={reset}
				on:focus={openSelect}
				on:input={filterMessages}
			/>
			<Search width="w-4" height="h4"></Search>
		</label>
	</summary>
	<ul bind:this={ul} class="p-2 shadow menu dropdown-content z-[1] bg-base-100 rounded-box w-full">
		{#each messagesFilter as message}
			<li value={message.id}>
				<button class="item" value={message.id} type="button" on:mousedown={itemSelected}
					>{message.name}</button
				>
			</li>
		{/each}
	</ul>
</details>
