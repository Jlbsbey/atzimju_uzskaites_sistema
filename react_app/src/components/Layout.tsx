import { Outlet } from "react-router-dom";
import '../styles/layout.css';

const Layout = () => {
	return (
		<>
			<div className="menu navbar">
				<a className="menu-logo" href="/">
					Grade
				</a>
				<a className="menu-link" href="/my-profile/">
					Grade
				</a>
			</div>

			<div className="content">
				<Outlet/>
			</div>
		</>
	)
};

export default Layout;