<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import Avatar from '$lib/icons/avatar.svelte';
	import Info from '$lib/icons/info.svelte';
	import Key from '$lib/icons/key.svelte';
	import Login from '$lib/icons/login.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<svelte:head>
	<title>Ingresar</title>
</svelte:head>

<div class="flex items-center justify-center mt-10">
	<div class="card card-bordered bg-accent/10 border-secondary w-96 text-secondary-content">
		<div class="card-body items-center text-center">
			<Login fill="fill-primary" stroke="stroke-neutral-content/10" width="w-14" height="h-14"
			></Login>
			<h1 class="card-title text-4xl font-bold mb-4">Bienvenido!</h1>

			<form
				action="?/login"
				method="POST"
				use:enhance={() => {
					return async ({ result }) => {
						invalidateAll();
						await applyAction(result);
					};
				}}
			>
				<label class="input input-bordered flex items-center gap-2 mb-4" for="username">
					<Avatar width="w-4" height="h-4" opacity="opacity-70"></Avatar>
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
					<Key width="w-4" height="h-4" opacity="opacity-70"></Key>
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
						<Info stroke="stroke-error" width="w-6" height="h-6"></Info>
						<span class="text-md">Credenciales incorrectas!</span>
					</div>
				{/if}

				<button class="btn btn-primary" type="submit">Ingresar</button>
			</form>
		</div>
	</div>
</div>
