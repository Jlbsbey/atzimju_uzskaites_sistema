import React, {useState} from "react";
import "../styles/login_component.css";

const LoginComponent: React.FC = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [error, setError] = useState('');

	const submit = () => {
		console.log("login as" + username + " with password " + password);
		setError("Invalid username or password");
	}

	return (
		<div className="input-container">

			{error === '' ? null : (
				<div className="login-alert alert alert-warning mb-4" role="alert">
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
				Log in to Grade
			</button>

		</div>
	);
}

export default LoginComponent;