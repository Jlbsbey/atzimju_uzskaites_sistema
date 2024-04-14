export function setAuthCookie(session_key: string, expire_time: string) {
	document.cookie = "session_key=" + session_key + "; expires=" + expire_time + "; path=/";
}

export function getAuthCookie() {
	let cookies = document.cookie.split(";");
	let session_key = "";

	cookies.forEach((cookie) => {
		if (cookie.includes("session_key")) {
			session_key = cookie.split("=")[1];
		} else {
			window.location.href = "/session_ended";
		}
	});

	return session_key;
}

export function removeAuthCookie() {
	document.cookie = "session_key=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}