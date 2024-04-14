export function sendRequest(request: string): Promise<any> {
	return fetch("https://grade.nevolodia.com:8443/" + request)
		.then(response => response.json())
		.then(data => {
			if (data.error === "Session expired") {
				window.location.href = "/session_ended";
			}
			return data;
		});
}