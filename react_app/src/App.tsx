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

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Layout />}>
					<Route index element={<LoginPage />} />
					<Route path="main" element={<MainPage />} />
					<Route path="about" element={<AboutPage />} />
					<Route path="*" element={<The404Page />} />
				</Route>
			</Routes>
		</BrowserRouter>
	);
}

export default App;
