const apiURL = "http://127.0.0.1:8001/admin/crpub";

function CreatePublication(){
	let title = document.getElementById("title").value;
	let text = document.getElementById("text").value;
	let sub = true;
    if (title.lenght < 1){
        let titlelabel = document.getElementById("titlelab");
        titlelabel.innerHTML = "Название слишком короткое";
        sub = false;
    }else {
        let titlelabel = document.getElementById("titlelab");
        titlelabel.innerHTML = "Название";
    }
    if (text.lenght < 1){
        let textlable = document.getElementById("textlab");
        textlable.innerHTML = "Содержание публикации должно быть больше";
        sub = false;
    }else {
        let textlable = document.getElementById("textlab");
        textlable.innerHTML = "Текст публикации";
    }
    if (!sub){
        return sub
    }
    
    var requestOptions = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            title: titlelable,
            text: textlable,
        }),
    };
    fetch(apiURL, requestOptions)
        .then(response => {
            if (!response.ok) {
                btn = document.getElementById("titletext");
                btn.innerHTML = "Ошибка при создании публикации";
            } else {
                console.log("All ok");
                window.location.replace(getQueryVariable("to"));
            }
        })
    return false;
}