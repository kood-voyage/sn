
import { writable } from "svelte/store";


type User = {
  id: string,
  username: string,
  email: string,
  date_of_birth: string,
  first_name: string,
  last_name: string,
  avatar: string,
  cover: string,
  description: string
}
export const currentUserStore = writable<User | object>()
export const currentUserToken = writable()
export const currentUserFollowers = writable([])
export const currentUserFollowing = writable([])
