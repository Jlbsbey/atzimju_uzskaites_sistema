import React from "react";
import "../styles/about_page.css";

const LoginPage: React.FC = () => {
	return (
		<div className="about-page">
			<div>

				<h1 className="login-title">
					This is <span className="logo-animation">
							Grade
						</span> !
				</h1>

				<p className="text-main w-50-desktop">
					Welcome here!
					Grade is an system for students and professors
					to see and manage their grades.
				</p>

				<p className="text-main w-50-desktop">
					Since It is a simple and easy system, we are
					based on the users` reviews and feedbacks
					which you can leave <a href="/about" className="link">
						here
					</a>!
				</p>

				<h3>
					<strong>
						Who created this?
					</strong>
				</h3>

				<p className="text-main w-50-desktop">
					Grade is made by <a href="https://www.linkedin.com/in/nevolodia/">
						Volodia Kiril Bickov
					</a> and
					Nikita Smorigo.
				</p>




			</div>
		</div>
	);
}

export default LoginPage;