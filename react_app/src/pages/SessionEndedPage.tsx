import React from "react";
import "../styles/404_page.css";
import the404image from "../assets/images/404.gif";

const SessionEndedPage: React.FC = () => {
	return (
		<div className="the-404-page d-flex justify-content-center">
			<div>
				<h1 className="the-404-title">
					Your <span className="logo-animation">
							Grade
						</span> session ended... <br/>
					Please, <a href="/">relogin</a>.
				</h1>

				<img className="the-404-image" src={the404image}/>

			</div>
		</div>
	);
}

export default SessionEndedPage;