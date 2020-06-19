function login() {
    window.location.href = "login";
}

function signup() {
    window.location.href = "signup";
}

function validateForm(thisform) {
    xmlhttp = new XMLHttpRequest
    xmlhttp.open("GET", "/finduser?username=" + thisform.username.value, false)
    xmlhttp.send()
    if (xmlhttp.responseText == "false") {
        alert("用户名已被占用");
        return false;
    }
    if (thisform.password.value != thisform.password2.value) {
        alert("两次输入的密码不一致");
        return false;
    }
    return true;
}

function ask() {
    window.location.href = "ask";
}