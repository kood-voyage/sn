// import { LOCAL_PATH } from '$env/static/private';
// import type { Actions } from './$types';
// // import { getProfile } from "$lib/server/db/profile"
// import type { PageServerLoad } from './$types';


// export const load: PageServerLoad = async (event) => {
//   const data = (await getProfile(event, event.params.username))
//   if (data.error) {
//     console.error(data)
//     return {
//       username: "Undefined"
//     }
//   }
//   return data
// }

// export const actions: Actions = {
// 	follow: async (event) => {
//   const data = await event.request.formData()
//   const target_id = data.get("target_id")

//     try {
//     await fetch(`${LOCAL_PATH}/api/v1/auth/follow/${target_id}`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')}`
//       }
//     })

//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }

//   },

//   unfollow: async (event) => {
//   const data = await event.request.formData()
//   const target_id = data.get("target_id")

//     try {
//     await fetch(`${LOCAL_PATH}/api/v1/auth/unfollow/${target_id}`, {
//       headers: {
//         "Authorization": `Bearer ${event.cookies.get('at')}`
//       }
//     })
//   } catch (err) {
//     if (err instanceof Error) {
//       return { ok: false, error: err, message: err.message }
//     } else {
//       return { ok: false, error: err, message: "Unknown Error" }
//     }
//   }
//   }
// }


export const ssr = false 



import { PUBLIC_LOCAL_PATH } from '$env/static/public';
import type { PageLoad } from './$types';

export const load: PageLoad = async({ params }) => {




	async function getUser(username: string) {
		const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/get/${username}`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'Access-Control-Request-Method': 'GET'
			},
			credentials: 'include'
		});

		if (response.ok) {
			return response.json();
		} else {
			throw new Error('Failed to fetch users');
		}
	}


	const user = await getUser(params.username)





	async function getPosts(userId: string) {
	if (userId === undefined) {
		return 0;
	}

	const response = await fetch(`${PUBLIC_LOCAL_PATH}/api/v1/auth/user/posts/${userId}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
			'Access-Control-Request-Method': 'GET'
		},
		credentials: 'include'
	});

	if (response) {
		return await response.json();
	} else {
		throw new Error('Failed to fetch posts');
	}
}



	const posts = await getPosts(user.data.id)






	

	return {params,posts,user};
};