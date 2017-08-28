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

function createUser(){
  var userid = document.getElementById("userid").value;
  var accountName = document.getElementById("account_name").value;
  var password = document.getElementById("password").value;
  var repassword = document.getElementById("repassword").value;
  var cautionPassword = document.getElementById("caution_pw").style;
  var cautionDataEnough = document.getElementById("caution_ne").style;
  var cautionUserAlreadyUsed = document.getElementById("caution_au").style;

  if(userid == "" || accountName == "" || password == "" || repassword == ""){
    cautionDataEnough.display = "block";
    return;
  }
  if(password != repassword){
    cautionPassword.display = "block";
    return;
  }else{
    var obj = {
      "user_id":userid,
      "password":password,
      "account_name":accountName
    }

    self.buildXHR("POST","json", { "Content-Type":"application/json" },"http://localhost:1323/createuser",JSON.stringify(obj), function(ev){
      var res=this.response;
      if(res.Code == 200){
        setInterval(function(){location.href = "http://localhost:1323/"+userid;},500);
      }else{
        cautionUserAlreadyUsed.display = "block";
      }
    });
  }
}
