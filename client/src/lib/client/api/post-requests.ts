import { PUBLIC_LOCAL_PATH } from "$env/static/public"



export async function getUserPosts(user_id: string) {
  try {
    const resp = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/posts/${user_id}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Access-Control-Request-Method': 'GET'
		},
		credentials: 'include'

    })
    const json = (await resp.json()).data


    return { ok: true, data: json }

  } catch (err) {
    if (err instanceof Error) {
      return { ok: false, error: err, message: err.message }
    } else {
      return { ok: false, error: err, message: "Unknown Error" }
    }
  }

}






// 	export async function getUserPosts(userId: string) {
// 	if (userId === undefined) {
// 		return 0;
// 	}

// 	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/posts/${userId}`, {
// 		method: 'GET',
// 		headers: {
// 			'Content-Type': 'application/json',
// 			'Access-Control-Request-Method': 'GET'
// 		},
// 		credentials: 'include'
// 	});

// 	if (response) {
// 		return await response.json();
// 	} else {
// 		throw new Error('Failed to fetch posts');
// 	}
// }