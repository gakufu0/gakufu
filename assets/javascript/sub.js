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
