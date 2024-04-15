import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function myDataEditor(
	username: string,
	email: string,
	oldPassword: string,
	newPassword: string
) {
	let auth = getAuthCookie();

	let request = "changeData?username=" + username +
		"changeData?email=" + email +
		"&oldPassword=" + oldPassword +
		"&newPassword=" + newPassword +
		"&auth=" + auth;
}