let arr_value = [];
let arr_material = [];
let a=1;
function final_values() {
    let arr_final = [];
    for (let i=1;i<=a;i+=1){
        if (document.getElementById(("select_materials"+i))) {
            if(document.getElementById(("input_materials" + i))) {
                arr_final.push(document.getElementById(("select_materials" + i)).value, document.getElementById(("input_materials" + i)).value);
            }
        }
    }
    document.getElementById(("input_materials"+a)).value=arr_final;
}

function addInput() {
    // arr.push(document.getElementById(("select_materials"+a)).value,document.getElementById(("input_materials"+a)).value)
    a+=1
    var container = document.getElementById("container");
    var newDiv = document.createElement("div");
    newDiv.id = "list_materials";
    newDiv.className = "input-group mb-3";

    var select = document.createElement("select");
    select.className = "btn btn-outline-secondary dropdown-toggle";
    select.name = "whatMaterial1";
    select.id=("select_materials"+a);

    var option1 = document.createElement("option");
    option1.value = "1";
    option1.innerHTML = "Песок карьерный";
    select.appendChild(option1);

    var option2 = document.createElement("option");

    option2.value = "2";
    option2.innerHTML = "Песок строительный";
    select.appendChild(option2);

    var option3 = document.createElement("option");
    option3.value = "3";
    option3.innerHTML = "Гравий";
    select.appendChild(option3);
    var option4 = document.createElement("option");
    option4.value = "4";
    option4.innerHTML = "ПГС (песчано-гравийная смесь)";
    select.appendChild(option4);
    var option5 = document.createElement("option");
    option5.value = "5";
    option5.innerHTML = "Щебень природный";
    select.appendChild(option5);
    var option6 = document.createElement("option");
    option6.value = "6";
    option6.innerHTML = "Пескогрунт";
    select.appendChild(option6);
    var option7 = document.createElement("option");
    option7.value = "7";
    option7.innerHTML = "Грунт";
    select.appendChild(option7);
    var option8 = document.createElement("option");
    option8.value = "8";
    option8.innerHTML = "Супесь";
    select.appendChild(option8);
    var option9 = document.createElement("option");
    option9.value = "9";
    option9.innerHTML = "Суглинок";
    select.appendChild(option9);
    var option10 = document.createElement("option");
    option10.value = "10";
    option10.innerHTML = "Глина";
    select.appendChild(option10);
    var option11 = document.createElement("option");
    option11.value = "11";
    option11.innerHTML = "Булыжник";
    select.appendChild(option11);
    var option12 = document.createElement("option");
    option12.value = "12";
    option12.innerHTML = "Валун для ландшафта";
    select.appendChild(option12);
    var option13 = document.createElement("option");
    option13.value = "13";
    option13.innerHTML = "Ландшафтный камень";
    select.appendChild(option13);


    var input = document.createElement("input");
    input.id=("input_materials"+a);
    input.name = "value1";
    input.className = "form-control";
    input.placeholder = "Укажите объем";
    input.setAttribute("aria-label", "Text input with dropdown button");

    var addButton = document.createElement("button");
    addButton.type = "button";
    addButton.className = "btn btn-info";
    addButton.innerHTML = "+";
    addButton.onclick = addInput;

    var removeButton = document.createElement("button");
    removeButton.type = "button";
    removeButton.className = "btn btn-danger";
    removeButton.innerHTML = "Удалить";
    removeButton.onclick = function() {
        container.removeChild(newDiv);
    };

    newDiv.appendChild(select);
    newDiv.appendChild(input);
    newDiv.appendChild(addButton);
    newDiv.appendChild(removeButton);
    container.appendChild(newDiv);
}