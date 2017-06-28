function buildXHR(method, type=null ,header = null, url, sendData = null, onloadFunc){
    var xhr= new XMLHttpRequest();
    xhr.open(method ,url);
    for(k in header)
        xhr.setRequestHeader(k,header[k]);

    if(type != null)
        xhr.responseType= type;
    xhr.setRequestHeader('Pragma', 'no-cache');
    xhr.setRequestHeader('Cache-Control', 'no-cache');
    xhr.send(sendData);
    xhr.onload = onloadFunc;

    return xhr;
}

function loadFile(event){
    var reader = new FileReader();
    reader.onload = function(){
        var loadImage = document.getElementById("loadImage");
        loadImage.src = reader.result;
    }
    reader.readAsDataURL(event.target.files[0]);
}

function postImage(){
    var formData = new FormData();
    var id = document.getElementById("inputId").value;
    var name = document.getElementById("inputName").value;
    var image = document.getElementById("inputImage").files[0];
    formData.append("music_name",name);
    formData.append("music_id",id);
    formData.append("file",image);

    var request = new XMLHttpRequest();
    request.open("POST", "http://192.168.2.104:1323/test/music");
    request.send(formData);
    request.onload=function(){console.log(this.responseText);}
}
