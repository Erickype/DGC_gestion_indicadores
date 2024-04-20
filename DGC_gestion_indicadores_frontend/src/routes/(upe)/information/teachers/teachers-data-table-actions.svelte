<script lang="ts">
	import CircleX from 'lucide-svelte/icons/circle-x';
	import Pencil from 'lucide-svelte/icons/pencil';

	import { Button } from '$lib/components/ui/button';
	import Dialog from '$lib/components/alert/dialog.svelte';
	import { hasTeacherDeleted, updateTeacherAction } from '../../../../stores';
	import { error } from '@sveltejs/kit';

	export let id: string;

    let dialogOpen = false

    async function deleteTeacher(event: CustomEvent) {                
        if(event.detail.status == true){
            console.log("Deleting...");
            const res = await fetchDeleteTeacher(id)
            if(res.status == "success"){
                hasTeacherDeleted.set(true)
            }
        }
    }

    async function fetchDeleteTeacher(teacherID: string) {
		const url = `/api/teacher/${teacherID}`;
		const response = await fetch(url, {
			method: 'DELETE',
			credentials: 'include'
		});
		if (!response.ok) {
            throw error(500, "Failed to delete teachers");            
		}

        return await response.json()
	}

    function sendUpdateTeacherAction() {
        updateTeacherAction.set({
            status: true,
            teacherID: parseInt(id)
        })
    }
</script>

<Dialog bind:dialogOpen={dialogOpen} on:dialog-continue={deleteTeacher}></Dialog>

<div class="flex flex-auto">
    <Button variant="ghost" size="icon" on:click={()=>{
		sendUpdateTeacherAction()
	}}>
        <Pencil class="h-4 w-4" />
    </Button>
    <Button variant="ghost" size="icon" on:click={()=>{
        dialogOpen = !dialogOpen
    }}>
        <CircleX class="stroke-primary h-4 w-4" />
    </Button>
</div>