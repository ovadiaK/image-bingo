// Get the modal
let activeModal;

// When the user clicks the button, open the modal
function openModal(modalId) {
    show(modalId);
    activeModal = document.getElementById(modalId);
}

// When the user clicks on <span> (x), close the modal
function closeModal(modalId) {
    hide(modalId);
    activeModal = null;
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == activeModal) {
        hide(activeModal.id);

    }
};

function show(id) {
    document.getElementById(id).classList.add("show");
    document.getElementById(id).classList.remove("hide");
}

function hide(id) {
    document.getElementById(id).classList.remove("show");
    document.getElementById(id).classList.add("hide");
}