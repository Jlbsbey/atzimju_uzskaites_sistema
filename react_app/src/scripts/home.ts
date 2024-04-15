import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";
import {Response, User} from "./data";

export function getHomeData(): Promise<Response> {
	let authKey = getAuthCookie();
	let response = sendRequest("home?auth=" + authKey);

	return response.then((data: Response) => {
		if (data.status === "error") {
			window.location.href = "/session_ended";
		}

		if (data.content.role === "admin") {
			window.location.href = "/admin";
			return data;
		}

		console.log(data);
		let marks = data.content.marks;
		let users = data.content.students;
		users.push(...data.content.professors);
		let subjects = data.content.subjects;

		console.log(marks)

		for (let mark of marks) {
			mark.professor_name = users.find((user: User) => user.user_id === mark.professor_id)?.name;
			mark.student_name = users.find((user: User) => user.user_id === mark.student_id)?.name;
		}

		data.content.marks = marks;

		return data;
	});
}