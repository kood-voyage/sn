import { addUserToChat, createChat } from "$lib/client/api/chat-requests"
import { currentUser } from "$lib/client/api/user-requests";
import type { ReturnEntryType } from "$lib/types/requests"
import { v4 as uuidv4 } from 'uuid';


type ChatCreate = ReturnEntryType<"", void>

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
  const userResp = currentUser()
  if (!userResp.ok) {
    console.error(userResp.error)
    return
  }



  // // Add User to the created Chat
  // let addUserResp = await addUserToChat(event, target_user, id)
  // if (!addUserResp.ok) {
  //   console.error(addUserResp.error)
  //   return
  // }
  // if (addUserResp.data <= 200, addUserResp.data >= 299) {
  //   const err = new Error("Adding User To Chat failed with Status Code >>> " + addUserResp.data.toString())
  //   console.error(err)
  //   return
  // }

  // // If the users are the same then just return and don't try to add a second one
  // if (userIDResp.user_id == target_user) {
  //   return redirect(303, "/app/chat")
  // }

  // // Otherwise add the second user
  // addUserResp = await addUserToChat(event, userIDResp.user_id as string, id)
  // if (!addUserResp.ok) {
  //   console.error(addUserResp.error)
  //   return fail(400, { ok: false })
  // }
  // if (addUserResp.data <= 200, addUserResp.data >= 299) {
  //   const err = new Error("Adding User To Chat failed with Status Code >>> " + addUserResp.data.toString())
  //   console.error(err)
  //   return fail(addUserResp.data, { ok: false })
  // }


  // return redirect(303, "/app/chat")
}


