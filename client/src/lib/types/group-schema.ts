import { z } from 'zod';

const fileSchema = z.instanceof(File)


export const groupPostSchema = z.object({
  groupId: z.string(),
  title: z.string(),
  content: z.string().min(8),
  images: z
    .array(z.any()) // Temporarily treat images as an array of any type
    .max(3, { message: 'You can only upload up to 3 images.' })
    .optional()
});

export type GroupSchema = typeof groupSchema


export const groupSchema = z.object({
  title: z.string().min(1),
  content: z.string().min(8),
  image: fileSchema, // Use the refined file schema
});

export type GroupSchema = typeof groupSchema