window.onload = function () {

var mainContents = new Vue({
    el:'#MainContents',
    data:{
        musics:{},
        items:[
            { message: 'Foo' },
            { message: 'Bar' }
        ]
    },
    beforeCreate:function(){
        self=this;
        buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "music/new",null, function(ev){
            self.musics = this.response;
            console.log(self.musics);
        });
    },
    methods:{
        postest:function(event){
            console.log("kita");
            var formData = new FormData();
            formData.append("userfile", document.getElementById("postImage"));
            buildXHR("POST","json", { "Content-Type":"multipart/form-data" } ,window.location.href + "kajk/music",formData,function(ev){
                console.log(this.response);
            });
        }
    }
});

}
