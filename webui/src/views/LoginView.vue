<template>

  <div class="container">
    <div class="center">
      <input type="text" v-model="textFieldValue" placeholder="Enter text here" />
    </div>
    <div class="side">
      <button @click="submitText">Submit</button>
    </div>
    <div class= "titleC">
    <h1 class="title">{{ pageTitle }}</h1>
    </div>
  </div>

</template>

<script>

export default {
  data() {
    return {
      pageTitle: "WasaPhoto",
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
  margin: 0;
  flex-shrink: 0; 
  background-color: rgba(255,0,255,100)
}
.titleC{
  position: absolute;
  top: 1;
  left: 1;
  z-index: 1;
}
.title{
  color: rgba(255, 0, 0, 100);/* Change the color of the heading text */
  font-family: Arial, sans-serif; /* Change the font family of the heading text */
  font-size: 24px; /* Change the font size of the heading text */
  font-weight: bold; /* Change the font weight of the heading text */
  /* Additional styles for the heading */

}



</style>
