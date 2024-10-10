<!-- src/routes/+layout.svelte -->
<script>
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import '../app.css';
</script>


<div class="bg-secondary font-fira-code">
    <Header />
	<main class="flex-grow">
		<slot />
	</main>
	<Footer />
</div>
