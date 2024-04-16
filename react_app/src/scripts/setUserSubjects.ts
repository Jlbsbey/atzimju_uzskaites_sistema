import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function setUserSubjects(
	username: string,
	subjectIds: string
) {
	let auth = getAuthCookie();

	let request = "changeUserSubjects?username=" + username +
		"&subjects=" + subjectIds +
		"&auth=" + auth;

	sendRequest(request).then(data => {
		console.log(data);
	});
}