import { db } from "."

export function getUser(username: string) {
  type RowType = {
    id: string,
    username: string,
    email: string,
    timestamp: string,
    date_of_birth: string,
    first_name: string,
    last_name: string
  }

  const query = `SELECT (id,
    username,
    email,
    timestamp,
    date_of_birth,
    first_name,
    last_name) FROM user WHERE username = ? LIMIT 1`

  try {
    const user = db.prepare(query).get(username) as RowType

    if (typeof user === 'object' && user !== null && user.id) {

      return { ok: true, user: user }
    } else {
      throw new Error("Session Not Found")
    }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }

}

