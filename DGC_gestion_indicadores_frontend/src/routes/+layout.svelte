<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/stores';
	import '../app.css';
</script>

<div class="navbar bg-primary">
	<div class="flex-1">
		{#if $page.data.user}
			<a href="/admin" class="btn btn-ghost text-lg">Admin</a>
		{/if}
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

<main>
	<slot />
</main>
