import { navigateTo } from '/frontend/js/pages.js';
import { showError,clearErrors,showToast } from '/frontend/js/error.js';

 
export async function Register() {
    console.log("eeee");
    clearErrors()
        //evnt.preventDefault();

        const obj = {
            nickname: document.getElementById("nickname").value.trim(),
            firstName: document.getElementById("firstName").value.trim(),
            lastName: document.getElementById("lastName").value.trim(),
            age: Number(document.getElementById("age").value), // Ensure number type
            gender: document.getElementById("gender").value.trim(),
            email: document.getElementById("email").value.trim(),
            password: document.getElementById("password").value.trim(),
        };


        console.log("Sending Data:", obj); // Debugging step

        const fields = {
            nickname: { value: obj.nickname, message: "Nickname is required." },
            firstName: { value: obj.firstName, message: "First name is required." },
            lastName: { value: obj.lastName, message: "Last name is required." },
            age: { value: obj.age, condition: obj.age <= 0 || isNaN(obj.age), message: "Please enter a valid age." },
            email: { value: obj.email, condition: !obj.email.includes("@"), message: "Please enter a valid email address." },
            password: { value: obj.password, condition: obj.password.length < 6, message: "Password must be at least 6 characters." }
        };
    
        let hasError = Object.keys(fields).some(field => {
            const { value, condition, message } = fields[field];
            if (!value || condition) {
                showError(field, message);
                return true;
            }
            return false;
        });
    
        if (hasError) return;


        try {
            const response = await fetch("/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(obj),
            });

            const result = await response.json(); // Ensure response is JSON
            console.log("Response:", result);

            if (!response.ok) {
                throw new Error(result.message );
            }

            showToast("Registered successfully!", "success");
        } catch (error) {
            console.error("Error:", error);
            showToast(error.message, "error");
        }
 
}

document.querySelector(".register-btn").addEventListener("click", function () {
    navigateTo("register")
})

//   ////////////   errors handling //////////////////////



/////////////////////////////////////////////////////////
async function Login() {

    document.querySelector(".login-btn").addEventListener("click", async function(event) {
        event.preventDefault()
    
        const obj = {
            user : document.querySelector("user"),
            password : document.querySelector("password")
        }
        try {
            const response = await fetch ("/login" , {
                method : "POST",
                body :  JSON.stringify(obj)
            })
        }catch (eror){
            console.log(eror);  
        }
       
    })
}


