const apiURL1 = "http://127.0.0.1:8080/admin/delpub";

function DeletePublication(){
	let id = document.getElementById("delpubid").value;
	let sub = true;
    if (id.length < 1){
        let idlabel = document.getElementById("delpubid");
        idlabel.placeholder = "Нужно ввести id публикации";
        sub = false;
    }else {
        let idlable = document.getElementById("delpubid");
        idlable.placeholder = "ID для удаления";
    }
    try{
        id = Number(id)
    } catch(err) {
        document.getElementById("delpubid").placeholder = "Нужно ввести число"
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
            id: id,
        }),
    };
    fetch(apiURL1, requestOptions)
        .then(response => {
            if (!response.ok) {
                btn = document.getElementById("btndel");
                btn.value = "Ошибка при при удалении публикации";
            } else {
                console.log("All ok");;
                btn = document.getElementById("btndel")
                btn.value = "Публикация удалена";
            }
        })
    return false;
}