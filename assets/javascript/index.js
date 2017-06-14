window.onload = function () {
    var mainContents = new Vue({
        el:'#MainContents',    
        data:{
            musics:[],
            items:[
                { message: 'Foo' },
                { message: 'Bar' }
            ]
        },
        beforeCreate:function(){
            self = this;
            var xhr=new XMLHttpRequest();
            xhr.open("GET","http://localhost:1323/music/new");
            xhr.setRequestHeader( 'Content-Type', 'application/json' );
            xhr.responseType = 'json';
            xhr.onload=function(){
                self.musics = xhr.response;
            }
            xhr.send();
        },
        methods:{
        }
    })
}
