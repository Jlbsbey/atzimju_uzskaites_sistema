import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";

export function markEditor(
	username: string,
	subjectId: number,
	value: number,
	markId: number
) {
	let auth = getAuthCookie();

	let request = "addMark?username=" + username +
						 "&subject_id=" + subjectId +
						 "&value=" + value;
	console.log(markId)
	if (markId !== -1) {
		request += "&mark_id=" + markId;
	}
	request += "&auth=" + auth;

	sendRequest(request).then(data => {
		console.log(data);
	});
}