// import { redirect, type Handle, type RequestEvent } from '@sveltejs/kit';
// import jwt from 'jsonwebtoken'
// import { JWT_KEY } from '$env/static/private';


// export const handle: Handle = async ({ event, resolve }) => {
//   const refresh_token = event.cookies.get("rt") as string
//   const pathname = event.url.pathname
//   jwt.verify(refresh_token, process.env.JWT_KEY || JWT_KEY, (err) => {
//     if (err != null) {
//       deleteTokens(event)
//       if (!(pathname.startsWith('/signin')) && !(pathname.startsWith('/signup'))) {
//         redirect(303, "/signin")
//       }
//     } else {
//       if (pathname.startsWith("/signin") || pathname.startsWith("/signup")) {
//         redirect(303, "/app")
//       }
//     }
//   })
//   const response = await resolve(event);
//   return response;
// };

// function deleteTokens(event: RequestEvent<Partial<Record<string, string>>, string | null>) {
//   const negTime = new Date((new Date()).getTime() - 1 * 1000)
//   event.cookies.set("rt", "", { path: "/", expires: negTime, secure: false })
//   event.cookies.set("at", "", {
//     path: "/", expires: negTime, secure: false
//   })
// }