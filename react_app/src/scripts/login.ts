import {setAuthCookie} from "./cookies";

export function user_login(login: string, password: string): boolean {
	if (login === "admin" && password === "http://localhost:3000/login") {
		setAuthCookie("admin");
		return true;
	}
	return false;
}
