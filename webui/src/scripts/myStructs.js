// ObjId struct equivalent in JavaScript
class ObjId {
    constructor(idObj) {
      this.idObj = idObj;
    }
  }
  
  // UserId struct equivalent in JavaScript
  class UserId {
    constructor(idUser) {
      this.idUser = idUser;
    }
  }
  
  // User struct equivalent in JavaScript
export class User {
    constructor(id, followers, follows, Posts, blockedArray) {
      this.id = new UserId(id); // UserId object
      this.followers = followers.map(id => new UserId(id)); // Array of UserId objects
      this.follows = follows.map(id => new UserId(id)); // Array of UserId objects
      this.Posts = Posts.map(id => new ObjId(id)); // Array of ObjId objects
      this.blockedArray = blockedArray.map(id => new UserId(id)); // Array of UserId objects
    }
  }
export class slideShowim{
    constructor(id, url) {
      this.id = id;
      this.url = url;
    }
  }
  
  // PostedImage struct equivalent in JavaScript
  class PostedImage {
    constructor(idPhoto, owner, image, likes, comments, numComments) {
      this.idPhoto = new ObjId(idPhoto); // ObjId object
      this.owner = new UserId(owner); // UserId object
      this.image = image;
      this.likes = likes;
      this.comments = comments.map(comment => new Comment(comment)); // Array of Comment objects
      this.numComments = numComments;
    }
  }
  
  // Comment struct equivalent in JavaScript
  class Comment {
    constructor(idComment, content, owner, photo) {
      this.idComment = new ObjId(idComment); // ObjId object
      this.content = content;
      this.owner = new UserId(owner); // UserId object
      this.photo = new ObjId(photo); // ObjId object
    }
  }
  
  // Like struct equivalent in JavaScript
  class Like {
    constructor(idLike, owner, toPhoto) {
      this.idLike = new ObjId(idLike); // ObjId object
      this.owner = new UserId(owner); // UserId object
      this.toPhoto = new ObjId(toPhoto); // ObjId object
    }
  }
  