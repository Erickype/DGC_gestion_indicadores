<script lang="ts">
	import Avatar from '$lib/icons/avatar.svelte';
	import Info from '$lib/icons/info.svelte';
	import Key from '$lib/icons/key.svelte';
	import Login from '$lib/icons/login.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<div class="card card-bordered bg-accent/10 border-secondary w-96 text-secondary-content">
	<div class="card-body items-center text-center">
		<Login customClass="fill-primary stroke-primary w-12 h-12"></Login>
		<h1 class="card-title text-4xl font-bold mb-4">Bienvenido!</h1>

		<form action="?/login" method="POST">
			<label class="input input-bordered flex items-center gap-2 mb-4" for="username">
				<Avatar customClass="w-4 h-4 opacity-70"></Avatar>
				<input
					class="glow"
					placeholder="Usuario"
					id="username"
					name="username"
					type="text"
					required
				/>
			</label>

			<label class="input input-bordered flex items-center gap-2 mb-4" for="password">
				<Key customClass="w-4 h-4 opacity-70"></Key>
				<input
					class="glow"
					placeholder="Contraseña"
					id="password"
					name="password"
					type="password"
					required
				/>
			</label>

			{#if form?.error}
				<div role="alert" class="alert alert-warning mb-4">
					<Info customClass="stroke-error w-6 h-6"></Info>
					<span class="text-md">Credenciales incorrectas!</span>
				</div>
			{/if}

			<button class="btn btn-primary" type="submit">Ingresar</button>
		</form>
	</div>
</div>
