<template>
  <div class="layout">
    <div class="container">
      <div class="userinfo">
        <p class="name">{{ "Username: "+profilename }}</p>
        <p class="follower">{{ "Followers: "+followers }}</p>
        <p class="followed">{{ "Follows: "+followed }}</p>
      </div>
    </div>
    <div class= "changUsername">
        <div v-if="showAlert" class="alert">
            {{ alertMessage }}
        <button @click="hideAlert">Close</button>
      </div>
      <button class changeUsButton @click="changeUser">Change username</button>
      <input type="text" v-model="inputText" class="txtInp">
    </div>

    <div class="photoUp">
        <input type="file" @change="handleFileUpload" ref="fileInput" accept="image/*" style="display: none;">
        <button @click="openFileInput">Choose File</button>
        <button @click="uploadPhoto" v-if="file!=null">Upload</button>
        <div class="photo"></div>
    </div>

    <div class="photo-stack">
            <div class="photo-container" v-for="(photo, index) in stack" :key="index">
              <button class="buttonIm" @click="toComm(photo.id)">
                <img :src=photo.url alt="Uploadedphoto">
              </button>
              <button class="deButton" @click="rightClickDel(photo.id)">Delete</button >
            </div> 
    </div>
    <div class="navigButton">
      <button @click="toHome">HOME</button>
      <button @click="toLogout">LOGOUT</button>
    </div>


  </div>


</template>

<style>
.alert{
  position: absolute;
  left:0;
  bottom:0;
  color: rgba(255,255,255,100);
  padding-bottom: 10px;

}
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
.changUsername{
  position: absolute;
  left:0;
  bottom:0;
}
.photoUp{
  position: absolute;
  right:0;
  top:0;
  background-color: rgba(255,255,255,100);
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

</style>

<script>
import { slideShowim } from '../scripts/myStructs.js';
export default {
  data() {
    return {
      showAlert: false,
      alertMessage: 'USER ALREADY EXISTS!',
      inputText: "",
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
    };
  },
  mounted() {
    this.fetchProfileData();
    window.addEventListener('popstate', this.handleBackButton);

  },
  methods: {
    reload(){
    window.location.reload();
    },
    toHome(){
      this.$router.push({path:'/'+this.user+'/home'});
    },
    toLogout(){
      this.$router.push({path:'/'});

    },
    launchAlert(){
      
      this.showAlert = true;

    },
    hideAlert() {
      this.showAlert = false;
    },




    handleBackButton() {
      // Perform your custom action when the back button is pressed
      // For example, navigate to the "NewPage" component
      this.$router.push({path:'/'+this.user+'/home'})
    },
    
    async changeUser(){
      if(this.inputText!=""){
      try{
      const response10 = await this.$axios.post("/userActions/"+this.user,JSON.stringify({idUser: this.inputText}),{ headers: {"Authorization" : this.user}});
      console.log(response10);
      localStorage.setItem('username', JSON.stringify(this.inputText));
      this.user=localStorage.getItem('username').replace('"', '').replace('"', '');
      this.$router.push({path:'/'+this.user+'/profile'});
      
      }
      catch(error){
        this.launchAlert();
      }
      }

    },
    async rightClickDel(photoName){ 
    const response7= await this.$axios.delete("/userActions/"+this.user+"/photoManager/"+photoName,{ headers: {"Authorization" : this.user}});
    console.log(response7);
    this.reload();

    },
    toComm(photoName){
      console.log("clicked");
      this.$router.push({path: '/'+this.user+'/profile/'+photoName})

    },
    openFileInput() {
      // Simulate click event on file input
      this.$refs.fileInput.click();
    },
    handleFileUpload(event) {
      this.file = event.target.files[0];
      console.log(this.file);
      this.filePreview = URL.createObjectURL(this.file)
      
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
    async uploadPhoto(){
      console.log("clicked");
      this.generateRandomString();
      console.log(this.randomString);
      await this.convertToBinaryString();
      const response2 = await this.$axios.post("/userActions/"+this.user+"/photoManager",JSON.stringify({idPhoto: { idObj: this.randomString }, owner: {idUser: this.user}, image: this.binaryIm, likes: 0, comments: null, numComments: 0}),{ headers: {"Authorization" : this.user}});
      console.log("here");
      console.log(response2);
      this.reload();
      
    },
    async convertToBinaryString() {
      if (!this.file) {
        alert('Please select a file.');
        return;
      }
      return new Promise((resolve, reject) => {
    const reader = new FileReader();

    // Set up onload event handler
    reader.onload = () => {
      const binaryString = reader.result;
      this.binaryIm = binaryString;
      console.log(this.binaryIm);
      resolve(); // Resolve the promise when reading is completed
    };

    // Set up onerror event handler
    reader.onerror = () => {
      reject(reader.error);
    };
     reader.readAsDataURL(this.file);
     });
  },
    async fetchProfileData() {
      try {
        const response = await this.$axios.get("/userActions/"+this.user+"/interactions/Profile/"+this.user, { headers: {"Authorization" : this.user}});
        console.log(response.data);
        //this.profileData.id = response.data.id.userId*/
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

        const response3 = await this.$axios.get("/userActions/"+this.user, { headers: {"Authorization" : this.user}});
        console.log(response3.data);
        if(response3.data!=null){
          console.log(response3.data.length);
          for(let i=0; i<response3.data.length;i++){
            try{
            /*var ur=new Blob([response3.data[i].image], { type: 'image/png' });
            var decodedImageUrl = URL.createObjectURL(ur);*/
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
