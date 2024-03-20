<template>

  <div class="container">
    <div class="profile">
      <button @click="toProfilePage">Profile</button>
    </div>
    <div class="userSearch">
      <input type="text" v-model="inputText" class="txtInp">
      <button @click="searchUser">Search</button>
      <p class="warning" v-if="userNotfound" :key="warnKey">User not found</p>
    </div>
    
    <div class="photo-stack">
            <div class="photo-container" v-for="(photo, index) in streamStack" :key="index">
            <button class="photoOwner" @click="toOwnerProfile(photo.owner)">{{ photo.owner }}</button>
              <button class="buttonIm" @click="toComm(photo.id, photo.owner)">
                <img :src=photo.url alt="Uploadedphoto">
              </button>
            </div> 


    </div>
  </div>

</template>

<script>
import { slideShowim } from '../scripts/myStructs.js';

export default {
  
  data() {
    return {
      streamStack:[],
      pageTitle: "WasaPhoto",
      user: localStorage.getItem('username').replace('"', '').replace('"', ''),
      inputText: "",
      userNotfound:false,
      warnKey:0,
    };
  },
   mounted() {
    this.fetchData();

  },
  methods: {
    toOwnerProfile(profileName){
      this.$router.push({path:'/'+this.user+'/profile/iteract/'+profileName});

    },
    toComm(photoName,photoOwner){
      this.$router.push({path: '/'+this.user+'/profile/iteract/'+photoOwner+'/'+photoName});
    },
    async fetchData(){
      console.log("fetched");
      const response5 = await this.$axios.get("/userActions/"+this.user+"/interactions/Profile/"+this.user, { headers: {"Authorization" : this.user}});
      console.log(response5.data);
      for(let i=0; i<response5.data.follows.length;i++){
        console.log("eme"+i);
        const response6 = await this.$axios.get("/userActions/"+response5.data.follows[i].idUser, { headers: {"Authorization" : response5.data.follows[i].idUser}});
        console.log(response6.data);
        var im= new slideShowim(response6.data.slice(-1)[0].idPhoto.idObj,response6.data.slice(-1)[0].image);
        im.owner=response5.data.follows[i].idUser;
        this.streamStack.push(im);
      }

    },
    toProfilePage(){
      this.$router.push({path: '/'+this.user+'/profile'});

    },
    async searchUser(){
      try{
        const response = await this.$axios.get("/userActions/"+this.user+"/interactions/Profile/"+this.inputText, { headers: {"Authorization" : this.user}});
        console.log(response.data);
        this.userNotfound=false;
        this.$router.push({path: '/'+this.user+'/profile/iteract/'+response.data.id.idUser});
        console.log("here");

      }
      catch(error){
        this.userNotfound=true;
        this.warnKey++;
        

      }
    }
  }
};
</script>

<style scoped>
.container{
  position:relative;
  width: 100vw;
  height: 100vh;
  background-color: rgba(71, 38, 77,100)
}
.profile{
  position:absolute;
  left:0;
  top:0;
  width: 20vw;
  height:15vh;
}
.userSearch{
  position:absolute;
  right:0;
  top:0;

}
.warning{
  color: rgba(255,255,255,100);
}
.photoOwner{
  padding-top:0px;
  

}
.photo-stack {
  padding-top: 20px;
  max-height: 500px;
  max-width: 400px;
  overflow-y: auto; /* Enable vertical scrolling */
  position: absolute;
  left: 50%;
  top: 10%;
  transform: translate(-50%, -10%);
}

</style>
