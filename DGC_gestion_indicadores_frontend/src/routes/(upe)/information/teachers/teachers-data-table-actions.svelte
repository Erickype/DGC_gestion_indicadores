<script lang="ts">
	import CircleX from 'lucide-svelte/icons/circle-x';
	import Pencil from 'lucide-svelte/icons/pencil';

	import { Button } from '$lib/components/ui/button';
	import Dialog from '$lib/components/alert/dialog.svelte';

	export let id: string;

    let dialogOpen = false
</script>

<Dialog bind:dialogOpen={dialogOpen}></Dialog>

<div class="flex flex-auto">
    <Button variant="ghost" size="icon" on:click={()=>{
		console.log("Update: "+id);
		
	}}>
        <Pencil class="h-4 w-4" />
    </Button>
    <Button variant="ghost" size="icon" on:click={()=>{
        console.log("Delete: "+id);      
        dialogOpen = !dialogOpen  
    }}>
        <CircleX class="stroke-primary h-4 w-4" />
    </Button>
</div>