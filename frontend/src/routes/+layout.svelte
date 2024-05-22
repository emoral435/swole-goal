<!-- theming for the entire webapge -->
<svelte:head>
	<!-- SMUI Styles -->
	{#if darkTheme === undefined}
	<link rel="stylesheet" href="/smui.css" media="(prefers-color-scheme: light)" />
	<link rel="stylesheet" href="/smui-dark.css" media="screen and (prefers-color-scheme: dark)" />
	{:else if darkTheme}
	<!-- this is here so that when the user prints page, light theme is used -->
	<link rel="stylesheet" href="/smui-dark.css" media="screen" />
	<!-- this is what the user sees on their screen -->
	<link rel="stylesheet" href="/smui.css" media="screen" />
	{:else}
	<link rel="stylesheet" href="/smui-dark.css" media="screen" />
	{/if}
</svelte:head>

<!-- Navbar for the landing page -->
<div class="landing-navbar">
	<TopAppBar variant="static"
		color={darkTheme ? 'secondary' : 'primary'}>
		<Row>
			<Section>
				
			</Section>
			<Section>
				<Button on:click={() => goto('/')}
					aria-label="Home">
					SWOLE GOAL
				</Button>
			</Section>
			<Section align="end" toolbar>
				<Button on:click={() => goto('/about')}
					aria-label="About">
					ABOUT
				</Button>
				<Button on:click={() => goto('/login')}
					aria-label="Login">
					LOGIN
				</Button>
				<IconButton title={darkTheme ? 'Light mode' : 'Dark mode'}
					on:click={() => (darkTheme = !darkTheme)}>
					<Icon tag="svg"
						viewBox="0 0 24 24">
						<path fill="currentColor" d={darkTheme ? mdiWeatherSunny : mdiWeatherNight} />
					</Icon>
				</IconButton>
			</Section>
		</Row>
	</TopAppBar>
</div>

<slot></slot>

<script lang="ts">
	import { onMount } from 'svelte'
	import IconButton, { Icon } from '@smui/icon-button'
	import Button from "@smui/button"
	import { mdiWeatherSunny, mdiWeatherNight } from '@mdi/js';
	import TopAppBar, { Row, Section, Title } from '@smui/top-app-bar';
	import { goto } from '$app/navigation';

	let darkTheme: boolean | undefined = undefined
	let topScreen

	onMount(() => {
		darkTheme = window.matchMedia('(prefers-color-scheme: dark)').matches
	})
</script>

<style lang="scss" scoped>
	.landing-navbar {
		font-family: "Comic Sans MS", "Comic Sans", cursive;
	}
</style>