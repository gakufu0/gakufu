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
        buildXHR("GET","json", { "Content-Type":"application/json" } ,window.location.href + "/music/new",null, function(ev){
            self.musics = this.response;
        });
    },
    methods:{
        getImagePath(image){
            return "music/picture" + image;
        },
    }
});

}
