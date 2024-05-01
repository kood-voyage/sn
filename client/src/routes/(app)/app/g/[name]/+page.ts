
import { superValidate, type SuperValidated } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { type ReturnType } from "$lib/types/requests";
// import { getGroupPosts, joinGroup, type GroupJson, type GroupPostJson } from "$lib/server/api/group-requests";
import { type User } from "$lib/types/user";
import { redirect } from '@sveltejs/kit';
import type { PageLoad } from '../$types';
import { groupPostSchema } from '$lib/types/group-schema';
import { GetGroup, GetGroupPosts } from '$lib/client/api/group-requests';
import { GetAllUsers } from '$lib/client/api/user-requests';

export const ssr = false;

// type GroupType = ReturnType<GroupJson>
// type AllUserType = ReturnType<UserRowType[]>

// type LoadType = {
//   group: GroupType,
//   form: SuperValidated<{ title: string; content: string; privacy: string; images?: any[] | undefined; }, any, { title: string; content: string; privacy: string; images?: any[] | undefined; }>,
//   data: User,
//   posts: GroupPostJson[],
// allusers: AllUserType
// }


export const load: PageLoad = async ({ fetch, url }) => {
    const form = await superValidate(zod(groupPostSchema));
    const u = new URL(url).pathname.split('/');
    const groupId = u[u.length - 1].replace("_", " ")
    const groupPostData = (await GetGroupPosts(groupId, fetch))
    if (!groupPostData.ok) {
        console.error(groupPostData.message)
        // return {data: {}, posts: undefined, form: form, group: undefined}
    }

      const groupInfo = (await GetGroup(groupId, fetch))
      if (!groupInfo.ok) {
        console.error(groupInfo.message)
      }

      console.log("GROUPINF", groupInfo)


      const allUsers = (await GetAllUsers(fetch))
      if (!allUsers.ok) {
        console.error("something wrong with getting all users")
      }
      console.log(allUsers)
    //   console.log("THIS IS ALL THE USESRS", allUsers)

    //   let info: LoadType = { data: parentData.data, posts: groupPostData.data, form: form, group: { ...data }, allusers: allUsers.data };
    // console.log(typeof data)
    // console.log(typeof data.data)

    return {}
}

// export const actions: Actions = {
//   groupPostSubmit: async (event) => {

//     // // HERE IS GET A USER_ID 
//     const { user_id } = getUserIdFromCookie(event)

//     // // HERE YOU CAN GET ALL FORM DATA
//     const formData = await event.request.formData();

//     // // Extracting other form fields
//     const post_id = uuidv4()
//     const title = formData.get('title') as string;
//     const content = formData.get('content') as string;
//     const privacy = formData.get('privacy') as string;
//     const groupId = event.params.name


//     // Extracting uploaded images

//     const imagesURL: string[] = []
//     const images = formData.getAll('images') as File[];

//     console.log(images)


//     for (const [i, image] of images.entries()) {
//       const resp = await saveToS3(("post" + (i + 1)), post_id, image, "post")
//       imagesURL.push(S3_BUCKET + resp)
//     }


//     const json = {
//       id: post_id,
//       user_id: user_id,
//       title: title,
//       content: content,
//       privacy: privacy,
//       community_id: groupId.replace("_", " "),
//       image_path: imagesURL
//     }

//     try {
//       const response = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/create`, {
//         method: 'POST',
//         headers: {
//           'Authorization': `Bearer ${event.cookies.get('at')}`,
//           'Content-Type': 'application/json' // Specify JSON content type
//         },
//         body: JSON.stringify(json) // Convert the JSON object to a string
//       });

//       if (!response.ok) {
//         throw new Error('Failed to create post'); // Throw an error if response is not OK
//       }
//       // Handle successful response
//     } catch (err) {
//       // Extracting only necessary information from the error object
//       const errorMessage = err instanceof Error ? err.message : "Unknown Error";

//       return { ok: false, error: errorMessage, message: errorMessage };
//     }
//     redirect(304, `/app/g/${event.params.name}`)
//   },
//   groupJoinSubmit: async (event) => {
//     try {
//       const resp = (await joinGroup(event, event.params.name))
//       if (!resp.ok) {
//         return resp.error
//       }
//     } catch (err) {
//       if (err instanceof Error) {
//         return { ok: false, error: err, message: err.message }
//       } else {
//         return { ok: false, error: err, message: "Unknown Error" }
//       }
//     }
//   }
// }


//1. server folderis api folder api folderis grouprequests.ts --> fetch request there
//2. page.server.ts peab looma g sisse 
//3. returnin loadis data --> data tuleb valja page.sveltes --> naeb marco name filedes script +page.server.ts ja page.svelte
//4. 