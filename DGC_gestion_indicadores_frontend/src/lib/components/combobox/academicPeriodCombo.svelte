<script lang="ts">
	import Combobox from './combobox.svelte';
	import type { Message } from './combobox';

	export let messages: Message[];
	export let selectedValue: number;
</script>

<Combobox
	data={messages}
	placeholder="Seleccionar periodo académico..."
	bind:selectedValue
	emptyMessage="No se encontró un periodo"
	on:message
></Combobox>
