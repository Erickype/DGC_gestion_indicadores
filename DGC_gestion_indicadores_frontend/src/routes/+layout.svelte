<script lang="ts">
	import { onMount } from 'svelte';
	import MenuIcon from '$lib/icons/menu.svelte';
	import MainMenu from '$lib/components/menu/mainMenu.svelte';
	import { enhance } from '$app/forms';
	import { page } from '$app/stores';
	import '../app.css';

	// Function to close the drawer
	function closeDrawer() {
		const drawerCheckbox = document.getElementById('my-drawer') as HTMLInputElement;
		if (drawerCheckbox) {
			drawerCheckbox.checked = false;
		}
	}

	onMount(() => {
		// Add click event listeners to all anchor elements in the sidebar
		const sidebarLinks = document.querySelectorAll('.menu a');
		sidebarLinks.forEach((link) => {
			link.addEventListener('click', closeDrawer);
		});

		// Clean up event listeners on component unmount
		return () => {
			sidebarLinks.forEach((link) => {
				link.removeEventListener('click', closeDrawer);
			});
		};
	});
</script>

<div class="navbar bg-primary">
	<div class="flex-1 pl-4">
		{#if $page.data.user}
			<label for="my-drawer" class="btn btn-ghost drawer-button mr-4">
				<MenuIcon></MenuIcon>
			</label>
		{/if}
		<h1 class="text-xl font-semibold">Gestión de Indicadores</h1>
	</div>
	<div class="flex-none">
		{#if !$page.data.user}
			<div class="gap-2">
				<a href="/login" class="btn btn-ghost text-lg">Ingresar</a>
				<a href="/register" class="btn btn-ghost text-lg">Registrarse</a>
			</div>
		{/if}
		{#if $page.data.user}
			<form action="/logout" method="post" use:enhance>
				<button class="btn btn-ghost" type="submit">Salir</button>
			</form>
		{/if}
	</div>
</div>

<div class="drawer">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content">
		<!-- Page content here -->
		<main>
			<slot />
		</main>
	</div>
	<div class="drawer-side">
		<label for="my-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
		<ul class="menu bg-base-200 w-1/4">
			<MainMenu></MainMenu>
		</ul>
	</div>
</div>
