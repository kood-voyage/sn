import { db } from ".";
import type { ReturnType } from "$lib/types/requests";

export type UserRowType = {
  id: string,
  username: string,
  first_name: string,
  last_name: string,
  avatar: string,
  cover: string,
  description: string

}

type userResp = ReturnType<UserRowType[]>

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