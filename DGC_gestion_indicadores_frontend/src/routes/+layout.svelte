<script lang="ts">
	import MenuIcon from '$lib/icons/menu.svelte';
	import MainMenu from '$lib/components/menu/mainMenu.svelte';
	import { applyAction, enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
	import '../app.css';

	function closeDrawer() {
		const drawerCheckbox = document.getElementById('my-drawer') as HTMLInputElement;
		if (drawerCheckbox) {
			drawerCheckbox.checked = false;
		}
	}
</script>

<div class="drawer">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content h-screen flex flex-col overflow-hidden">
		<div class="h-min">
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
						<form
							action="/logout"
							method="POST"
							use:enhance={() => {
								return async ({ result }) => {
									invalidateAll();
									await applyAction(result);
								};
							}}
						>
							<button class="btn btn-ghost" type="submit">Salir</button>
						</form>
					{/if}
				</div>
			</div>
		</div>

		<main style="height: 90%;">
			<slot />
		</main>
	</div>
	<div class="drawer-side">
		<label for="my-drawer" aria-label="close sidebar" class="drawer-overlay"></label>
		<ul class="h-full menu bg-base-200 w-1/4">
			<MainMenu on:clicked={closeDrawer}></MainMenu>
		</ul>
	</div>
</div>
