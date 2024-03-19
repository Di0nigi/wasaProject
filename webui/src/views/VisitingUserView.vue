<template>
  <div class="layout">
    <div class="container">
      <div class="userinfo">
        <p class="name">{{ "Username: "+profilename }}</p>
        <p class="follower">{{ "Followers: "+followers }}</p>
        <p class="followed">{{ "Follows: "+followed }}</p>
      </div>
    </div>
    <div class="follow">
      <button class="followBt" @click="toFollow" :style="btModel" :key="btKey">{{ followState }}</button>

    </div>

    <div class="photo-stack">
            <div class="photo-container" v-for="(photo, index) in stack" :key="index">
              <button class="buttonIm" @click="toComm(photo.id)">
                <img :src=photo.url alt="Uploadedphoto">
              </button>
            </div> 
    </div>
  </div>


</template>

<style>
.layout{
  position: relative;
  height: 100vh;
  width: 100vw;
  background-color: rgba(71, 38, 77,100);

}

.container {
  height: 100vh;
  width: 100vw;
  
}
.userinfo{
  position: absolute;
  left: 0;
  top: 0;
  background-color: rgba(255, 255, 255,0);
}
.name{
  color: rgba(255, 255, 255,100);
}
.followed{
  color: rgba(255, 255, 255,100);
}
.follower{
  color: rgba(255, 255, 255,100);
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

.photo-container {
  margin-bottom: 10px; /* Adjust margin between photos as needed */
}

.photo-container img {
  width:100%;
}
.follow{
  position:absolute;
  right:0;
  top:0;
}
.followBt{

  color: rgba(0,255,0,100);
  width:15vw;
  height:10vh;
}

</style>

<script>
import { slideShowim } from '../scripts/myStructs.js';
export default {
  data() {
    return {
      btKey: 0,
      followVal:true,
      followState: "",
      randomString: "1",
      binaryIm:"10",
      file: null,
      uploadedImageUrl: null,
      filePreview: null,
      stack:[],
      followers:10,
      followed:10,
      profilename: 'EO',
      user: localStorage.getItem('username').replace('"', '').replace('"', ''),
      profileId : this.$route.path.split("/").slice(-1)[0],
    };
  },
  mounted() {
    this.fetchProfileData();
  },
  methods: {
    async checkFollow(){
      try{
      const response4 = await this.$axios.get("/userActions/"+this.user+"/interactions/followingActions/"+this.profileId, { headers: {"Authorization" : this.user}});
      console.log(response4);
      this.followVal=false;
      console.log("triggered");
     /* console.log(this.followVal);*/
      
      }
      catch(error){
        console.log("triggered by err");
        this.followVal=true;
      }


    },
    async toFollow(){
      await this.checkFollow();
      console.log("in bt: "+this.followVal)
      
      
      if (this.followVal==true){
        let response9 = await this.$axios.post("/userActions/"+this.user+"/interactions/followingActions", JSON.stringify({idUser: this.profileId}),{ headers: {"Authorization" : this.user}})
        console.log(response9);
        this.followState="Follow";
        
        this.btKey++;
        }
      else{
        let response10 = await this.$axios.delete("/userActions/"+this.user+"/interactions/followingActions/"+this.profileId,{ headers: {"Authorization" : this.user}});
        console.log(response10);
        this.followState="Unfollow";
        
        
        this.btKey++;
      }
    },
    toComm(photoName){
      console.log("clicked");
      this.$router.push({path: '/'+this.user+'/profile/iteract/'+this.profileId+'/'+photoName})    
     
  },
    async fetchProfileData() {
      try {
        await this.checkFollow();
        console.log("here: "+this.followVal)
        
        if(this.followVal===true){
         /* console.log("true? ");
          console.log(this.followVal);*/
          this.followState="Follow";
          console.log("should be follow");
          this.btKey++;
        }else{
         /* console.log("false? ");
          console.log(this.followVal);*/
          this.followState="Unfollow";
          console.log("should be Unfollow");
          this.btKey++;
        }
        const response = await this.$axios.get("/userActions/"+this.user+"/interactions/Profile/"+this.profileId, { headers: {"Authorization" : this.profileId}});
        console.log(response.data);
        
        this.profilename =response.data.id.idUser;
        if (response.data.followers==null){
          this.followers=0;
        }
        else{
          this.followers= response.data.followers.length;
        }
        
        if (response.data.follows==null){
          this.followed=0;
        }
        else{this.followed=response.data.follows.length;}

        const response3 = await this.$axios.get("/userActions/"+this.profileId, { headers: {"Authorization" : this.profileId}});
        console.log(response3.data);
        if(response3.data!=null){
          console.log(response3.data.length);
          for(let i=0; i<response3.data.length;i++){
            try{
            var im= new slideShowim(response3.data[i].idPhoto.idObj,response3.data[i].image);
            this.stack.push(im);
            }
            catch (error) {
        console.error("Error fetching profile data:", error);
      }
           
          }

        }

      } catch (error) {
        console.error("Error fetching profile data:", error);
      }
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
}
</style>
