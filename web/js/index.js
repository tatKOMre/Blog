function checkCookie() {
    let cookie = document.cookie
    if (cookie === null || cookie === "" || cookie === undefined) {
        return false;
    } else {
        return true;
    }
}
function hideRegisterButton() {
    const registerButton = document.getElementById('regb');
    const profileButton = document.getElementById('profile')
    
    if (checkCookie()) {
        registerButton.classList.add('hidden'); // Скрыть кнопку
        profileButton.classList.remove('hidden') // показать кнопку
    }
}

// Вызов функции
document.addEventListener('DOMContentLoaded', hideRegisterButton);