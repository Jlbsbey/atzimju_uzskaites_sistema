import React, {useState} from "react";
import "../styles/login_component.css";
import {user_login} from "../scripts/login";

const Login: React.FC = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [error, setError] = useState('');

	const submit = () => {
		if (user_login(username, password)) {
			window.location.href = "/main";
		} else {
			setError("Invalid username or password");
		}
	}

	return (
		<div className="input-container">

			{error === '' ? null : (
				<div className="login-alert alert alert-warning" role="alert">
					{error}
				</div>
			)}

			<div className="input-group mb-2">
				<span className="input-group-text"> @ </span>
				<input
					type="text"
					className="form-control"
					placeholder="Username"
					onChange={(event) => setUsername(event.target.value)}
				/>
			</div>

			<div className="input-group mb-2">
				<input
					type="password"
					className="form-control"
					placeholder="Password"
					onChange={(event) => setPassword(event.target.value)}
				/>
			</div>

			<button
				type="button"
				className="btn btn-outline-primary login-button"
				onClick={submit}
			>
				Go to Grade
			</button>

		</div>
	);
}

export default Login;