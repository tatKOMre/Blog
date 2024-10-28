function checkCookie() {
    let cookie = document.cookie
    if (cookie === null || cookie === "" || cookie === undefined) {
        return false;
    } else {
        return true;
    }
}
console.log(checkCookie())
function hideRegisterButton() {
    const registerButton = document.getElementById('regb');
    
    if (checkCookie()) {
        registerButton.classList.add('hidden'); // Скрыть кнопку
        registerButton.classList.remove('hidden') // показать кнопку
    }
}

// Вызов функции
document.addEventListener('DOMContentLoaded', hideRegisterButton);