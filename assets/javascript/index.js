window.onload = function () {

var mainContents = new Vue({
  el:'#contentsWrapper',
  data:{
    musics:{},
    post_mode:false,
    zoomed:false
  },
  beforeCreate:function(){
    self=this;
    buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/music/new",null, function(ev){
      self.musics = this.response;
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

      location.href = "http://localhost:1323/test";

    },
    postMode:function(){
      self.post_mode = !(self.post_mode);
      console.log(self.post_mode);
    },
    sentenceAdju:function(sentence,max_length){
      return sentence.slice(0,max_length);
    },
    searchWord:function(){
      var words = document.getElementById("searchData").value;
      words=words.replace(/　/g,"&");
      words=words.replace(/ /g,"&");
      var data = words.split("&");
      for( var i=0; i<data.length;i++){
        words=words.replace(data[i],"s"+i+"="+data[i]);
      }
      buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/music/search?"+words,null, function(ev){
        self.musics=this.response;
      });
    },
    zoom:function(music){
      var el = document.getElementById("zoomImage");
      el.src = self.getImagePath(music.content);
      self.zoomed = true;
    }
  }
});

}
