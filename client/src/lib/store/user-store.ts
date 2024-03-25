import {writable} from "svelte/store"; 




export const currentUserStore = writable({})
export const currentUserToken = writable()
export const currentUserFollowers = writable([])
export const currentUserFollowing = writable([])
