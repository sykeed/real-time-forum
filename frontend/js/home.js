const chatWindows = {};

function openChat(username) {
    if (chatWindows[username]) return; // Prevent opening multiple chats

    const chatBox = document.createElement("div");
    chatBox.className = "chat-box";
    chatBox.innerHTML = `
        <div class="chat-header" onclick="closeChat('${username}')">${username} ✖</div>
        <div class="chat-body" id="chat-body-${username}"></div>
        <div class="chat-footer">
            <input type="text" id="chat-input-${username}" placeholder="Type a message...">
            <button onclick="sendMessage('${username}')">Send</button>
        </div>
    `;
    
    document.getElementById("chat-container").appendChild(chatBox);
    chatWindows[username] = chatBox;
}

function closeChat(username) {
    document.getElementById("chat-container").removeChild(chatWindows[username]);
    delete chatWindows[username];
}

function sendMessage(username) {
    const input = document.getElementById(`chat-input-${username}`);
    if (input.value.trim() === "") return;

    const chatBody = document.getElementById(`chat-body-${username}`);
    chatBody.innerHTML += `<p><strong>You:</strong> ${input.value}</p>`;
    input.value = "";
    chatBody.scrollTop = chatBody.scrollHeight; // Auto-scroll
}
