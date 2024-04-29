import { currentUser, getUserFollowers, getUserFollowing } from "$lib/client/api/user-requests";
import { fail, type LoadEvent } from "@sveltejs/kit";
import type { LayoutLoad } from "./$types";




export const load: LayoutLoad = async ({ fetch }: LoadEvent) => {

  // console.log(event)

  const userResp = await currentUser(fetch)
  // console.log(user)
  if (!userResp.ok) {
    console.error("User not found please try logging in again or call customer support!")
    fail(400)
    // alert(("User not found please try logging in again or call customer support"))
    return { data: [], followers: [], following: [] }
  }

  const data = userResp.data
  const user_id = data.id as string

  const followers = await getUserFollowers(user_id, fetch)
  // console.log(followers)
  if (!followers.ok) {
    console.error("Error getting who Follows the user!")
    fail(400)
    // alert(("User not found please try logging in again or call customer support"))
    return { data: userResp.data, followers: [], following: [] }
  }

  const following = await getUserFollowing(user_id, fetch)
  // console.log(following)
  if (!following.ok) {
    console.error("Error getting who the user is Following!")
    fail(400)
    // alert(("User not found please try logging in again or call customer support"))
    return { data: userResp.data, followers: followers.data, following: [] }
  }


  return { data: userResp.data, followers: followers.data, following: following.data }

}
