import Database from 'better-sqlite3'; 
import { readFileSync } from 'fs';
import {DB_PATH, SCHEMA_PATH} from '$env/static/private'; 
import {v4 as uuidv4} from 'uuid';
import type { User } from '$lib/types/user';


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

export function createUser(userInfo: User) {
  try {
    const query = `
      INSERT INTO user 
      (id, username, email, password, date_of_birth, first_name, last_name)
      VALUES
      (?, ?, ?, ?, ?, ?, ?)
      `
    const id = uuidv4()
    
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


