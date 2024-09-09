const apiURL = "http://127.0.0.1:8001/register/";

function getQueryVariable(variable) {
    var query = window.location.search.substring(1);
    var vars = query.split("&");
    for (var i=0; i<vars.length; i++) {
        var pair = vars[i].split("=");
        if (pair[0] == variable) {
            return pair[1];
        }
    }
    return true;
}

function registerUser(){
    let login = document.getElementById("login").value;
    let password = document.getElementById("password").value;
    let anpassword = document.getElementById("anpassword").value;
    let sub = true

    if (login.length <  4) {
        let loginLabel = document.getElementById("ll");
        loginLabel.innerHTML = "Логин (слишком короткий)";
        sub = false;
    
    } else {
        let loginLabel = document.getElementById("ll");
        loginLabel.innerHTML = "Логин";
    } 

    if (password.length < 8) {
        let passwordLabel = document.getElementById("pl");
        passwordLabel.innerHTML = "Пароль (слишком короткий)";
        sub = false;
    } else {
        let passwordLabel = document.getElementById("pl");
        passwordLabel.innerHTML = "Пароль";
    } 
    if (password != anpassword){
        let passwordLabel = document.getElementById("plr");
        passwordLabel.innerHTML = "Пароли должны совпадать";
        sub = false;
    } else {
        let passwordLabel = document.getElementById("plr");
        passwordLabel.innerHTML = "Повторите пароль";
        
    }
    if (!sub) {
        return sub;
    }

    var requestOptions = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            login: login,
            password: password,
        }),
    };

    fetch(apiURL, requestOptions)
        .then(response => {
            if (!response.ok) {
                btn = document.getElementById("sub-text");
                btn.innerHTML = "Ошибка при входе";
            } else {
                console.log("All ok");
                tkn = response.json.token;
                console.log(tkn);
                var d = new Date();
                d.setTime(d.getTime() + (30*24*60*60*1000));
                var expires = "expires="+ d.toUTCString();
                document.cookie = "token" + "=" + tkn + "; expires=" + expires + ";path=/";

                window.location.replace(getQueryVariable("to"));
            }
        })
    return false;
}