document.getElementById("getallbutton").onclick = getAllmember
function getAllmember() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("allmember").innerHTML = this.responseText;
        }
    };
   
    xhttp.open("GET", "http://localhost:8080/members", true);
    xhttp.send();
}





document.getElementById("getbutton").onclick = getMember

function getMember() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("response").innerHTML = this.responseText;
        }
    };
    var input= document.getElementById("id").value;
    xhttp.open("GET", "http://localhost:8080/member/"+input, true);
    xhttp.send();
}



document.getElementById("deletebutton").onclick = DeleteMember

function DeleteMember() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById("response").innerHTML = this.responseText;
        }
    };
    var input= document.getElementById("delete").value;
    xhttp.open("DELETE", "http://localhost:8080/member/delete/"+input, true);
    xhttp.send();
}


