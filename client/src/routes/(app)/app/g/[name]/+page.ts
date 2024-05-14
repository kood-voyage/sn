
import { superValidate, type SuperValidated } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { type UserType } from "$lib/types/user";
import type { PageLoad } from '../$types';
import { groupPostSchema } from '$lib/types/group-schema';
import { GetAllInvitedUsers, GetGroup, GetGroupEvents, GetGroupPosts, type Group, type GroupEventJson, type GroupPostJson } from '$lib/client/api/group-requests';
import { GetAllUsers } from '$lib/client/api/user-requests';

export const ssr = false;

type LoadType = {
  group: Group | undefined,
  form: SuperValidated<{
    title: string;
    content: string;
    groupId: string;
    images?: any[] | undefined;
  }, any, {
    title: string;
    content: string;
    groupId: string;
    images?: any[] | undefined;
  }>,
  posts: GroupPostJson[] | undefined,
  allusers: UserType[] | undefined,
  allevents: GroupEventJson[] | undefined,
  allInvitedUsers: string[] | undefined
}


export const load: PageLoad = async ({ fetch, url }) => {
  const form = await superValidate(zod(groupPostSchema));
  const u = new URL(url).pathname.split('/');
  const groupId = u[u.length - 1].replaceAll("_", " ")

  const info: LoadType = {
    group: undefined, form: form, posts: undefined, allusers: undefined, allevents: undefined, allInvitedUsers: undefined
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
  if (allEvents.ok) {
    info.allevents = allEvents.allGroupEvents
  }

  const allInvitedUsers = (await GetAllInvitedUsers(groupInfo.group.id, fetch))
  if (!allInvitedUsers) {
    console.error("something went wrong with getting all invited users")
    return info
  }
  if (allInvitedUsers.ok){
    info.allInvitedUsers = allInvitedUsers.allInvitedUsers
  }

  return { ...info }
}