import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";
import {Response, User} from "./data";

export function getAdminData(): Promise<Response> {
	let authKey = getAuthCookie();
	let response = sendRequest("home?auth=" + authKey);

	return response.then((data: Response) => {
		console.log(data);
		let users = data.content.students;
		users.push(...data.content.professors);
		let subjects = data.content.subjects;

		return data;
	});
}