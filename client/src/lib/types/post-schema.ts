import { z } from 'zod';

export const postSchema = z.object({
    title: z.string(),
    content: z.string().min(8),
    privacy: z.string(),
    images: z.array(z.any()) // Temporarily treat images as an array of any type
        .max(3, { message: "You can only upload up to 3 images." })
        .optional(),
});
export type PostSchema = typeof postSchema