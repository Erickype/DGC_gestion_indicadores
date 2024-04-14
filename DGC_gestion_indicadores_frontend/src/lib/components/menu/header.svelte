<script lang="ts">
	import Search from 'lucide-svelte/icons/search';
	import CircleUser from 'lucide-svelte/icons/circle-user';

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
