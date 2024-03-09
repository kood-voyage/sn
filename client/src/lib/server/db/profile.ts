import { db } from "."
import type { RequestEvent } from "@sveltejs/kit"
import { getUserFollowers, getUserFollowing } from "../api/userRequests";

interface Response {
  ok: boolean;
  data: unknown;
  error: Error;
  message: string;
}

type User = {
  id: string,
  username: string,
  email: string,
  timestamp: string,
  date_of_birth: string,
  first_name: string,
  last_name: string
}

export async function mainGetProfile(event: RequestEvent, username: string) {
  const userResp = getUser(username) as Response
  if (!userResp.ok) {
    console.error("'profile.ts' >>>", userResp.message)

    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user = userResp.data as User
  // console.log(user)

  const followingResp = await getUserFollowing(event, user.id)
  if (!followingResp.ok) {
    return { ok: followingResp.ok, error: followingResp.error, message: followingResp.message }
  }

  const followersResp = await getUserFollowers(event, user.id)
  if (!followersResp.ok) {
    return { ok: followersResp.ok, error: followersResp.error, message: followersResp.message }
  }

  // GET THE POSTS
  // GET PROFILE IF THEY HAVE 
  // GET BANNER IMAGE 

  return {
    user: user,
    following: followingResp.data,
    followers: followersResp.data
  }
}

function getUser(username: string) {


  const query = `SELECT id,
    username,
    email,
    timestamp,
    date_of_birth,
    first_name,
    last_name FROM user WHERE username = ?`

  try {
    const user = db.prepare(query).get(username) as User

    if (typeof user === 'object' && user !== null && user.id) {

      return { ok: true, data: user }
    } else {
      throw new Error("User Not Found")
    }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }

}

