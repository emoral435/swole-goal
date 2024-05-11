<div>Hello.</div>
<Welcome project="Not Swole Goal" />
<Button 
	on:click={() => alert('It worked!')}
	color="secondary">
	Click me!
</Button>
<Button 
	on:click={() => alert('It worked!')}
	color="primary">
	Click me too!
</Button>
<IconButton on:click={() => clicked++}>
	<Icon tag="svg" viewBox="0 0 24 24">
		<path fill="currentColor" d={mdiFormatColorFill} />
	</Icon>
</IconButton>

<script lang="ts">
	import Welcome from "./Welcome.svelte"
	import Button from "@smui/button"
	import IconButton, { Icon } from '@smui/icon-button'
	import { mdiFormatColorFill } from '@mdi/js';

	let clicked = 0
</script>
