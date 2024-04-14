export function sendRequest(request: string): Promise<any> {
	return fetch("https://grade.nevolodia.com:8443/" + request)
		.then(response => response.json())
		.then(data => {
			return data;
		});
}