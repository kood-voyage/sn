import jwt from 'jsonwebtoken'
import { v4 as uuidv4 } from 'uuid';
import { JWT_KEY } from '$env/static/private';
import type { RequestEvent } from '@sveltejs/kit';
import type { RouteParams } from '../../routes/(auth)/signin/$types';

type CreateEvent = RequestEvent<RouteParams, "/(auth)/signin">

export function createTokens(event: CreateEvent, user_id: string) {

  const access_token_id:string = uuidv4()

  const refresh_token = jwt.sign({
      exp: Math.floor(Date.now() / 1000) + (60 * 60 *24 *7),
      access_token_id
  }, JWT_KEY,{ algorithm: 'HS256' })


  const access_token = jwt.sign({
			exp: Math.floor(Date.now() / 1000) + (60 * 15),
			user_id,
      access_token_id
	}, JWT_KEY,{ algorithm: 'HS256' })

  event.cookies.set("at", access_token,{path:"/"})
	event.cookies.set("rt", refresh_token,{path:"/"})

}

type RefreshEvent = RequestEvent<Partial<Record<string, string>>, string | null>

export function refreshAccessToken(event: RefreshEvent, access_token_id: string, user_id:string) {
  const access_token = jwt.sign({
			exp: Math.floor(Date.now() / 1000) + (60 * 15),
			user_id,
      access_token_id
	}, JWT_KEY,{ algorithm: 'HS256' })
  
  event.cookies.set("at", access_token,{path:"/"})
}

