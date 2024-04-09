import { messageStore, userStatusStore, type ServerMessage } from "$lib/store/websocket-store";




let webSocket: WebSocket;
// let messagess = [];


// Function to establish WebSocket connection
export function connectWebSocket(access_token: string) {
  webSocket = new WebSocket(`ws://localhost:8080/cookie/ws?at=${access_token}`);
  webSocket.onmessage = function (event) {
    console.log('Received message:', event.data);
  };

  webSocket.onopen = () => {
    console.log('WebSocket connection established');
  };

  webSocket.onmessage = (event) => {
    console.log('Message from server ', event.data);
    const eventData = JSON.parse(event.data) as ServerMessage
    // messagess = [...messagess, JSON.parse(event.data)]
    messageStore.update((value) => [...value, JSON.parse(event.data)])
    const data = eventData.data as Array<{ id: string }>
    switch (eventData.type) {
      case "status":
        console.log("data", data)
        console.log("data is boolean", typeof data)
        if (data instanceof Object) {
          for (const val of data) {
            console.log(val.id)
            storeUserStatus(val.id, true)
          }
          return
        }
        if (typeof data == "boolean") {
          storeUserStatus(eventData.source_id, true)
          return
        }

        break;

    }

  };

  webSocket.onerror = (error) => {
    console.error('WebSocket error: ', error);
  };

  webSocket.onclose = () => {
    console.log('WebSocket connection closed');
    // Optionally, implement reconnection logic here
  };
  return webSocket
}

function storeUserStatus(user_id: string, bool: boolean) {
  userStatusStore.update((value) => {
    value[user_id] = bool
    return value
  })
}

export function closeWebSocket() {
  if (webSocket) {
    webSocket.close(1000)
  }
}

// export const getAllMessages = () => { return messagess }


export function sendMessage(message: string) {
  if (webSocket.readyState === WebSocket.OPEN) {
    webSocket.send(message);
  } else {
    console.error('WebSocket is not open.');
  }
}


