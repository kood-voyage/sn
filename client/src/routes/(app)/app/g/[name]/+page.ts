
import { superValidate, type SuperValidated } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { type ReturnType } from "$lib/types/requests";
// import { getGroupPosts, joinGroup, type GroupJson, type GroupPostJson } from "$lib/server/api/group-requests";
import { type User, type UserType } from "$lib/types/user";
import { redirect, type Actions } from '@sveltejs/kit';
import type { PageLoad } from '../$types';
import { eventSchema, groupPostSchema } from '$lib/types/group-schema';
import { GetAllInvitedUsers, GetGroup, GetGroupEvents, GetGroupPosts, JoinGroup, type AllGroupPosts, type Group, type GroupEventJson, type GroupPostJson } from '$lib/client/api/group-requests';
import { GetAllUsers, currentUser, getUserFollowers, type AllUsers } from '$lib/client/api/user-requests';
import { currentUserFollowers } from '$lib/store/user-store';

export const ssr = false;

// type GroupType = ReturnType<GroupJson>
// type AllUserType = ReturnType<UserRowType[]>

type LoadType = {
  group: Group | undefined,
  eventForm: SuperValidated<{
    id: string,
    userId: string,
    groupId: string,
    name: string,
    description: string,
}>,
  posts: GroupPostJson[] | undefined,
  allusers: UserType[] | undefined,
  allevents: GroupEventJson[] | undefined,
  allInvitedUsers: string[] 
}


export const load: PageLoad = async ({ fetch, url }) => {
    // const form = await superValidate(zod(groupPostSchema));
    const eventForm = await superValidate(zod(eventSchema))
    const u = new URL(url).pathname.split('/');
    const groupId = u[u.length - 1].replace("_", " ")
    
    let info: LoadType = {group: undefined, eventForm: eventForm, posts: undefined, allusers: undefined, allevents: undefined, allInvitedUsers: []
    }
    const groupInfo = (await GetGroup(groupId, fetch))
    if (!groupInfo.ok) {
      console.error(groupInfo.message)
      return info
    }
    info.group = groupInfo

    const groupPostData = (await GetGroupPosts(groupId, fetch))
    if (!groupPostData.ok) {
        console.error(groupPostData.message)
        return info
    }
    info.posts = groupPostData.allGroupPosts

    const allUsers = (await GetAllUsers(fetch))
      if (!allUsers.ok) {
        console.error("something wrong with getting all users")
        return info
      }
    if (allUsers.ok) {
      info.allusers = allUsers.allUsers
    }

    const allEvents = (await GetGroupEvents(groupInfo.group.id, fetch))
    if (!allEvents.ok) {
      console.error("something wrong with getting all the events")
      return info
    }
    if (allEvents.ok){
      info.allevents = allEvents.allGroupEvents
    }
    
    const allInvitedUsers = (await GetAllInvitedUsers(groupInfo.group.id, fetch))
    if (!allInvitedUsers.ok) {
      console.error("something wront with getting all the invited users", allInvitedUsers.message)
      return info
    }
    if (allInvitedUsers.allInvitedUsers) {
      info.allInvitedUsers = allInvitedUsers.allInvitedUsers
    }

    return info
}

// export const actions: Actions = {
  // groupPostSubmit: async (event) => {

  //   // // HERE IS GET A USER_ID 
  //   const { user_id } = getUserIdFromCookie(event)

  //   // // HERE YOU CAN GET ALL FORM DATA
  //   const formData = await event.request.formData();

  //   // // Extracting other form fields
  //   const post_id = uuidv4()
  //   const title = formData.get('title') as string;
  //   const content = formData.get('content') as string;
  //   const privacy = formData.get('privacy') as string;
  //   const groupId = event.params.name


  //   // Extracting uploaded images

  //   const imagesURL: string[] = []
  //   const images = formData.getAll('images') as File[];

  //   console.log(images)


  //   for (const [i, image] of images.entries()) {
  //     const resp = await saveToS3(("post" + (i + 1)), post_id, image, "post")
  //     imagesURL.push(S3_BUCKET + resp)
  //   }


  //   const json = {
  //     id: post_id,
  //     user_id: user_id,
  //     title: title,
  //     content: content,
  //     privacy: privacy,
  //     community_id: groupId.replace("_", " "),
  //     image_path: imagesURL
  //   }

  //   try {
  //     const response = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/create`, {
  //       method: 'POST',
  //       headers: {
  //         'Authorization': `Bearer ${event.cookies.get('at')}`,
  //         'Content-Type': 'application/json' // Specify JSON content type
  //       },
  //       body: JSON.stringify(json) // Convert the JSON object to a string
  //     });

  //     if (!response.ok) {
  //       throw new Error('Failed to create post'); // Throw an error if response is not OK
  //     }
  //     // Handle successful response
  //   } catch (err) {
  //     // Extracting only necessary information from the error object
  //     const errorMessage = err instanceof Error ? err.message : "Unknown Error";

  //     return { ok: false, error: errorMessage, message: errorMessage };
  //   }
  //   redirect(304, `/app/g/${event.params.name}`)
  // },
  // groupJoinSubmit: async (event) => {
  //   try {
  //     const resp = (await JoinGroup())
  //     if (!resp.ok) {
  //       return resp.error
  //     }
  //   } catch (err) {
  //     if (err instanceof Error) {
  //       return { ok: false, error: err, message: err.message }
  //     } else {
  //       return { ok: false, error: err, message: "Unknown Error" }
  //     }
  //   }
  // }
// }


//1. server folderis api folder api folderis grouprequests.ts --> fetch request there
//2. page.server.ts peab looma g sisse 
//3. returnin loadis data --> data tuleb valja page.sveltes --> naeb marco name filedes script +page.server.ts ja page.svelte
//4. 