
import type { UserType } from "$lib/types/user";
import { writable } from "svelte/store";

export const currentUserStore = writable<UserType>()
export const currentUserToken = writable()
export const currentUserFollowers = writable([])
export const currentUserFollowing = writable([])
