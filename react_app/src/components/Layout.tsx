import { Outlet } from "react-router-dom";
import '../styles/layout.css';
import {getAuthCookie} from "../scripts/cookies";

const Layout = () => {
	let authKey = getAuthCookie();

	return (
		<>
			<div className="menu navbar">
				<a className="menu-logo" href="/">
					Grade
				</a>
				<div>
					<a className="menu-link" href="/main/">
						Grades
					</a>
					<a className="menu-link" href="/user/">
						My Profile
					</a>
				</div>
			</div>

			<div className="content">
				<Outlet/>
			</div>

			<div className="footer">
				<a className="footer-link" href="/about/">
					About Grade
				</a>
				<p className="footer-text">
					© 2024 Grade
				</p>
			</div>
		</>
	)
};

export default Layout;