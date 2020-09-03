let j = 0;

function nextTurn() {
    if (j === order.length) {
        document.getElementById("next").onclick = window.location.href = '/';
    }
    document.getElementById("main-image").src = order[j].path;
    document.getElementById("main-image").alt = order[j].title;
    document.getElementById("main-image").title = order[j].title;
    document.getElementById("current-image-title").innerText = order[j].title;

    document.getElementById(j + 1).src = order[j].path;
    document.getElementById(j + 1).alt = order[j].title;
    document.getElementById(j + 1).title = order[j].title;

    document.getElementById("current-turn").innerText = j + 1;


    j++;
    if (j === order.length) {
        document.getElementById("next").innerText = "Restart";
    }
}

