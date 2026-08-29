
export type User = {
	id: number;
	email: string;
	username: string;
	is_active: boolean;
	role_id: number;
	role: Role;
	created_at: string;
	updated_at: string;
};


export type Role = {
	id: number
	name: string
	level: number
}
