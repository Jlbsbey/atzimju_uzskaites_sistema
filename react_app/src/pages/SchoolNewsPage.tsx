import React from "react";
import "../styles/user_page.css";
import 'simplebar-react/dist/simplebar.min.css';
import logo from "../assets/images/kabul.png";
import PeVideo from "../components/PeVideo";

const SchoolNewsPage: React.FC = () => {

	return (
		<div className="">
			<h1 className="">
				[Archive] Kabul International - 26.12.1979
			</h1>
			<p style={{fontSize: "20px", fontWeight: "500"}}>
				The breaking news about soviet invasion of Afghanistan. <br/>
				Learn how Afghanistan is living, fighting, and suffering <br/>
				on the 3rd day of the invasion.
			</p>
			<img src={logo}
				 style={{width: "35%", marginBottom: "30px"}}
			/>

			<div style={{width: "85%"}}>
				<PeVideo src={"https://negallery.ams3.cdn.digitaloceanspaces.com/temp/final-lq.mp4"}
						 provider={"html5"}
				/>
			</div>

		</div>
	);
}

export default SchoolNewsPage;