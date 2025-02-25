
export function showError(field, message) {
    let errorElement = document.getElementById(`${field}-error`);
    if (!errorElement) {
        errorElement = document.createElement("p");
        errorElement.id = `${field}-error`;
        errorElement.className = "error-message";
        document.getElementById(field).parentNode.appendChild(errorElement);
    }
    errorElement.innerText = message;
}

export function clearErrors() {
    document.querySelectorAll(".error-message").forEach(el => el.innerText = "");
}
export function showToast(message, type = "error") {
    const toast = document.getElementById("toast");
    toast.innerText = message;
    toast.className = "toast " + type;
    toast.style.display = "block";
    setTimeout(() => { toast.style.display = "none"; }, 3000);
}