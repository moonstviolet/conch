function login() {
    window.location.href="login";
}

function signup() {
    window.location.href="signup";
}

function validateForm(thisform) {
    if (thisform.password.value != thisform.password2.value) {
        alert("两次输入的密码不一致");
        return false;
    }
    return true;
}