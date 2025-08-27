let socket = null;
let messageHandlers = new Map();
let reconnectAttempts = 0;
const maxReconnectAttempts = 5;
let reconnectDelay = 1000;
import { fetchInfo } from "../api/getInfo.js";
function showNotification(message, type) {
  const notification = document.createElement("div");
  notification.className = `notification ${type}`;
  notification.textContent = message;

  notification.style.position = "fixed";
  notification.style.top = "20px";
  notification.style.right = "20px";
  notification.style.padding = "15px 25px";
  notification.style.borderRadius = "4px";
  notification.style.color = "white";
  notification.style.zIndex = "1000";
  notification.style.transition = "opacity 0.3s ease-in-out";

  if (type === "success") {
    notification.style.backgroundColor = "#4CAF50";
  } else if (type === "error") {
    notification.style.backgroundColor = "#f44336";
  } else {
    notification.style.backgroundColor = "#2196F3";
  }

  document.body.appendChild(notification);

  setTimeout(() => {
    notification.style.opacity = "0";
    setTimeout(() => {
      document.body.removeChild(notification);
    }, 300);
  }, 3000);
}

async function connect(groupId) {
  const gg = await fetchInfo();
  console.log("gg", gg);
  const currentUserId = gg.Username;

  if (socket && socket.readyState === WebSocket.OPEN) {
    return;
  }

  const wsUrl = `ws://localhost:8080/ws/group/${groupId}`;
  socket = new WebSocket(wsUrl);

  socket.onopen = () => {
    console.log("Group WebSocket connected");
    showNotification("Connected to group chat", "success");
    reconnectAttempts = 0;
    reconnectDelay = 1000;
  };

  socket.onmessage = (event) => {
    try {
      const data = JSON.parse(event.data);
      if (data.error) {
        console.error("WebSocket error:", data.error);
        showNotification(data.error, "error");
        socket.close();
        return;
      }
      if (Array.isArray(data)) {
        messageHandlers.get("initial")?.forEach((handler) => handler(data));
      } else {
        if (data.username !== currentUserId) {
          showNotification(data.username + " sent a message", "success");
        } else {
          showNotification("I sent a message", "success");
        }

        messageHandlers.get("message")?.forEach((handler) => handler(data));
      }
    } catch (error) {
      console.error("Error parsing WebSocket message:", error);
      showNotification("Error processing message", "error");
    }
  };

  socket.onclose = () => {
    console.log("Group WebSocket disconnected");
    attemptReconnect(groupId);
  };

  socket.onerror = (error) => {
    console.error("Group WebSocket error:", error);
  };
}

function attemptReconnect(groupId) {
  if (reconnectAttempts >= maxReconnectAttempts) {
    console.error("Max reconnection attempts reached");
    return;
  }

  setTimeout(() => {
    reconnectAttempts++;
    reconnectDelay *= 2;
    console.log(
      `Attempting to reconnect (${reconnectAttempts}/${maxReconnectAttempts})...`
    );
    connect(groupId);
  }, reconnectDelay);
}

function disconnect() {
  if (socket) {
    socket.close();
    socket = null;
  }
}

function sendMessage(content) {
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    console.error("WebSocket is not connected");
    return;
  }

  const message = {
    type: "message",
    content: content,
    timestamp: new Date().toISOString(),
  };

  console.log("Sending message:", message);
  socket.send(JSON.stringify(message));
}

function onMessage(handler) {
  if (!messageHandlers.has("message")) {
    messageHandlers.set("message", new Set());
  }
  messageHandlers.get("message").add(handler);
}

function onInitialMessages(handler) {
  if (!messageHandlers.has("initial")) {
    messageHandlers.set("initial", new Set());
  }
  messageHandlers.get("initial").add(handler);
}

function removeMessageHandler(handler) {
  messageHandlers.get("message")?.delete(handler);
}

function removeInitialHandler(handler) {
  messageHandlers.get("initial")?.delete(handler);
}

export const groupWebSocket = {
  connect,
  disconnect,
  sendMessage,
  onMessage,
  onInitialMessages,
  removeMessageHandler,
  removeInitialHandler,
  showNotification,
};
