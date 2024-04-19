<script lang="ts">
	import Combobox from "./combobox.svelte";
	import type { Message } from "./combobox";

    export let messages: Message[]
    export let selectedValue: number
    export let pixelsWidth: string = "300"
    
</script>

<Combobox
	data={messages}
	placeholder="Seleccionar periodo académico..."
	bind:selectedValue
	bind:pixelsWidth
	emptyMessage="No se encontró un periodo"
	on:message
></Combobox>
