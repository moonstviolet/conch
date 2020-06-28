function validateForm(thisform) {
    xmlhttp = new XMLHttpRequest
    xmlhttp.open("GET", "/user/find?username=" + thisform.username.value, false)
    xmlhttp.send()
    var text = xmlhttp.responseText;
    var obj = JSON.parse(text)
    if (obj["isValid"] == false) {
        alert("用户名已被占用");
        return false;
    }
    if (thisform.password.value != thisform.password2.value) {
        alert("两次输入的密码不一致");
        return false;
    }
    return true;
}

function newQuestion() {
    //TODO 登录提示
    window.location.href = "/question/new";
}

function followQuestion() {
    var obj = document.getElementById("followQuestion");
    if (obj.value == "1") {
        obj.value = "2";
        obj.innerHTML = "已关注";
        obj.classList.remove("orangered-button");
        obj.classList.add("grey-button");
    } else {
        obj.value = "1";
        obj.innerHTML = "关注问题";
        obj.classList.add("orangered-button");
        obj.classList.remove("grey-button");
    }
}