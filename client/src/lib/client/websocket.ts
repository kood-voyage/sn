import { messageStore, notificationStore, userStatusStore, webSocketStore, type CustomNotification, type ServerMessage } from "$lib/store/websocket-store";
import { isChatLine, isNotification } from "$lib/types/type-checks";
import { toast } from "svelte-sonner";
import { type ChatLine } from "./api/chat-requests";



let webSocket: WebSocket

export const ssr = false
export async function connectWebSocket() {

  webSocket = new WebSocket(`ws://localhost:8080/auth/ws`);
  if (webSocket)
    webSocketStore.update((obj) => {
      obj.websocket = webSocket
      return { ...obj }
    })
  // console.log("THIS IS WEBSOCKET >>>", webSocket)

  webSocket.onmessage = function (event) {
    // console.log('Received message:', event.data);
  };

  webSocket.onopen = () => {
    // console.log('WebSocket connection established');
  };

  webSocket.onmessage = (event) => {
    console.log('Message from server ', event.data);
    const eventData = JSON.parse(event.data) as ServerMessage
    const data = eventData.data as Array<{ id: string }>
    // console.log(eventData.type)
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
        // console.log("data of websocket ??>>>", data)
        if (typeof data == "boolean") storeUserStatus(eventData.source_id, data)
        if (typeof data == "number") {
          if (data == 1) {
            storeUserStatus(eventData.source_id, true)
          } else if (data == 0) {
            storeUserStatus(eventData.source_id, false)
          } else if (data == 2) {
            sendMessage('status', 'direct', eventData.source_id, eventData.id, 1);
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
      case "notification":
        // console.log(eventData.data)
        if (isNotification(eventData.data))
          notificationStore.update((old) => {
            old.push(eventData.data as CustomNotification)
            sendAlert(eventData.data as CustomNotification)
            return old
          })
        break;

    }

  };

  webSocket.onerror = (error) => {
    // console.error('WebSocket error: ', error);

    // webSocketStore.set({ websocket: undefined, access_token: undefined })
    // if (access_token == undefined) {
    //   redirect(303, window.location.href)
    // }

  };

  webSocket.onclose = () => {
    // console.log('WebSocket connection closed');

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


export function sendMessage(type: string,
  address: string,
  id: string,
  source_id: string,
  data: number | ChatLine | CustomNotification) {

  const message = JSON.stringify({ type, address, id, source_id, data })
  if (webSocket.readyState === WebSocket.OPEN) {
    webSocket.send(message);
  } else {
    console.error('WebSocket is not open.');
  }
}


export function sendNotification(target_id: string, source_id: string, notification: CustomNotification) {

  console.log(notification)
  sendMessage("notification", "direct", target_id, source_id, notification)
}


function sendAlert(notification: CustomNotification) {
  toast.success(notification.source_information.username)
}