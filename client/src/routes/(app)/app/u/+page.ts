
// import { PUBLIC_LOCAL_PATH } from "$env/static/public"
// export const load = async () => {
//     try {
//     const resp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/all`, {
//       method: "GET",
//       headers: {
//         "Content-Type": "application/json",
//         "Access-Control-Request-Method": "GET",
//       },
//       credentials: "include",
//     })


//     console.log(resp)
//     return {resp}

//   } catch (err) {
//     console.log("ERRRRRR", err)
//     if (err instanceof Error) {

//       console.log(err)
//       return { ok: false, error: err, message: err.message }
//     } else {

//       console.log(err)
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }

// }