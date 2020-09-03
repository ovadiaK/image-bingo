// Restricts input for the given textbox to the given inputFilter.
function setInputFilter(textbox, inputFilter) {
    ["input", "keydown", "keyup", "mousedown", "mouseup", "select", "contextmenu", "drop"].forEach(function(event) {
        textbox.addEventListener(event, function() {
            if (inputFilter(this.value)) {
                this.oldValue = this.value;
                this.oldSelectionStart = this.selectionStart;
                this.oldSelectionEnd = this.selectionEnd;
            } else if (this.hasOwnProperty("oldValue")) {
                this.value = this.oldValue;
                this.setSelectionRange(this.oldSelectionStart, this.oldSelectionEnd);
            }
        });
    });
}

setInputFilter(document.getElementById("cards"), function(value) {
    return /^\d*$/.test(value) && (value === "" || parseInt(value) <= 100); });

// setInputFilter(document.getElementById("quantity"), function(value) {
//     return /^\d*$/.test(value) && (value === "" || parseInt(value) <= 100); });
//
// function numberLimit(id, limit) {
//     setInputFilter(document.getElementById(id), function(value) {
//         return /^\d*$/.test(value) && (value === "" || parseInt(value) <= limit); });
// }