<script lang="ts">
    import MoveLeft from "lucide-svelte/icons/move-left";

    import { Button } from "$lib/components/ui/button/index";
	import { browser } from "$app/environment";

    function returnToTeachers() {
        if(browser){
            window.history.back()
        }
    }

</script>

<Button variant="outline" on:click={returnToTeachers}>
    <MoveLeft class="h-4 w-4" />
</Button>