import React from "react";
import "../styles/404_page.css";
import the404image from "../assets/images/404.gif";

const The404Page: React.FC = () => {
	return (
		<div className="the-404-page d-flex justify-content-center">
			<div>
				<h1 className="the-404-title">
					This page is not on the <span className="logo-animation">
							Grade
						</span> ...
				</h1>

				<img className="the-404-image" src={the404image}/>

			</div>
		</div>
	);
}

export default The404Page;