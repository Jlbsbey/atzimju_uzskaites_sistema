import {setAuthCookie} from "./cookies";
import {sendRequest} from "./request";

export function user_login(username: string, password: string): Promise<boolean> {
	let response = sendRequest("login?username=" + username + "&password=" + password);

	return response.then((data: any) => {
		if (data.login_status) {
			setAuthCookie(data.session_key, data.expire_time);
			return true;
		}
		return false;
	});
}
