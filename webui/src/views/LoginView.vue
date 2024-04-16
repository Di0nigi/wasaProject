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
	  let response = await this.$axios.post("/session", { idUser: this.textFieldValue})
    console.log(response.data);
    localStorage.setItem('username', JSON.stringify(this.textFieldValue));
    localStorage.setItem('token', response.data);

	  this.$router.push({path: '/'+this.textFieldValue+'/home'})
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
  background-color: rgba(71, 38, 77)
}
.titleC{
  position: absolute;
  top: 0;
  left: 0;
  z-index: 1;
  padding-left: 350px;

}
.title{
  color: rgba(255,255,255,100);
  font-size: 120px; 
  
}



</style>
