<script>
	import { ModeWatcher } from 'mode-watcher';
	import { page } from '$app/stores';
	import '../app.pcss';
	import MainMenu from '$lib/components/menu/mainMenu.svelte';
	import Header from '$lib/components/menu/header.svelte';
	import { Toaster } from "$lib/components/ui/sonner";
</script>

<ModeWatcher />
<Toaster />

{#if $page.data.user}
	<div class="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
		<MainMenu></MainMenu>
		<div class="flex flex-col">
			<Header></Header>
			<main class="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
				<slot />
			</main>
		</div>
	</div>
{/if}

{#if !$page.data.user}
	<slot />
{/if}
