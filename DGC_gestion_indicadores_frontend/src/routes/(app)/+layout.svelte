<script lang="ts">
	import MainMenu from '$lib/components/menu/mainMenu.svelte';
	import Header from '$lib/components/menu/header.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import { onMount } from 'svelte';
	import type { LayoutData } from '../$types';
	import { writable } from 'svelte/store';

	export let data: LayoutData;
	const userData = data.user;
	const timer = writable(0);

	onMount(() => {
		const remainingTime = parseInt((userData.eat - Date.now() / 1000).toFixed(0));

		const interval = setInterval(() => {
			timer.update((n) => n + 1);
			if ($timer >= remainingTime+1) {
				goto('/login');
			}
		}, 1000);

		return async () => {
				await invalidateAll()
			clearInterval(interval);
		}
	});
</script>

<div class="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
	<MainMenu></MainMenu>
	<div class="flex flex-col">
		<Header></Header>
		<main class="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
			<slot />
		</main>
	</div>
</div>
