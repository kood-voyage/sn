import { z } from 'zod';

export const postSchema = z.object({
    title: z.string(),
    content: z.optional(z.string()),
    privacy: z.string(),
});
export type PostSchema = typeof postSchema