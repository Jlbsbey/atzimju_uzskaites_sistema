import React from "react";
import "../styles/login_page.css";
import Login from "../components/Login";

const LoginPage: React.FC = () => {
	return (
		<div className="login-page d-flex justify-content-center">
			<div>
				<h1 className="login-title">
					Welcome to the <span className="logo-animation">
							Grade
						</span> !
				</h1>

				<Login/>
			</div>
		</div>
	);
}

export default LoginPage;