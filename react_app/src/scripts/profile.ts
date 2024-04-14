import {sendRequest} from "./request";
import {getAuthCookie} from "./cookies";
import {Response} from "./data";

export function getProfileData(id: number): Promise<Response> {
	let authKey = getAuthCookie();
	let request = "profile?auth=" + authKey;
	request += id === -1 ? "" : "&user=" + id;
	let response = sendRequest(request);

	return response.then((data: Response) => {
		if (data.content.username == "") {
			window.location.href = "/404";
		}

		return data;
	});
}