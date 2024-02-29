import { z } from 'zod';

export const signUpSchema = z.object({
	username: z.string().refine((value) => /^[a-zA-Z].{3,31}$/.test(value), {
		message: 'Username must start with a letter and be between 4 and 32 characters.'
	}),
	email: z.string().email(),
	dateOfBirth: z.string(),
	password: z.string().min(8).max(32),
	firstName: z.string().min(2).max(32),
	lastName: z.string().min(2).max(32),
	// gender: z.string()
});



export const signInSchema = z.object({
	login: z.string(),
	password: z.string().min(8).max(32),

});




export type SignUpSchema = typeof signUpSchema;

export type SignInSchema = typeof signInSchema;
