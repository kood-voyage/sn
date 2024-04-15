import { db } from "."
import type { RequestEvent } from "@sveltejs/kit"
import { getUserFollowers, getUserFollowing } from "../api/user-requests";
import { getUserPosts } from "../api/post-requests";
import type { ReturnType } from "$lib/types/requests";
import type { UserModel } from "$lib/types/user";

interface Response {
  ok: boolean;
  data: unknown;
  error: Error;
  message: string;
}



export async function getProfile(event: RequestEvent, username: string) {
  const userResp = getUser(username) as Response
  if (!userResp.ok) {
    console.error("'profile.ts' >>>", userResp.message)
    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user = await userResp.data as UserModel


  const followingResp = await getUserFollowing(event, user.id)

  const posts = await getUserPosts(event, user.id)




  if (!posts.ok) {
    return { ok: posts.ok, error: posts.error, message: posts.message }
  }




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
    followers: followersResp.data,
    posts: posts.data
  }
}

type GetUser = ReturnType<UserModel>

export function getUser(unique_credentials: string): GetUser {


  const query = `SELECT id,
    username,
    email,
    date_of_birth,
    first_name,
    last_name,
    avatar,
    cover,
    description
    FROM user WHERE username = ? OR id = ?`

  try {
    const user = db.prepare(query).get(unique_credentials, unique_credentials) as UserModel

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



export function setAvatar(link: string, user_id: string) {


  const path = "https://profilemediabucket-voyage.s3.amazonaws.com/" + link

  const query = `UPDATE user SET avatar = ? WHERE id = ?`
  try {
    db.prepare(query).run(path, user_id)

    return { ok: true }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }

}


export function setCover(link: string, user_id: string) {

  const path = "https://profilemediabucket-voyage.s3.amazonaws.com/" + link

  const query = `UPDATE user SET cover = ? WHERE id = ?`
  try {
    db.prepare(query).run(path, user_id)

    return { ok: true }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }

}


export function setDescription(link: string, user_id: string) {

  const query = `UPDATE user SET description = ? WHERE id = ?`
  try {
    db.prepare(query).run(link, user_id)

    return { ok: true }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }

}