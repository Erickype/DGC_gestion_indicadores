<script lang="ts">
	import { enhance } from '$app/forms';
	import Avatar from '$lib/icons/avatar.svelte';
	import Info from '$lib/icons/info.svelte';
	import Key from '$lib/icons/key.svelte';
	import Letter from '$lib/icons/letter.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<svelte:head>
	<title>Registrarse</title>
</svelte:head>

<div class="flex items-center justify-center mt-10">
	<div class="card card-bordered bg-accent/10 border-secondary w-96 text-secondary-content">
		<div class="card-body items-center text-center">
			<Avatar fill="fill-primary" stroke="secondary-content/10" width="w-12" height="h-12"></Avatar>
			<h1 class="card-title text-4xl font-bold mb-4">Crear cuenta!</h1>

			<form action="?/register" method="POST" use:enhance>
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

				<label class="input input-bordered flex items-center gap-2 mb-4" for="email">
					<Letter width="w-4" height="h-6" opacity="opacity-70"></Letter>
					<input class="glow" placeholder="E-mail" id="email" name="email" type="email" required />
				</label>

				<label class="input input-bordered flex items-center gap-2 mb-4" for="password">
					<Key width="w-4" height="h-4" opacity="opacity-70"></Key>
					<input
						class="glow"
						placeholder="Password"
						id="password"
						name="password"
						type="password"
						required
					/>
				</label>

				{#if form?.error}
					<div role="alert" class="alert alert-warning mb-4">
						<Info stroke="stroke-error" width="w-6" height="h-6"></Info>
						<span class="text-md">Error registrando usuario!</span>
					</div>
				{/if}

				<button class="btn btn-primary" type="submit">Registarse</button>
			</form>
		</div>
	</div>
</div>
