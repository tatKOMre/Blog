const apiURL = "http://127.0.0.1:8080/login/";

function loginUser(){
	let login = document.getElementById("login").value;
	let password = document.getElementById("password").value;
	let sub = true;

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
	if (!sub){
		return sub;
	}

	var requestOptions = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            login: login,
            password: password
        }),
    };
	fetch(apiURL, requestOptions)
		.then(response => {
			if (!response.ok) {
				btn = document.getElementById("sub-text");
				btn.innerHTML = "Ошибка при авторизации";
			}
			return response.json()
		})
		.then(data => {
			if (data &&data.token){
				let tkn = data.token
				var d =  new Date();
				d.setTime(d.getTime() + (30*24*60*60*1000));
				var expires = "expires=" + d.toUTCString();
				document.cookie = "token" + "=" + tkn + "; expires=" + expires + ";path=/";
				window.location.replace("/")
			} else {
				console.log("Токен не найден")
			}
		})
		return false;
}

