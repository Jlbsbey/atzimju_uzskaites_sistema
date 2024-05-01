import React from 'react';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import "./styles/global.css";
import "./styles/typography.css";
import "./styles/layout_global.css";
import "bootstrap/dist/js/bootstrap.bundle.js.map";
import "bootstrap/dist/css/bootstrap.min.css";
import Layout from "./components/Layout";
import LoginPage from "./pages/LoginPage";
import MainPage from "./pages/MainPage";
import The404Page from "./pages/The404Page";
import AboutPage from "./pages/AboutPage";
import SessionEndedPage from "./pages/SessionEndedPage";
import UserPage from "./pages/UserPage";
import AdminHomePage from "./pages/AdminHomePage";
import SchoolNewsPage from "./pages/SchoolNewsPage";
import SchoolNewsPage2 from "./pages/SchoolNewsPage2";

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Layout />}>
					<Route index element={<LoginPage />} />
					<Route path="main" element={<MainPage />} />
					<Route path="user" element={<UserPage />} />
					<Route path="about" element={<AboutPage />} />
					<Route path="session_ended" element={<SessionEndedPage />} />
					<Route path="admin" element={<AdminHomePage />} />

					<Route path="kabul" element={<SchoolNewsPage />} />
					<Route path="kabul-lq" element={<SchoolNewsPage2 />} />
					<Route path="*" element={<The404Page />} />
				</Route>
			</Routes>
		</BrowserRouter>
	);
}

export default App;
