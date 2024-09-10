function loginUser(){
	let login = document.getElementById("login").value;
	let password = document.getElementById("password").value;
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
	return sub;
}
