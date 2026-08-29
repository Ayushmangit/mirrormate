import type { User } from "./user";

export type LoginPayload = {
	email: string;
	password: string;
};

export type RegisterPayload = {
	username: string;
	email: string;
	password: string;
};

export type LoginResponse = {
	user: User;
	access_token: string;
};

export type RegisterResponse = {
	user: User;
};
