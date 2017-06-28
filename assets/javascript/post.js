window.onload = function () {

var mainContents = new Vue({
    el:'#postImage',
    data:{
    },
    methods:{
        postImage(){
            var formData = new FormData();
            var id = document.getElementById("inputId").value;
            var name = document.getElementById("inputName").value;
            var image = document.getElementById("inputImage").files[0];
            console.log(id);
            console.log(name);
            console.log(image);
        }
    }
});

}
