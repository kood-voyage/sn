import { db } from ".";
import type { ReturnType } from "$lib/types/requests";
import { getUser } from "./profile";
import type { UserModel } from "$lib/types/user";


type ApiUsers = {
  id: string;
  member_type: number;
}

export function getUsersFromArray(users: Array<ApiUsers>) {
  const arr_out: Array<UserModel> = []
  for (const user of users) {
    const userResp = getUser(user.id)
    if (!userResp.ok) {
      continue
    }

    arr_out.push(userResp.data)

  }
  return arr_out
}

export type UserRowType = {
  id: string,
  username: string,
  first_name: string,
  last_name: string,
  avatar: string,
  cover: string,
  description: string
}

export type userResp = ReturnType<UserRowType[]>

export function mainGetAllUsers(): userResp {


  const query = `SELECT id,
    username,
    first_name,
    last_name,avatar,cover,description FROM user`
  try {
    const prep = db.prepare(query)

    const row = prep.all() as UserRowType[]

    if (typeof row === 'object' && row !== null && row.length != 0) {

      return { ok: true, data: row }
    } else {
      throw new Error("Error querying get all users!")
    }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }


}