/* eslint-disable @typescript-eslint/no-explicit-any */
import type { ChatLine } from "$lib/client/api/chat-requests";
import type { CustomNotification } from "$lib/store/websocket-store";
import type { UserType } from "./user";

export function isChatLine(obj: any): obj is ChatLine {
  return typeof obj.chat_id === 'string' &&
    typeof obj.created_at === 'string' &&
    typeof obj.id === 'string' &&
    typeof obj.message === 'string' &&
    typeof obj.user_id === 'string';
}

export function isNotification(obj: any): obj is CustomNotification {
  return typeof obj === 'object' &&
    obj !== null &&
    typeof obj.created_at === 'string' &&
    typeof obj.id === 'string' &&
    typeof obj.message === 'string' &&
    typeof obj.parent_id === 'string' &&
    typeof obj.source_id === 'string' &&
    // isUserType(obj.source_information) &&
    typeof obj.target_id === 'string' &&
    // isUserType(obj.target_information) &&
    typeof obj.type_id === 'number';
}

export function isUserType(obj: any): obj is UserType {
  return typeof obj === 'object' &&
    obj !== null &&
    typeof obj.id === 'string' &&
    typeof obj.username === 'string' &&
    typeof obj.email === 'string' &&
    typeof obj.password === 'string' &&
    typeof obj.dateOfBirth === 'string' &&
    typeof obj.firstName === 'string' &&
    typeof obj.lastName === 'string' &&
    typeof obj.event_status === 'string' &&
    (typeof obj.avatar === 'string' || obj.avatar === undefined) &&
    (typeof obj.cover === 'string' || obj.cover === undefined) &&
    (typeof obj.description === 'string' || obj.description === undefined);
}