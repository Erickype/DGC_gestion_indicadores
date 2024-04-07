<script lang="ts">
	import Avatar from '$lib/icons/avatar.svelte';
	import Key from '$lib/icons/key.svelte';
	import Letter from '$lib/icons/letter.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<h1 class="text-4xl font-bold">Registrarse</h1>

<form action="?/register" method="POST">
	<label class="input input-bordered flex items-center gap-2" for="username">
		<Avatar customClass="w-4 h-4 opacity-70"></Avatar>
		<input class="glow" placeholder="Usuario" id="username" name="username" type="text" required />
	</label>

	<label class="input input-bordered flex items-center gap-2" for="email">
		<Letter customClass="w-4 h-4 opacity-70"></Letter>
		<input class="glow" placeholder="Email" id="email" name="email" type="email" required />
	</label>

	<label class="input input-bordered flex items-center gap-2" for="password">
		<Key customClass="w-4 h-4 opacity-70"></Key>
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
		<p class="error">{form?.error}</p>
	{/if}

	<button class="btn btn-success" type="submit">Enviar</button>
</form>
