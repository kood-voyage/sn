// import { CLIENT_IP, CLIENT_PROTO } from '$env/static/private';
import type { Handle } from '@sveltejs/kit';
import { JWT_KEY } from '$env/static/private';
import jwt from 'jsonwebtoken'
import { refreshAccessToken } from '$lib/server/jwt-handle';


export const handle: Handle = async ({ event, resolve }) => {
  const access_token = event.cookies.get("at") as string
  const refresh_token = event.cookies.get("st") as string

  //check if tokens are in cookies?

  try {
    jwt.verify(refresh_token, JWT_KEY, (err, decoded) => {
      if (err == null) {
        jwt.verify(access_token, JWT_KEY, (err, decoded) => {
          if (err?.name == "TokenExpiredError") {
            refreshAccessToken(event, "asd", "dsa")
            console.log("error >>>", err.message)
          }
          console.log(decoded)
          // console.log(err)
        })

      }
      // console.log(err)
      console.log(decoded)
    })

  } catch (err) {
    if (err instanceof Error) {
      console.error(err)
    } else {
      console.error(err)
    }
  }




  // const time_now = Date.now()
  // console.log(time_now)
  // console.log(at_decoded)

  if (event.url.pathname.startsWith('/custom')) {
    // console.log(response)
    return new Response('custom response');
  }

  const response = await resolve(event);
  return response;
};

