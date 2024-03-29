import {setAuthCookie} from "./cookies";

interface Response {
	login_status: boolean;
	session_key: string;
	expire_time: string;
}

export function user_login(username: string, password: string): Promise<boolean> {
	let response = try_login("https://grade.nevolodia.com:8443/login", username, password);

	return response.then((data: Response) => {
		if (data.login_status) {
			setAuthCookie(data.session_key, data.expire_time);
			return true;
		}
		return false;
	});
}

function try_login(url: string, username: string, password: string): Promise<Response> {
	let urlRequest = url + "?username=" + username + "&password=" + password;
	// return as Response
	return fetch(urlRequest)
		.then(response => {
			if (!response.ok) {
				throw new Error('Network response was not ok');
			}
			return response.json();
		})
		.then(jsonData => {
			return jsonData;
		})
		.catch(error => {
			console.error('Error fetching data:', error.message);
			throw error;
		});
}