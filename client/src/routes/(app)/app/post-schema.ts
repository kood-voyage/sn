import { z } from 'zod';

export const postSchema = z.object({
    title: z.string(),
    content: z.optional(z.string()),
    privacy: z.string(),
    images: z.optional(z.any())
});
export type PostSchema = typeof postSchema