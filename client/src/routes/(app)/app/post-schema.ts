import { z } from 'zod';

export const postSchema = z.object({
    title: z.string().min(4),
    content: z.string().min(12),
    privacy: z.enum(["public", "private", "selected"], {
      required_error: "You need to select a notification type"
    }),
    images: z.any()
});
export type PostSchema = typeof postSchema



