<script lang="ts">
	import Combobox from "./combobox.svelte";
	import type { Message } from "./combobox";

    export let messages: Message[]
    export let selectedValue: string
    export let pixelsWidth: string = "300"
    
</script>

<Combobox
	data={messages}
	placeholder="Seleccionar persona existente..."
	bind:selectedValue
	bind:pixelsWidth
	emptyMessage="No se encontró a la persona un periodo"
></Combobox>
