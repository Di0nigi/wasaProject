<template>
 <div class="layout">
    <div class="likeBts">
      <button @click="LikePost" class="like">Like</button>
      <button @click="UnLikePost" class="unLike">unlike</button>
      <p class="nLike">{{ nLikes+" people liked your post!" }}</p>

    </div>
    <div class="imageContainer">
    <img :src=photo alt="Uploadedphoto">
    </div>
    <div class="comments">
      <div class= "make comment">
        <button @click="uploadComment">comment</button>
        <input type="text" v-model="inputText" class="txtInp">
      </div>
      <div class="commDisplay" v-for="(comm, index) in commStack" :key="index">
      <div comContainer>
        <p class="sComment">{{ comm.owner+": "+comm.content }}</p>
        <button @click="delComment( comm.idComment, comm.owner )">delete</button>
      </div>
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
.nLike{
  color: rgba(255,255,255,100);

}
.likeBts{
  position: absolute;
  left:0;
  top:0;
  
}
.like{
  width:10vw;
  height:6vh;
  background-color: rgba(0,255,0,100);
}
.unLike{
  width:10vw;
  height:6vh;
  background-color: rgba(255,0,0,100);
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
import { dispComment } from '../scripts/myStructs.js';
export default {
  data() {
    return {
      valLike:false,
      commStack:[],
      photoId: "",
      photo: null,
      user : localStorage.getItem('username').replace('"', '').replace('"', ''),
      randomString: "",
      inputText: "",
      nLikes:0,
      };
  },
  mounted() {
    this.fetchData();
    console.log("called");
  },
methods: {
  async delComment(commId,owner){
    if (this.user==owner){

    console.log(commId);
    const response7= await this.$axios.delete("/userActions/"+this.user+"/interactions/comments/"+this.photoId+"/actions/"+commId,{ headers: {"Authorization" : this.user}});
    console.log(response7);}
    
  },
  async checkLikevalidity(){
    try{
    const response4 = await this.$axios.get("/userActions/"+this.user+"/interactions/likes/"+this.photoId,{ headers: {"Authorization" : this.user}});
    console.log(response4);
    if (response4.status == 200  || response4.status == 204) {
      this.valLike=true;
      }
        }
    catch (error) {
      this.valLike=false;

    }
  },
  async LikePost(){
    await this.checkLikevalidity();
    console.log(this.valLike);
    if(this.valLike==false){
      this.generateRandomString();
    const response3 = await this.$axios.post("/userActions/"+this.user+"/interactions/likes",JSON.stringify({idLike: { idObj: this.randomString}, owner: {idUser: this.user}, toPhoto: { idObj: this.photoId}}),{ headers: {"Authorization" : this.user}});
    console.log(response3);}
  },
  async UnLikePost(){
    await this.checkLikevalidity();
    console.log(this.valLike);
    if (this.valLike==true){
    const response6 = await this.$axios.get("/userActions/"+this.user+"/interactions/likes/"+this.photoId,{ headers: {"Authorization" : this.user}});
    console.log(response6);
    const response5 = await this.$axios.delete("/userActions/"+this.user+"/interactions/likeaction/"+this.photoId+"/actions/"+response6.data.idLike.idObj+"/likes",{ headers: {"Authorization" : this.user}});
    console.log(response5);}

  },
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
          var cm=new dispComment( response.data.comments[i].owner.idUser,response.data.comments[i].content,response.data.comments[i].idComment.idObj);
          this.commStack.push(cm);
        }
      }

      if (response.data.likes!=null){
      this.nLikes=response.data.likes;
      }
    },
    async uploadComment(){
      this.generateRandomString();
      const response2 = await this.$axios.post("/userActions/"+this.user+"/interactions/comments",JSON.stringify({idComment: { idObj: this.randomString }, content: this.inputText, owner: {idUser: this.user}, photo: {idObj: this.photoId} }),{ headers: {"Authorization" : this.user}});
      
    },
  }
};

</script>

<style scoped>





</style>
