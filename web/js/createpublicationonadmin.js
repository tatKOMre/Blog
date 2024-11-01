const apiURL = "http://127.0.0.1:8080/admin/crpub";

function CreatePublication(){
	let title = document.getElementById("crpub").value;
	let text = document.getElementById("text").value;
	let sub = true;
    if (title.length < 3){
        let titlelabel = document.getElementById("crpubtitle");
        titlelabel.placeholder = "Название слишком короткое";
        sub = false;
    }else {
        let titlelabel = document.getElementById("crpubtitle");
        titlelabel.placeholder = "Заголовок1";
    }
    if (text.length < 10){
        let textlable = document.getElementById("text");
        textlable.placeholder = "Содержание публикации должно быть больше";
        sub = false;
    }else {
        let textlable = document.getElementById("text");
        textlable.placeholder = "Текст публикации1";
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
            title: title,
            text: text,
        }),
    };
    fetch(apiURL, requestOptions)
        .then(response => {
            if (!response.ok) {
                btn = document.getElementById("btn");
                btn.value = "Ошибка при создании публикации";
            } else {
                console.log("All ok");;
                btn = document.getElementById("btn")
                btn.value = "Публикация создана";
            }
        })
    return false;
}