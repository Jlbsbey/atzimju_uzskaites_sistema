export interface Grade {
	id: number;
	subject_id: number;
	value: number;
	professor: string;
	created_unix: number;
	last_updated_unix: number;
}

export interface Subject {
	subject_id: number;
	name: string;
	description: string;
	last_updated: string;
	professor: string;
	is_active: boolean;
}

export const subjects : Subject[] = [
	{
		subject_id: 1,
		name: "Mathematics",
		description: "The basics of mathematics with extended probability theory.",
		last_updated: "1",
		professor: "John von Neumann",
		is_active: true,
	},
	{
		subject_id: 2,
		name: "Calculus",
		description: "Calculus theory with practical examples.",
		last_updated: "2",
		professor: "John von Neumann",
		is_active: false,
	},
	{
		subject_id: 3,
		name: "Hebrew",
		description: "עברית, שפה שמית, היא השפה הכנענית היחידה החיה.",
		last_updated: "5",
		professor: "Jesus Christ",
		is_active: false,
	},
	{
		subject_id: 4,
		name: "Mathematics",
		description: "The basics of mathematics with extended probability theory.",
		last_updated: "1",
		professor: "John von Neumann",
		is_active: true,
	},
	{
		subject_id: 5,
		name: "Calculus",
		description: "Calculus theory with practical examples.",
		last_updated: "2",
		professor: "John von Neumann",
		is_active: false,
	},
	{
		subject_id: 6,
		name: "Hebrew",
		description: "עברית, שפה שמית, היא השפה הכנענית היחידה החיה.",
		last_updated: "5",
		professor: "Jesus Christ",
		is_active: false,
	},
	{
		subject_id: 7,
		name: "Mathematics",
		description: "The basics of mathematics with extended probability theory.",
		last_updated: "1",
		professor: "John von Neumann",
		is_active: true,
	},
	{
		subject_id: 8,
		name: "Calculus",
		description: "Calculus theory with practical examples.",
		last_updated: "2",
		professor: "John von Neumann",
		is_active: false,
	},
	{
		subject_id: 9,
		name: "Hebrew",
		description: "עברית, שפה שמית, היא השפה הכנענית היחידה החיה.",
		last_updated: "5",
		professor: "Jesus Christ",
		is_active: false,
	},
	{
		subject_id: 10,
		name: "Mathematics",
		description: "The basics of mathematics with extended probability theory.",
		last_updated: "1",
		professor: "John von Neumann",
		is_active: true,
	},
	{
		subject_id: 11,
		name: "Calculus",
		description: "Calculus theory with practical examples.",
		last_updated: "2",
		professor: "John von Neumann",
		is_active: false,
	},
	{
		subject_id: 12,
		name: "Hebrew",
		description: "עברית, שפה שמית, היא השפה הכנענית היחידה החיה.",
		last_updated: "5",
		professor: "Jesus Christ",
		is_active: false,
	},
];

export const grades : Grade[] = [
	{
		id: 1,
		subject_id: 1,
		value: 8,
		professor: "John von Neumann",
		created_unix: 24.04,
		last_updated_unix: 25.04
	},
	{
		id: 2,
		subject_id: 1,
		value: 10,
		professor: "John von Neumann",
		created_unix: 13.03,
		last_updated_unix: 13.03
	},
	{
		id: 3,
		subject_id: 2,
		value: 6,
		professor: "John von Neumann",
		created_unix: 24.04,
		last_updated_unix: 25.04
	},
	{
		id: 4,
		subject_id: 2,
		value: 7,
		professor: "John von Neumann",
		created_unix: 13.03,
		last_updated_unix: 13.03
	},
	{
		id: 5,
		subject_id: 3,
		value: 10,
		professor: "Jesus Christ",
		created_unix: 24.04,
		last_updated_unix: 25.04
	},
	{
		id: 6,
		subject_id: 3,
		value: 10,
		professor: "Jesus Christ",
		created_unix: 13.03,
		last_updated_unix: 13.03
	},
];


