import { LOCAL_PATH, S3_BUCKET } from "$env/static/private";
import { mainGetGroup } from "$lib/server/db/group"
import { saveToS3 } from "$lib/server/images/upload";
import { getUserIdFromCookie } from "$lib/server/jwt-handle";
import { postSchema } from "$lib/types/post-schema";
import type { Actions, PageServerLoad } from "./$types"
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';


export const load: PageServerLoad = async (event) => {
  const form = await superValidate(zod(postSchema));

  console.log(event.params.name)
  console.log((await event.parent()))
  const data = (await mainGetGroup(event, event.params.name))
  if (!data.ok) {
    console.error(data.message)
    return {
      group: { ...data }, form
    }
  }
  // console.log(typeof data)
  // console.log(typeof data.data)
  return { group: { ...data }, form }

}



//1. server folderis api folder api folderis grouprequests.ts --> fetch request there
//2. page.server.ts peab looma g sisse 
//3. returnin loadis data --> data tuleb valja page.sveltes --> naeb marco name filedes script +page.server.ts ja page.svelte
//4. 





//TODO: HERE IS ACTION TEMPLATE

export const actions: Actions = {
    groupPostSubmit: async (event) => {
        // Validate the form data
        // const form = await superValidate(event, zod(postSchema));


        // HERE IS GET A USER_ID 
        const { user_id } = getUserIdFromCookie(event)

        // HERE YOU CAN GET ALL FORM DATA
        const formData = await event.request.formData();

        // Extracting other form fields
        const post_id = uuidv4()
        const title = formData.get('title') as string;
        const content = formData.get('content') as string;
        const groupId = formData.get('groupId') as string;


        console.log(title)
        console.log(content)
        console.log(groupId) 




        // Extracting uploaded images

        //         const imagesURL: string[] = []
        // const images = formData.getAll('images') as File[];


        // for (const [i, image] of images.entries()) {
        //     const resp = await saveToS3(("post" + (i + 1)), post_id, image, "post")
        //     imagesURL.push(S3_BUCKET + resp)
        // }


        const json: Post = {
            id: post_id,
            user_id: user_id,
            title: title,
            content: content,
            community_id: "",
            image_path: imagesURL

        }
        console.log(json)

        try {
            const response = await fetch(`${LOCAL_PATH}/api/v1/auth/posts/create`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${event.cookies.get('at')}`,
                    'Content-Type': 'application/json' // Specify JSON content type
                },
                body: JSON.stringify(json) // Convert the JSON object to a string
            });

            if (!response.ok) {
                throw new Error('Failed to create post'); // Throw an error if response is not OK
            }

            // Handle successful response
        } catch (err) {
            if (err instanceof Error) {
                return { ok: false, error: err, message: err.message }
            } else {
                return { ok: false, error: err, message: "Unknown Error" }
            }
        }
        redirect(304, `/app/g`)
    }
};

function uuidv4() {
  throw new Error("Function not implemented.");
}

