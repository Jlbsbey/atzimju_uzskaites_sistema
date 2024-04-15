import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function addUser(
	name: string,
	surname: string,
	email: string,
	password: string,
	avatarLink: string,
	role: string
) {
	let auth = getAuthCookie();

	let request = "addMark?name=" + name +
						 "&surname=" + surname +
						 "&email=" + email +
						 "&password=" + password +
						 "&avatar_link=" + avatarLink +
						 "&role=" + role +
						 "&auth=" + auth;

	sendRequest(request).then(data => {
		console.log(data);
	});
}