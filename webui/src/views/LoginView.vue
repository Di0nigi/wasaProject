<template>

  <div class="container">
    <div class="center">
      <input type="text" v-model="textFieldValue" placeholder="Enter text here" />
    </div>
    <div class="side">
      <button @click="submitText">Submit</button>
    </div>
  </div>

</template>

<style>
.container{
	height: 100vh;
	width: 100vw;
	background-color: #fcba03;
}
</style>

<script>

export default {
  data() {
    return {
      textFieldValue: ''
    };
  },
  methods: {
    async submitText() {
      // Handle submission here
      console.log("Text submitted:", this.textFieldValue);
	  let response = this.$axios.post("/session", { userId: this.textFieldValue})
	  console.log(response)
	  this.$router.push({path: '/'+this.textFieldValue+'/profile'})
    }
  }
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
}


</style>
