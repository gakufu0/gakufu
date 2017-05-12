function test(){
    var xhr=new XMLHttpRequest();
    xhr.open("POST","http://localhost:1323/aaaa123/music");
    xhr.setRequestHeader('Content-Type', 'application/json' );
    var data = {
        music_id : "test2",
        music_name : "testttaaa",
        content : "faillllf"
    } 
    xhr.onload=function(){console.log(xhr.response);}
    xhr.send(JSON.stringify(data));
}

function test2(){
    var xhr=new XMLHttpRequest();
    xhr.open("GET","http://localhost:1323/music/new");
    xhr.setRequestHeader( 'Content-Type', 'application/json' );
    xhr.onload=function(){console.log(xhr.response);}
    xhr.send();
}

