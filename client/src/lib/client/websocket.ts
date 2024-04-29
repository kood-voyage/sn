import { messageStore, userStatusStore, webSocketStore, type ServerMessage } from "$lib/store/websocket-store";




let webSocket: WebSocket
webSocketStore.subscribe((obj) => {
  if (obj.access_token == undefined) return
  if (obj.websocket != undefined) return
  if (obj.access_token != undefined && obj.websocket == undefined) connectWebSocket()
})

export const ssr = false
export async function connectWebSocket() {

  webSocket = new WebSocket(`ws://localhost:8080/cookie/ws`);
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

    // webSocketStore.set({ websocket: undefined, access_token: undefined })
    // if (access_token == undefined) {
    //   redirect(303, window.location.href)
    // }

  };

  webSocket.onclose = () => {
    webSocketStore.set({ websocket: undefined, access_token: undefined })
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


