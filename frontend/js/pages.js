
import { Register, Login } from '/frontend/js/app.js'

function addListers() {
    document.body.addEventListener("click", function (event) {
        if (event.target.matches(".register-btn")) {
            navigateTo("register");
        }
        else if (event.target.matches("#register-submit")) {
            Register();
        }
        else if (event.target.matches("#link-login")) {
            navigateTo("login");
        } else if (event.target.matches(".login-btn")) {
            Login();
        }
    });
}

export function navigateTo(page) {
    let content = ""


    if (page === "login") {

        content = `

    <div id="loginform">
        <div class="container">
            <h2>Login</h2>
            <input type="text" id="user" placeholder="Username or Email" required>
            <input type="password" id="password" placeholder="Password" required>
            <button class="login-btn">Login</button>
            <button class="register-btn">register</button>
        </div>
    </div>
       
        `;
    }
    else if (page === "register") {
        content = `
    <div class="container">
        <h2>Register</h2>
        <div class="form-group">
            <label for="nickname">Nickname</label>
            <input type="text" id="nickname" required>
        </div>
        <div class="form-group">
            <label for="age">Age</label>
            <input type="number" id="age" required>
        </div>
        <div class="form-group">
            <label for="gender">Gender</label>
            <select id="gender">
                <option value="male">Male</option>
                <option value="female">Female</option>
            </select>
        </div>
        <div class="form-group">
            <label for="firstName">First Name</label>
            <input type="text" id="firstName" required>
        </div>
        <div class="form-group">
            <label for="lastName">Last Name</label>
            <input type="text" id="lastName" required>
        </div>
        <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" required>
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" required>
        </div>
        <button  id="register-submit">Register</button>
        <p>Already have an account ? <a id="link-login">Login</a></p>
    </div>
        `;
    } else if (page === "home") {
        content = `
    <div class="container">
        <!-- Left Sidebar: Categories -->
        <aside class="sidebar">
            <h3>Categories</h3>
            <ul>
                <li onclick="filterPosts('all')">All</li>
                <li onclick="filterPosts('tech')">Tech</li>
                <li onclick="filterPosts('gaming')">Gaming</li>
                <li onclick="filterPosts('sports')">Sports</li>
            </ul>
        </aside>

        <!-- Main Content: Create Posts & Feed -->
        <main class="main-content">
            <h2>Create a Post</h2>
            <div class="post-form">
                <input type="text" id="post-title" placeholder="Title" required>
                <textarea id="post-content" placeholder="What's on your mind?" required></textarea>
                <select id="post-category">
                    <option value="tech">Tech</option>
                    <option value="gaming">Gaming</option>
                    <option value="sports">Sports</option>
                </select>
                <button onclick="createPost()">Post</button>
            </div>

            <h2>Recent Posts</h2>
            <div id="posts-container">
                <!-- Posts will be dynamically added here -->
            </div>
        </main>

        <!-- Right Sidebar: Online Users -->
        <aside class="users">
            <h3>Online Users</h3>
            <ul id="users-list">
                <li onclick="openChat('John')">John</li>
                <li onclick="openChat('Emma')">Emma</li>
                <li onclick="openChat('Mike')">Mike</li>
            </ul>
        </aside>
    </div>

    <!-- Floating Chat Windows -->
    <div id="chat-container"></div>

        `
    }

    const stylo = document.getElementById('page-style');
    if (stylo) {
        stylo.href = `/frontend/css/${page}.css`
    }

    const app = document.getElementById('app');
    if (app) {
        app.innerHTML = content;
    }

    //  addListers()
}


addListers()