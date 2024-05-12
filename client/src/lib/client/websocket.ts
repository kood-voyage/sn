import { messageStore, userStatusStore, webSocketStore, type ServerMessage } from "$lib/store/websocket-store";
import { isChatLine, type ChatLine } from "./api/chat-requests";



let webSocket: WebSocket

export const ssr = false
export async function connectWebSocket() {

  webSocket = new WebSocket(`ws://localhost:8080/auth/ws`);
  if (webSocket)
    webSocketStore.update((obj) => {
      obj.websocket = webSocket
      return { ...obj }
    })
  console.log("THIS IS WEBSOCKET >>>", webSocket)

  webSocket.onmessage = function (event) {
    console.log('Received message:', event.data);
  };

  webSocket.onopen = () => {
    console.log('WebSocket connection established');
  };

  webSocket.onmessage = (event) => {
    console.log('Message from server ', event.data);
    const eventData = JSON.parse(event.data) as ServerMessage
    const data = eventData.data as Array<{ id: string }>
    console.log(eventData.type)
    switch (eventData.type) {
      case "status":
        // console.log("data", data)
        // console.log("data is boolean", typeof data)
        if (data instanceof Object) {
          for (const val of data) {
            console.log(val.id)
            storeUserStatus(val.id, true)
          }
          return
        }
        console.log("data of websocket ??>>>", data)
        if (typeof data == "boolean") storeUserStatus(eventData.source_id, data)
        if (typeof data == "number") {
          if (data == 1) {
            storeUserStatus(eventData.source_id, true)
          } else if (data == 0) {
            storeUserStatus(eventData.source_id, false)
          } else if (data == 2) {
            sendMessage(
              JSON.stringify({
                type: 'status',
                address: 'direct',
                id: eventData.source_id,
                source_id: eventData.id,
                data: 1
              })
            );
            storeUserStatus(eventData.source_id, true)
          }
          return
        }

        break;
      case "message":
        if (isChatLine(eventData.data)) messageStore.update((old) => {
          old.push(eventData.data as ChatLine)
          return old
        })
        break;

    }

  };

  webSocket.onerror = (error) => {
    console.error('WebSocket error: ', error);

    // webSocketStore.set({ websocket: undefined, access_token: undefined })
    // if (access_token == undefined) {
    //   redirect(303, window.location.href)
    // }

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
    console.log("USER WENT BOOL >>>", bool)
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


