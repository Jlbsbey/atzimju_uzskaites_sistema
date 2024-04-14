import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";
import {Response, User} from "./data";

export function getProfileData(id: number): Promise<Response> {
	let authKey = getAuthCookie();
	let response = sendRequest("profile?auth=" + authKey + "&id=" + id);

	return response.then((data: Response) => {
		return data;
	});
}