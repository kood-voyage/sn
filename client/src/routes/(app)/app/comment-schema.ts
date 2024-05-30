import { z } from 'zod';
export const commentSchema = z.object({
    user_id: z.string(),
    user_name: z.string(),
    user_avatar: z.string(),
    post_id: z.string(),
    parent_id: z.optional(z.string()),
    content: z.string().min(1).max(300),
    privacy: z.string(),
    path: z.optional(z.array(z.string())),
    count: z.optional(z.string())
});
export type CommentSchema = typeof commentSchema