<template>
 <div class="layout">
    <div class="imageContainer">
    <img :src=photo alt="Uploadedphoto">
    </div>
    <div class="comments">
      <div class= "make comment">
        <button @click="uploadComment">comment</button>
        <input type="text" v-model="inputText" class="txtInp">
      </div>
      <div class="commDisplay" v-for="(comm, index) in commStack" :key="index">
      <p class="sComment">{{ comm }}</p>
      </div>
    </div>
  </div>

</template>

<style>
.layout{
  flex-shrink: 0; 
  position: relative;
  height: 100vh;
  width: 100vw;
  background-color: rgba(71, 38, 77,100);
}
.imageContainer{
  position: absolute;
  left: 50%;
  top: 10%;
  transform: translate(-50%, -10%);
  
}
.imageContainer img{
  max-width:500px;
}
.comments{
  position:absolute;
  right:0;
  top:0;
  max-height: 500px;
   overflow-y: auto;
}
.commDisplay{
 
}
.sComment{
  color: rgba(255,255,255,100);

}


</style>

<script>
export default {
  data() {
    return {
      commStack:[],
      photoId: "",
      photo: null,
      user : localStorage.getItem('username').replace('"', '').replace('"', ''),
      randomString: "",
      inputText: "",
      };
  },
  mounted() {
    this.fetchData();
    console.log("called");
  },
methods: {
  generateRandomString() {
      const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
      let result = '';
      const charactersLength = characters.length;
      for (let i = 0; i < 5; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
      }
      this.randomString = result;
    },
    async fetchData(){
      var phId = this.$route.path.split("/").slice(-1)[0];
      this.photoId=phId;
      const response = await this.$axios.get("/userActions/"+this.user+"/photoManager/"+phId, { headers: {"Authorization" : this.user}});
      console.log(response.data);
      this.photo=response.data.image;
      if (response.data.comments!=null){
        for(let i=0; i<response.data.comments.length; i++){
          this.commStack.push(response.data.comments[i].content);
        }

      }
    },
    async uploadComment(){
      this.generateRandomString();
      const response2 = await this.$axios.post("/userActions/"+this.user+"/interactions/comments",JSON.stringify({idComment: { idObj: this.randomString }, content: this.inputText, owner: {idUser: this.user}, photo: {idObj: this.photoId} }),{ headers: {"Authorization" : this.user}});
      console.log(JSON.stringify({idComment: { idObj: this.randomString }, content: this.inputText, owner: {idUser: this.user}, photo: {idObj: this.photoId} }));
      console.log(response2);
    },
  }
};

</script>

<style scoped>





</style>
