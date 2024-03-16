<template>

  <div class="container">
    <div class="userinfo">
      <p class="name">{{ "Username: "+profilename }}</p>
      <p class="follower">{{ "Followers: "+followers }}</p>
      <p class="followed">{{ "Follows: "+followed }}</p>
    </div>
      
    <div class="photoUp">
<input type="file" @change="handleFileUpload" ref="fileInput" accept="image/*">
  <button @click="uploadPhoto">Upload</button>
    <div class="photo">
      <div class="photo-stack">
          <div class="photo-container" v-for="(photo, index) in stack" :key="index">
            <img :src=photo alt="Uploadedphoto">
          </div>
    
      </div>

    </div>
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
.photo-stack {
  width: 100%;
  max-height: 300px; /* Adjust the maximum height as needed */
  overflow-y: auto; /* Enable vertical scrolling */
}

.photo-container {
  margin-bottom: 10px; /* Adjust margin between photos as needed */
}

.photo-container img {
  width: 100%;
  height: auto;
}

</style>

<script>
import { User } from '../scripts/myStructs.js';
export default {
  data() {
    return {
      randomString: "1",
      binaryIm:"10",
      file: null,
      uploadedImageUrl: null,
      filePreview: null,
      stack:[],
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
            this.stack.push(response3.data[i].image);
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
