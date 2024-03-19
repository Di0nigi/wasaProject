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
  </div>

</template>

<script>

export default {
  data() {
    return {
      pageTitle: "WasaPhoto",
      user: localStorage.getItem('username').replace('"', '').replace('"', ''),
      inputText: "",
      userNotfound:false,
      warnKey:0,
    };
  },
  methods: {
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

</style>
