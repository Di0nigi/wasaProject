<template>

  <div class="container">
    <div class="userinfo">
      <p class="name">{{ "Username: "+profilename }}</p>
      <p class="follower">{{ "Followers: "+followers }}</p>
      <p class="followed">{{ "Follows: "+followed }}</p>
    </div>

  </div>

</template>

<style>
.container {
  height: 100vh;
  width: 100vw;
  background-color: rgba(71, 38, 77,100);
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

</style>

<script>
import { User } from '../scripts/myStructs.js';
export default {
  data() {
    return {
      followers:10,
      followed:10,
      profilename: 'EO',
      profileData: User,
      user: localStorage.getItem('username').replace('"', '').replace('"', ''),
    };
  },
  mounted() {
    this.fetchProfileData();

  },
  methods: {
    async fetchProfileData() {
      try {
        const response = await this.$axios.get("/userActions/"+this.username+"/interactions/Profile/"+this.user);
        console.log(response.data);
        //this.profileData.id = response.data.id.userId*/
        this.profilename =response.data.id.idUser;
        if (response.data.followers==null){
          this.followers=0;
        }
        else{this.followers= response.data.followers.length;}
        
        if (response.data.follows==null){
          this.followed=0;
        }
        else{this.followed=response.data.follows.length;}
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
