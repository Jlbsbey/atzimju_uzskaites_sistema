import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function addSubject(
	subject: string,
	description: string
) {
	let auth = getAuthCookie();

	let request = "addSubject?subject=" + subject +
						 "&description=" + description +
						 "&auth=" + auth;

	sendRequest(request).then(data => {
		console.log(data);
	});
}