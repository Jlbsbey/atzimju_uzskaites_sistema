import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function myDataEditor(
	ifAdmin: boolean,
	username: string,
	email: string,
	oldPassword: string,
	newPassword: string,
	newName: string,
	newSurname: string
) {
	let auth = getAuthCookie();

	let request = "changeData?username=" + username +
		"&email=" + email +
		"&auth=" + auth;

	if (!ifAdmin) {
		request +=
			"&oldPassword=" + oldPassword +
			"&newPassword=" + newPassword;
	} else {
		request +=
			"&newName=" + newName +
			"&newSurname=" + newSurname +
			"&newPassword=" + newPassword;
	}

	sendRequest(request).then(data => {
	});
}