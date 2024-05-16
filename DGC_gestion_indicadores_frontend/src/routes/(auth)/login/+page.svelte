<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import * as Card from '$lib/components/ui/card/index.js';
	import { onMount } from 'svelte';
	import type { PageData } from './$types.js';
	import LoginForm from './loginForm.svelte';
	export let data: PageData;

	onMount(()=>{
		invalidateAll()
	})

</script>

<svelte:head>
	<title>Ingreso</title>
</svelte:head>

<div class="flex h-screen w-screen">
	<div
		class="bg-primary flex h-full w-2/5 flex-col items-center justify-center p-5 text-center"
	>
		<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-primary-foreground">
			Gestión de indicadores
		</h1>
		<p class="text-primary-foreground text-sm">Dirección de Gestión de Calidad</p>
	</div>
	<div class="flex h-full w-3/5 items-center justify-center">
		<Card.Root class="w-full max-w-sm">
			<Card.Header>
				<Card.Title class="text-2xl">Ingreso</Card.Title>
				<Card.Description>Ingresar el nombre de usuario para iniciar sesión.</Card.Description>
			</Card.Header>
			<Card.Content class="grid gap-4">
				<LoginForm data={data.loginForm} />
			</Card.Content>
		</Card.Root>
	</div>
</div>
