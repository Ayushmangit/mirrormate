export interface LoginUser {
	email: string
	password: string
}


export interface LoginResponse {
	user: User,
	access_token: string
}

export interface RegisterUser {
	name: string
	email: string
	password: string
}

export interface User {
	id: number
	name: string
	email: string
	createdAt: string
	updatedAt: string
}

