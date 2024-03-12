import { redirect, type Handle } from '@sveltejs/kit';
import jwt from 'jsonwebtoken'
import { deleteTokens, refreshTokens } from '$lib/server/jwt-handle';
import { JWT_KEY } from '$env/static/private';


export const handle: Handle = async ({ event, resolve }) => {
  const access_token = event.cookies.get("at") as string
  const refresh_token = event.cookies.get("rt") as string
  const pathname = event.url.pathname



  jwt.verify(refresh_token, process.env.JWT_KEY || JWT_KEY, (err, rdecoded) => {

    if (err != null) {
      if (!(pathname.startsWith('/signin')) && !(pathname.startsWith('/signup'))) {
        deleteTokens(event)
        redirect(303, "/signin")
      }
    } else {

      jwt.verify(access_token, process.env.JWT_KEY || JWT_KEY, (err) => {

        if (err != null) {

          if (rdecoded && typeof rdecoded == "object") {
            refreshTokens(event, rdecoded.access_token_id)
          }

        } else {
          if (pathname.startsWith("/signin") || pathname.startsWith("/signup")) {
            redirect(303, "/app")
          }
        }

      })
    }
  })

  const response = await resolve(event);
  return response;
};
