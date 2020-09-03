function printCards() {
    // hide("print");
    // hide("back");
    window.print();
    // show("print");
    // show("back");
}


function show(id) {
    document.getElementById(id).classList.add("show");
    document.getElementById(id).classList.remove("hide");
}

function hide(id) {
    document.getElementById(id).classList.remove("show");
    document.getElementById(id).classList.add("hide");
}