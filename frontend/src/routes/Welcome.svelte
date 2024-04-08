<script lang="ts">
	export let project = 'Swole Goal'
	let age = 0
	$: calcedAge = age < 1 ? 0 : age * 6
	$: if (age < 0) {
		alert(`your cats input age ${age} is invalid`)
	}
</script>

<h1>Welcome to project {project}</h1>
<p>Visit <a href="https://github.com/emoral435/swole-goal">the repository</a> to find out more about this project!</p>
<a href="/about">Click here to visit the about page.</a>

<input type="number" bind:value={age} />
<!-- <button on:click={calculateAge}>Calculate</button> -->
<p>{ calcedAge }</p>