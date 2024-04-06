<template>
  <div class="layout">
  <button class="homeButton" @click="backHome">Home</button>
  <div class="navigButton">
      <button @click="toHome">HOME</button>
      <button @click="toProfile">PROFILE</button>
      <button @click="toLogout">LOGOUT</button>
  </div>
  </div>


</template>

<style>
.navigButton{
  position:absolute;
  left: 50%;
  bottom: 0%;
  transform: translate(-50%, -0%);
}
.layout{
  position: relative;
  height: 100vh;
  width: 100vw;
  background-color: rgba(71, 38, 77,100);

}
.homeButton{
  color: rgba(0,255,0,100);
  width: 20vw;
  height:15vh;
}


</style>

<script>
export default {
  data() {
    return {
      
      followers:10,
      followed:10,
      profilename: 'EO',
      user: localStorage.getItem('username').replace('"', '').replace('"', ''),
    };
  },
  mounted() {
    
  },
  methods: {
    toHome(){
      this.$router.push({path:'/'+this.user+'/home'});
    },
    toProfile(){
      this.$router.push({path:'/'+this.user+'/profile'});
    },
    toLogout(){
      this.$router.push({path:'/'});

    },
    backHome(){
      this.$router.push({path:'/'+this.user+'/home'});

    },
    /*async fetchProfileData() {
        const response = await this.$axios.get("/userActions/"+this.user+"/interactions/Profile/"+this.profileId, { headers: {"Authorization" : this.profileId}});
        console.log(response.data);
        this.profilename =response.data.id.idUser;
        if (response.data.followers==null){
          this.followers=0;}
        else{
          this.followers= response.data.followers.length;}
        if (response.data.follows==null){
          this.followed=0;}
        else{this.followed=response.data.follows.length;}
        }*/
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
