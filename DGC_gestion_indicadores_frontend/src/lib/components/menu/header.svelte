<script lang="ts">
	import Search from 'lucide-svelte/icons/search';
	import CircleUser from 'lucide-svelte/icons/circle-user';

	import Sun from 'lucide-svelte/icons/sun';
	import Moon from 'lucide-svelte/icons/moon';

	import { toggleMode } from 'mode-watcher';

	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { enhance } from '$app/forms';
</script>

<header class="bg-muted/40 flex h-14 items-center gap-4 border-b px-4 lg:h-[60px] lg:px-6">
	<div class="w-full flex-1">
		<!-- <form>
			<div class="relative">
				<Search class="text-muted-foreground absolute left-2.5 top-2.5 h-4 w-4" />
				<Input
					type="search"
					placeholder="Search products..."
					class="bg-background w-full appearance-none pl-8 shadow-none md:w-2/3 lg:w-1/3"
				/>
			</div>
		</form> -->
	</div>
	<Button on:click={toggleMode} variant="outline" size="icon">
		<Sun
			class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
		/>
		<Moon
			class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
		/>
		<span class="sr-only">Cambiar tema</span>
	</Button>
	<DropdownMenu.Root>
		<DropdownMenu.Trigger asChild let:builder>
			<Button builders={[builder]} variant="secondary" size="icon" class="rounded-full">
				<CircleUser class="h-5 w-5" />
				<span class="sr-only">Menu usuario</span>
			</Button>
		</DropdownMenu.Trigger>
		<DropdownMenu.Content align="end">
			<DropdownMenu.Label>Cuenta</DropdownMenu.Label>
			<DropdownMenu.Separator />
			<DropdownMenu.Item asChild>
				<form action="/logout" method="POST" use:enhance>
					<Button
						type="submit"
						variant="ghost"
						class="h-8 w-full justify-start px-2 text-sm font-light">Salir</Button
					>
				</form>
			</DropdownMenu.Item>
		</DropdownMenu.Content>
	</DropdownMenu.Root>
</header>
