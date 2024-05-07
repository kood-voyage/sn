import { z } from 'zod';

export const postSchema = z.object({
    title: z.string(),
    content: z.optional(z.string()),
    privacy: z.enum(["public", "private", "selected"], {
      required_error: "You need to select a notification type"
    }),
    images: z
    .instanceof(File, { message: 'Please upload a file.'})
    .refine((f) => f.size < 100_000, 'Max 100 kB upload size.')
    .array().max(3)


});
export type PostSchema = typeof postSchema



