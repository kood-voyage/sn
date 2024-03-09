import Database from 'better-sqlite3';
import { readFileSync } from 'fs';
import { DB_PATH, SCHEMA_PATH } from '$env/static/private';
import { v4 as uuidv4 } from 'uuid';
import type { User } from '$lib/types/user';
import bcrypt from 'bcrypt'


const db = new Database(DB_PATH);

try {
  const sqlSchema = readFileSync(SCHEMA_PATH, { encoding: 'utf8' });
  db.exec(sqlSchema);

} catch (err) {
  if (err instanceof Error) {
    console.log(err.message);
  } else {
    console.log('An unknown error occurred');
  }
}



export function checkSessionExists(access_or_user_id: string) {
  type RowType = {
    user_id: string
  }

  try {
    const query = `
      SELECT user_id FROM session WHERE access_id = ? OR user_id = ? LIMIT 1;
      `

    const row = db.prepare(query).get(access_or_user_id, access_or_user_id) as RowType

    if (typeof row === 'object' && row !== null && row.user_id) {

      return { ok: true, user_id: row.user_id }
    } else {
      throw new Error("Session Not Foune")
    }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Misc Error" }
    }
  }
}


export function checkUserExists(login: string, password: string) {
  type RowType = {
    id: string
    password: string
  }
  try {
    const query = `
    SELECT id, password FROM user WHERE username = ? OR email = ? LIMIT 1;
    `

    const row = db.prepare(query).get(login, login) as RowType

    if (typeof row === 'object' && row !== null && row.password && row.id) {

      const bool = bcrypt.compareSync(password, row.password)

      return { ok: true, authorized: bool, id: row.id }
    } else {
      throw new Error("User with that username/email not found")
    }
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, authorized: false, error: err, message: err.message }
    } else {
      return { ok: false, authorized: false, error: err, message: "Misc Error" }
    }
  }
}

export function deleteSession(user_id: string) {
  const query = `DELETE FROM session WHERE user_id = ?;`

  try {
    db.prepare(query).run(user_id)
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err }
    }
  }

  return { ok: true }
}
export function createSession(access_id: string, user_id: string) {
  const query = `INSERT INTO session (access_id, user_id) VALUES (?, ?);`

  try {
    db.prepare(query).run(access_id, user_id)
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err }
    }
  }

  return { ok: true }
}


export function createUser(userInfo: User) {
  // return userInfo
  try {
    const query = `
    INSERT INTO 
      user 
      (id, username, email, password, date_of_birth, first_name, last_name) 
    VALUES
      (?, ?, ?, ?, ?, ?, ?);`
    const id = uuidv4()
    const salt = bcrypt.genSaltSync(10);

    const hash = bcrypt.hashSync(userInfo.password, salt);


    userInfo.password = hash


    db.prepare(query).run(id, ...userInfo)
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err }
    }
  }

  return { ok: true }
}
