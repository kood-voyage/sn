import { date, z } from 'zod';

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

export const groupSchema = z.object({
  name: z.string().min(1),
  description: z.string().min(8),
  image: fileSchema, // Use the refined file schema
  privacy: z.enum(["public", "private"], {
      required_error: "You need to select a notification type"
    }),
  
});

export type GroupSchema = typeof groupSchema

export const groupEventSchema = z.object({
  eventId: z.string(),
  userId: z.string(),
  groupId: z.string(),
  name: z.string().min(4).max(25),
  description: z.string().min(4).max(25),
  date: z.date(),
})




export const eventSchema = z.object({
	id: z.string(),
	userId: z.string(),
	groupId: z.string(),
  name: z.string().min(4).max(25),
  description: z.string().min(4).max(100)
});




export type EventSchema = typeof eventSchema;
