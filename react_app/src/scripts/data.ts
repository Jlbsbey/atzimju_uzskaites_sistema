export interface Response {
	status: string;
	error: string;
	content: any;
}

export interface Mark {
	mark_id: number;
	subject_id: number;
	professor_id: string;
	professor_name: string;
	student_id: number;
	student_name: string;
	value: number;
	create_date: string;
	edit_date: string;
}

export interface Subject {
	subject_id: number;
	subject_name: string;
	subject_description: string;
	is_active: boolean;
}

export interface User {
	user_id: number;
	username: string;
	name?: string;
}