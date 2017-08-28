window.onload = function () {

var mainContents = new Vue({
  el:'#contentsWrapper',
  data:{
    user:{},
    musics:{},
    notices:{},
    post_mode:false,
    zoomed:false
  },
  beforeCreate:function(){
    self=this;
    buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/getUser",null, function(ev){
      self.user = this.response;
    });
    buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/music/new",null, function(ev){
      self.musics = this.response;
    });
    buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/notice",null, function(ev){
      self.notices = this.response;
    });
  },
  methods:{
    getImagePath:function(image){
      return "music/picture" + image;
    },
    postImage:function(){
      var formData = new FormData();
      var name = document.getElementById("inputName").value;
      var description = document.getElementById("description").value;
      var tags = document.getElementById("musicTags").value;
      var image = document.getElementById("inputImage").files[0];
      formData.append("music_name",name);
      formData.append("tags",tags);
      formData.append("description",description);
      formData.append("file",image);

      buildXHR("POST","", {},window.location.href + "/music",formData, function(ev){
      });

      location.href = location.href;

    },
    postMode:function(){
      self.post_mode = !(self.post_mode);
      console.log(self.post_mode);
    },
    sentenceAdju:function(sentence,max_length){
      return sentence.slice(0,max_length);
    },
    searchWord:function(){
    },
    zoom:function(music){
      var el = document.getElementById("zoomImage");
      el.src = self.getImagePath(music.music_id);
      self.zoomed = true;
    },
    zoomout:function(){
      var el = document.getElementById("zoomImage");
      el.src = "";
      self.zoomed = false;
    },
    favorite:function(){

      var src = document.getElementById("zoomImage").src;
      var splitData = src.split("/");
      var id = "/" + splitData[splitData.length - 2] + "/" + splitData[splitData.length - 1]
      var data = {
        "music_id":id
      }
      console.log(data)
      buildXHR("POST","json", { "Content-Type":"application/json" },window.location.href+"/fav",JSON.stringify(data), function(ev){
      });
    },
    getFavorited:function(){
      self = this;
      buildXHR("GET","json", { "Content-Type":"application/json" },window.location.href+"/fav",null, function(ev){
        self.musics = this.response;
      });
    },
    getMusicHome:function(){
      buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/music/new",null, function(ev){
        self.musics = this.response;
      });
    }
  }
});
}
