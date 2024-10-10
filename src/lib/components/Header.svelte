<!-- src/lib/components/Header.svelte -->
<script>
	import { page } from '$app/stores';
	import { slide } from 'svelte/transition';

	let isOpen = false;

	$: currentPath = $page.url.pathname;

	function toggleMenu() {
		isOpen = !isOpen;
	}
</script>

<nav
	class="container fixed bg-secondary z-[999] mx-auto py-3 px-4 lg:px-52 md:static md:h-auto md:w-auto"
>
	<div class="flex items-center justify-between flex-wrap">
		<!-- Logo -->
		<div class="flex gap-x-3">
			<img src="/images/Logo.svg" alt="logo" />
			<p class="text-white">Syndrom</p>
		</div>

		<!-- Mobile nav button -->
		<button class="relative group md:hidden" on:click={toggleMenu}>
			<div
				class="relative flex overflow-hidden items-center justify-center rounded-full w-[50px] h-[50px] transform transition-all duration-200"
			>
				<div
					class="flex flex-col justify-between w-[20px] h-[20px] transform transition-all duration-300 origin-center overflow-hidden"
				>
					<div
						class="bg-white h-[2px] w-7 transform transition-all duration-300 origin-left"
						class:translate-y-6={isOpen}
						class:delay-100={isOpen}
					></div>
					<div
						class="bg-white h-[2px] w-7 rounded transform transition-all duration-300"
						class:translate-y-6={isOpen}
						class:delay-75={isOpen}
					></div>
					<div
						class="bg-white h-[2px] w-7 transform transition-all duration-300 origin-left"
						class:translate-y-6={isOpen}
						class:delay-100={isOpen}
					></div>

					<!-- Close btn -->
					<div
						class="absolute items-center justify-between transform transition-all duration-500 top-2.5 -translate-x-10 flex w-0"
						class:w-12={isOpen}
						class:translate-x-0={isOpen}
					>
						<div
							class="absolute bg-white h-[2px] w-5 transform transition-all duration-500 rotate-0 delay-300"
							class:rotate-45={isOpen}
						></div>
						<div
							class="absolute bg-white h-[2px] w-5 transform transition-all duration-500 -rotate-0 delay-300"
							class:-rotate-45={isOpen}
						></div>
					</div>
				</div>
			</div>
		</button>

		<!-- Links -->
		<div class="hidden md:block">
			<div class="flex justify-center gap-5 mt-5 text-2xl md:text-lg md:space-x-4">
				<a
					href="/"
					class:text-white={currentPath === '/'}
					class:text-tertiary={currentPath !== '/'}
					class:hover:text-white={currentPath !== '/'}
					><span class="text-primary">#</span>home</a
				>
				<a
					href="/projects"
					class:text-white={currentPath === '/projects'}
					class:text-tertiary={currentPath !== '/projects'}
					class:hover:text-white={currentPath !== '/projects'}
					><span class="text-primary">#</span>works</a
				>
				<a
					href="/aboutme"
					class:text-white={currentPath === '/aboutme'}
					class:text-tertiary={currentPath !== '/aboutme'}
					class:hover:text-white={currentPath !== '/aboutme'}
					><span class="text-primary">#</span>about-me</a
				>
				<a
					href="/contact"
					class:text-white={currentPath === '/contact'}
					class:text-tertiary={currentPath !== '/contact'}
					class:hover:text-white={currentPath !== '/contact'}
					><span class="text-primary">#</span>contacts</a
				>
			</div>
		</div>
	</div>

	{#if isOpen}
		<div transition:slide={{ duration: 300 }} class="md:hidden">
			<div class="flex flex-col justify-center flex-wrap gap-5 mt-5 text-2xl">
				<a
					href="/"
					class="w-[100%]"
					class:text-white={currentPath === '/'}
					class:text-tertiary={currentPath !== '/'}
					class:hover:text-white={currentPath !== '/'}
					><span class="text-primary">#</span>home</a
				>
				<a
					href="/projects"
					class="w-[100%]"
					class:text-white={currentPath === '/projects'}
					class:text-tertiary={currentPath !== '/projects'}
					class:hover:text-white={currentPath !== '/projects'}
					><span class="text-primary">#</span>works</a
				>
				<a
					href="/aboutme"
					class="w-[100%]"
					class:text-white={currentPath === '/aboutme'}
					class:text-tertiary={currentPath !== '/aboutme'}
					class:hover:text-white={currentPath !== '/aboutme'}
					><span class="text-primary">#</span>about-me</a
				>
				<a
					href="/contact"
					class="w-[100%]"
					class:text-white={currentPath === '/contact'}
					class:text-tertiary={currentPath !== '/contact'}
					class:hover:text-white={currentPath !== '/contact'}
					><span class="text-primary">#</span>contacts</a
				>
			</div>
		</div>
	{/if}
</nav>
