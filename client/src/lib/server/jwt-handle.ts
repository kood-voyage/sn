import jwt from 'jsonwebtoken'
import { v4 as uuidv4 } from 'uuid';
//import { JWT_KEY } from '$env/static/private';
import { redirect, type RequestEvent } from '@sveltejs/kit';
import type { RouteParams } from '../../routes/(auth)/signin/$types';
import { checkSessionExists, createSession, deleteSession } from './db';

type CreateEvent = RequestEvent<RouteParams, "/(auth)/signin">
type RefreshEvent = RequestEvent<Partial<Record<string, string>>, string | null>


const min15 = 60 * 15; // 15 minutes in seconds
const week = 60 * 60 * 24 * 7; // 7 days in seconds

export function createTokens(event: CreateEvent | RefreshEvent, user_id: string) {

  const access_token_id: string = uuidv4()

  const resp = createSession(access_token_id, user_id)
  if (!resp.ok) {
    return resp
  }

  const refresh_token = jwt.sign({
    exp: Math.floor(Date.now() / 1000) + week,
    access_token_id
  }, process.env.JWT_KEY, { algorithm: 'HS256' })


  const access_token = jwt.sign({
    exp: Math.floor(Date.now() / 1000) + min15,
    user_id,
    access_token_id
  }, process.env.JWT_KEY, { algorithm: 'HS256' })

  function timeConvert(time: number) {
    return new Date((new Date()).getTime() + time * 1000);
  }

  event.cookies.set("at", access_token, { path: "/", expires: timeConvert(min15) })
  event.cookies.set("rt", refresh_token, { path: "/", expires: timeConvert(week) })
  return { ok: true }
}

export function deleteTokens(event: RefreshEvent) {
  const negTime = new Date((new Date()).getTime() - 1 * 1000)
  event.cookies.set("rt", "", { path: "/", expires: negTime })
  event.cookies.set("at", "", { path: "/", expires: negTime })
}

export function refreshTokens(event: RefreshEvent, access_token_id: string) {
  const resp = checkSessionExists(access_token_id)

  if (resp.ok) {
    const user_id = resp.user_id as string
    deleteSession(user_id)
    createTokens(event, user_id)
  } else {
    if (!(event.url.pathname.startsWith('/signin')) || !(event.url.pathname.startsWith('/signup'))) {
      deleteTokens(event)
      console.error("Refresh token access_id wasn't found", resp.error)
      redirect(303, "/signin")
    }
  }
}

