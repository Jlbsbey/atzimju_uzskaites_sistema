import React, {useEffect, useState} from "react";
import "../styles/user_page.css";
import {Subject} from "../scripts/data";
import 'simplebar-react/dist/simplebar.min.css';
import {getProfileData} from "../scripts/profile";

const UserPage: React.FC = () => {
	const [subjects, setSubjects] = useState<Subject[]>(
		[
			{
				subject_id: 0,
				subject_name: "Loading...",
				subject_description: "Loading...",
				is_active: false
			}
		]
	);
	const [generalData, setGeneralData] = useState({
		loaded: false,
		user_id: 0,
		role: "",
		name: "",
		surname: "",
		username: "",
		email: "",
		avatar_link: ""
	});

	// read from html params
	const urlParams = new URLSearchParams(window.location.search);
	const userIdString = urlParams.get('id');
	let userId = parseInt(userIdString ? userIdString : "-1");
	if (!userIdString)
		userId = -1;

	useEffect(() => {
		getProfileData(userId).then((data) => {
			console.log(data)
			setSubjects(data.content.subjects);
			setGeneralData(
				{
					loaded: true,
					user_id: data.content.user_id,
					role: data.content.role,
					username: data.content.username,
					name: data.content.name,
					surname: data.content.surname,
					email: data.content.email,
					avatar_link: data.content.avatar_link
				}
			)
		});
	}, []);

	return (
		<div className="user-page">

			<div className="user-beautiful">
				<img src={generalData.avatar_link}
				     alt="avatar"
				     className="avatar"
				/>
				<h2 className="user-name">
					{generalData.name} {generalData.surname} <br/>
				</h2>
				<h4 className="role">
					{generalData.role}
				</h4>
			</div>

			<h4 className="user-data">
				Username: {generalData.username} <br/>
				Email: <a href={"mailto:" + generalData.email}>{generalData.email}</a>
				{generalData.loaded ? "" : "Loading..."}
			</h4>

			<button className="btn btn-primary"
			        style={{
				        padding: "4px 16px",
				        fontSize: "16px",
				        marginBottom: "16px"
			        }}
			>
				Change my data
			</button>

			<table className="grade-table table table-sm table-striped">
				<thead>
					<tr>
						<th className="grade-table-t-first">
							{generalData.name} {generalData.surname}'s subjects
						</th>
					</tr>
				</thead>
				<tbody className="table-group-divider">
					{subjects.map((subject) => (
						<tr key={subject.subject_id}>
							<td className="grade-table-t-first">
								{subject.subject_name}
							</td>
						</tr>
					))}
				</tbody>
			</table>

			<h4>
			</h4>

		</div>
	);
}

export default UserPage;