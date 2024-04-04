import { z } from 'zod';

const fileSchema = z.instanceof(File)


export const groupSchema = z.object({
  title: z.string().min(1),
  content: z.string().min(8),
  image: fileSchema, // Use the refined file schema
});

export type GroupSchema = typeof groupSchema