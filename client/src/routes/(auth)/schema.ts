import { z } from 'zod';

export const signUpSchema = z.object({
username: z.string().refine((value) => /^[a-zA-Z][a-zA-Z0-9]{3,31}$/.test(value), {
    message: 'Username must start with a letter and be between 4 and 32 characters.[A-z0-9]'
  }),
	email: z.string().email(),
	dateOfBirth: z.string(),
	password: z.string().min(8).max(32),
	repeatPassword:z.string(),
	firstName: z.string().min(2).max(32),
	lastName: z.string().min(2).max(32),
}).refine(data=> data.password === data.repeatPassword,{
	message:'Password do not match',
	path:['repeatPassword'],
});

export const signInSchema = z.object({
	login: z.string(),
	password: z.string().min(8).max(32),
});
export type SignUpSchema = typeof signUpSchema;
export type SignInSchema = typeof signInSchema;
