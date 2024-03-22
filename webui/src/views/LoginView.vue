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
	  let response = this.$axios.post("/session", { idUser: this.textFieldValue})
    localStorage.setItem('username', JSON.stringify(this.textFieldValue));
	  this.$router.push({path: '/'+this.textFieldValue+'/home'})
    }
  }
};
</script>

<style scoped>
@font-face {
    font-family: "VP Pixel W03 Bold";
    src: url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.eot");
    src: url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.eot?#iefix")format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.woff2")format("woff2"),
    url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.woff")format("woff"),
    url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.ttf")format("truetype"),
    url("https://db.onlinewebfonts.com/t/a59a2fbd8417fe68d1259c51f8f03313.svg#VP Pixel W03 Bold")format("svg");
}
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
  
  color: rgba(255, 255, 255, 100);/* Change the color of the heading text */
  font-family: "VP Pixel W03 Bold"; /* Change the font family of the heading text */
  font-size: 120px; /* Change the font size of the heading text */
  
}



</style>
