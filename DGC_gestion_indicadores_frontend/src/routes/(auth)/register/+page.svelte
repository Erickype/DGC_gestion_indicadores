<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import type { PageData } from './$types.js';
	import RegisterForm from './registerForm.svelte';
	export let data: PageData;
</script>

<svelte:head>
	<title>Registro</title>
</svelte:head>

<div class="flex h-screen w-screen">
	<div
		class="bg-muted flex h-full w-2/5 flex-col items-center justify-center p-5 text-center"
	>
		<h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-muted-foreground">
			Gestión de indicadores
		</h1>
		<p class="text-muted-foreground text-sm">Dirección de Gestión de Calidad</p>
	</div>
	<div class="flex h-full w-3/5 items-center justify-center">
		<Card.Root class="w-full max-w-sm">
			<Card.Header>
				<Card.Title class="text-2xl">Registro</Card.Title>
				<Card.Description>Ingresar datos para registrarse.</Card.Description>
			</Card.Header>
			<Card.Content class="grid gap-4">
				<RegisterForm data={data.registerForm} />
			</Card.Content>
		</Card.Root>
	</div>
</div>
