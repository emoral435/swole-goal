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

<nav>
	<a href="/">Home</a>
	<a href="/about">About</a>
	<a href="/settings">Settings</a>
	<IconButton title={darkTheme ? 'Light mode' : 'Dark mode'}
		on:click={() => (darkTheme = !darkTheme)}>
		<Icon tag="svg"
			viewBox="0 0 24 24">
			<path fill="currentColor" d={darkTheme ? mdiWeatherSunny : mdiWeatherNight} />
		</Icon>
	</IconButton>
</nav>

<slot></slot>

<script lang="ts">
	import { onMount } from 'svelte'
	import IconButton, { Icon } from '@smui/icon-button'
	import { mdiWeatherSunny, mdiWeatherNight } from '@mdi/js';

	let darkTheme: boolean | undefined = undefined

	onMount(() => {
		darkTheme = window.matchMedia('(prefers-color-scheme: dark)').matches
	})
</script>