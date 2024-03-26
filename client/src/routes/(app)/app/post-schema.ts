import { z } from 'zod';

export const postSchema = z.object({
    title: z.string(),
    content: z.string().min(8),
    privacy: z.string()
});
export type PostSchema = typeof postSchema