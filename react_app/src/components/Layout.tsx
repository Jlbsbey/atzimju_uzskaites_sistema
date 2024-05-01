import { Outlet } from "react-router-dom";
import '../styles/layout.css';

const Layout = () => {
	let title = "Grade1";
	let path = window.location.pathname;
	if (path === "/kabul" || path === "/kabul/") {
		title = "Kabul International";
	}

	return (
		<>
			<div className="menu navbar">
				<a className="menu-logo" href="/">
					{title}
				</a>
				<div>
					<a className="menu-link" href="/main/">
						{title === "Grade" ? "Grades" : "News archive"}
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
					About {title}
				</a>
				<p className="footer-text">
					© 2024 {title}. All rights reserved.
				</p>
			</div>
		</>
	)
};

export default Layout;