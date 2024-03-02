import Database from 'better-sqlite3'; 
import { readFileSync } from 'fs';
import {DB_PATH, SCHEMA_PATH} from '$env/static/private'; 
import {v4 as uuidv4} from 'uuid';
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

export function checkUserExists(login:string, password:string) {
  return {login, password}
  // try {
  //   const query = `
  //     SELECT 
  //     `
  //   const id = uuidv4()

  //   const password = 'secret';

  //   const salt = bcrypt.genSaltSync(10);

  //   const hash = bcrypt.hashSync(password, salt);

  //   console.log(hash)
    
  //   db.prepare(query).get(login)
  // } catch (err) {
  //   if (err instanceof Error) {
  //     return { ok: false, error: err, message: err.message}
  //   } else {
  //     return { ok: false, error: err}
  //   }
  // }
}

export function createUser(userInfo: User) {
  // return userInfo
  try {
    const query = 'INSERT INTO user (id, username, email, password, date_of_birth, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?, ?);'
    const id = uuidv4()

    const password = 'secret';

    const salt = bcrypt.genSaltSync(10);

    const hash = bcrypt.hashSync(password, salt);

    console.log(hash)
    userInfo.password = hash
    
    db.prepare(query).run(id, ...userInfo)
  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message}
    } else {
      return { ok: false, error: err}
    }
  }

  return { ok: true }
}

