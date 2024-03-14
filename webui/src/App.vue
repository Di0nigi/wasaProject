<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
	<div class="container">
		<RouterView />
	</div>
</template>

<style>
.container{
	display: flex;
  	justify-content: center;
  	align-items: center;
  	height: 100vh;
  	width: 100vw;
	margin: 0;
	flex-shrink: 0;
	max-width: 2000px; 
	
	 
}
</style>
