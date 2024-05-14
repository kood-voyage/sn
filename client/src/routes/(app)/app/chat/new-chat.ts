import { addUserToChat, createChat } from "$lib/client/api/chat-requests"
import { currentUser } from "$lib/client/api/user-requests";
import type { ReturnEntryType } from "$lib/types/requests"
import { v4 as uuidv4 } from 'uuid';


type ChatCreate = ReturnEntryType<"chatCreated", boolean>

export const newChatCreate = async (formData: FormData): Promise<ChatCreate> => {


  // Create and Get variables like chat_id and user_id, respectively.
  const id = uuidv4()

  // Create Chat 
  const createResp = await createChat(id)
  if (!createResp.ok) {
    console.error(createResp.error)
    alert("Something went wrong whilst creating chat!")
    return { ok: false, error: createResp.error, message: createResp.message }
  }
  if (createResp.status <= 200, createResp.status >= 299) {
    const err = new Error("Creating chat failed with Status Code >>> " + createResp.status.toString())
    console.error(err)
    alert(err)
    return { ok: false, error: err, message: err.message }
  }

  // Get the target User to add to chat
  const target_user = formData.get("target") as string
  if (typeof target_user != "string" || target_user == undefined) {
    const err = new Error("User ID not found on form")
    console.error(err)
    return { ok: false, error: err, message: err.message }
  }

  // Get the user creating the post 
  const userResp = await currentUser()
  if (!userResp.ok) {
    console.error(userResp.error)
    return { ok: false, error: userResp.error, message: userResp.message }
  }

  console.log("CUREENT USER ID >>>", userResp.data.id)
  console.log("TARGET USER ID >>>", target_user)
  if (userResp.data.id == target_user) {
    return { ok: true, chatCreated: true }
  }

  // Add User to the created Chat
  const addUserResp = await addUserToChat(target_user, id)
  if (!addUserResp.ok) {
    console.error(addUserResp.error)
    return { ok: false, error: addUserResp.error, message: addUserResp.message }
  }
  if (addUserResp.status <= 200, addUserResp.status >= 299) {
    const err = new Error("Adding User To Chat failed with Status Code >>> " + addUserResp.status.toString())
    console.error(err)
    return { ok: false, error: err, message: err.message }
  }

  return { ok: true, chatCreated: true }
}


