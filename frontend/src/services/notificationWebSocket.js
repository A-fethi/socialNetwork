let notificationSocket = null;
let globalNotificationHandler = null;
let reconnectAttempts = 0;
const maxReconnectAttempts = 5;
let reconnectDelay = 1000;
let isConnecting = false;

function connectNotificationWebSocket() {
  if (isConnecting) {
    console.log("Notification WebSocket connection already in progress");
    return;
  }

  if (
    notificationSocket &&
    (notificationSocket.readyState === WebSocket.CONNECTING ||
      notificationSocket.readyState === WebSocket.OPEN)
  ) {
    console.log("Notification WebSocket already connected or connecting");
    return;
  }

  isConnecting = true;

  try {
    notificationSocket = new WebSocket("ws://localhost:8080/ws/notifications");

    notificationSocket.onopen = function (event) {
      console.log("Connected to notification WebSocket");
      isConnecting = false;
      reconnectAttempts = 0;
      reconnectDelay = 1000;
    };

    notificationSocket.onmessage = function (event) {
      try {
        const data = JSON.parse(event.data);
        console.log("Received notification:", data);

        if (data.type === "new_notification" && globalNotificationHandler) {
          try {
            globalNotificationHandler(data.notification);
          } catch (error) {
            console.error("Error in global notification handler:", error);
          }
        }
      } catch (error) {
        console.error("Error parsing notification WebSocket message:", error);
      }
    };

    notificationSocket.onclose = function (event) {
      console.log("Notification WebSocket connection closed");
      notificationSocket = null;
      isConnecting = false;

      if (
        event.code !== 1000 &&
        reconnectAttempts < maxReconnectAttempts &&
        globalNotificationHandler
      ) {
        reconnectAttempts++;
        console.log(
          `Attempting to reconnect notification WebSocket (${reconnectAttempts}/${maxReconnectAttempts}) in ${reconnectDelay}ms`
        );

        setTimeout(() => {
          connectNotificationWebSocket();
        }, reconnectDelay);

        reconnectDelay = Math.min(reconnectDelay * 2, 30000);
      }
    };

    notificationSocket.onerror = function (error) {
      console.error("Notification WebSocket error:", error);
      isConnecting = false;
    };
  } catch (error) {
    console.error("Failed to create notification WebSocket connection:", error);
    isConnecting = false;
  }
}

function disconnectNotificationWebSocket() {
  if (notificationSocket) {
    notificationSocket.close(1000, "Manual disconnect");
    notificationSocket = null;
  }
  globalNotificationHandler = null;
  isConnecting = false;
}

function setGlobalNotificationHandler(handler) {
  if (typeof handler !== "function") {
    console.error("Notification handler must be a function");
    return;
  }

  console.log("Setting global notification handler");
  globalNotificationHandler = handler;

  if (
    !notificationSocket ||
    notificationSocket.readyState === WebSocket.CLOSED
  ) {
    connectNotificationWebSocket();
  }
}

function removeGlobalNotificationHandler() {
  console.log("Removing global notification handler");
  globalNotificationHandler = null;

  if (!globalNotificationHandler) {
    disconnectNotificationWebSocket();
  }
}

function isConnected() {
  return notificationSocket && notificationSocket.readyState === WebSocket.OPEN;
}

let activeComponents = new Set();
let latestComponent = null;

function registerComponent(componentName) {
  activeComponents.add(componentName);
  latestComponent = componentName;
  console.log(
    `Registered component: ${componentName}, Active components:`,
    Array.from(activeComponents)
  );
}

function unregisterComponent(componentName) {
  activeComponents.delete(componentName);
  if (latestComponent === componentName) {
    latestComponent =
      activeComponents.size > 0
        ? Array.from(activeComponents)[activeComponents.size - 1]
        : null;
  }
  console.log(
    `Unregistered component: ${componentName}, Active components:`,
    Array.from(activeComponents)
  );
}

let componentHandlers = new Map();

function onNotification(componentName, handler) {
  if (typeof handler !== "function") {
    console.error("Notification handler must be a function");
    return;
  }

  registerComponent(componentName);
  componentHandlers.set(componentName, handler);

  if (!globalNotificationHandler) {
    setGlobalNotificationHandler((notification) => {
      if (latestComponent && componentHandlers.has(latestComponent)) {
        const handler = componentHandlers.get(latestComponent);
        try {
          handler(notification);
        } catch (error) {
          console.error(
            `Error in notification handler for ${latestComponent}:`,
            error
          );
        }
      }
    });
  }
}

function removeNotificationHandler(componentName) {
  unregisterComponent(componentName);
  componentHandlers.delete(componentName);

  if (activeComponents.size === 0) {
    removeGlobalNotificationHandler();
  }
}

export const notificationWebSocket = {
  connect: connectNotificationWebSocket,
  disconnect: disconnectNotificationWebSocket,
  onNotification,
  removeNotificationHandler,
  isConnected,
};
