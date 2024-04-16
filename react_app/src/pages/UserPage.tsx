import React, {useEffect, useState} from "react";
import "../styles/user_page.css";
import {Subject} from "../scripts/data";
import 'simplebar-react/dist/simplebar.min.css';
import {getProfileData} from "../scripts/profile";
import MyDataEditor from "../components/MyDataEditor";
import {myDataEditor} from "../scripts/myDataEditor";
import {setUserSubjects} from "../scripts/setUserSubjects";

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
		if_myself: false,
		if_admin: false,
		user_id: 0,
		role: "",
		name: "",
		surname: "",
		username: "",
		email: "",
		avatar_link: ""
	});

	const [editMyData, setEditMyData] = useState(false);
	const [editSubjects, setEditSubjects] = useState(false);

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
					if_myself: data.content.if_myself,
					if_admin: data.content.if_admin,
					user_id: data.content.user_id,
					role: data.content.role,
					username: data.content.username,
					name: data.content.name,
					surname: data.content.surname,
					email: data.content.email,
					avatar_link: data.content.avatar_link
				}
			)

			if (generalData.if_myself && generalData.if_admin) {
				window.location.href = "/404";
			}
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

			{(generalData.if_myself || generalData.if_admin) ? (
				<>
					{editMyData ? <MyDataEditor isAdmin={generalData.if_admin}/> : ""}

					<button className="btn btn-primary"
					        style={{
						        padding: "4px 16px",
						        fontSize: "16px",
						        marginBottom: "16px"
					        }}
					        onClick={() => {
								if (editMyData) {
									let newEmailElement = document.getElementById("new-email-input") as HTMLInputElement;
									let oldPasswordElement = document.getElementById("old-password-input") as HTMLInputElement;
									let newPasswordElement = document.getElementById("new-password-input") as HTMLInputElement;
									let newNameElement = document.getElementById("new-name-input") as HTMLInputElement;
									let newSurnameElement = document.getElementById("new-surname-input") as HTMLInputElement;

									let newEmail = newEmailElement ? newEmailElement.value : '';
									let oldPassword = oldPasswordElement ? oldPasswordElement.value : '';
									let newPassword = newPasswordElement ? newPasswordElement.value : '';
									let newName = newNameElement ? newNameElement.value : '';
									let newSurname = newSurnameElement ? newSurnameElement.value : '';

									myDataEditor(
										generalData.if_admin,
										generalData.username,
										newEmail,
										oldPassword,
										newPassword,
										newName,
										newSurname
									);

									// reload after .5s
									setTimeout(() => {
										window.location.reload();
									}, 500);
								}

								setEditMyData(!editMyData)
					        }}
					>
						Change data
					</button>
				</>
			) : ""}

			<table className="grade-table table table-sm table-striped">
				<thead>
					<tr>
						<th className="grade-table-t-first">
							{generalData.name} {generalData.surname}'s subjects
						</th>
					</tr>
				</thead>
				<tbody className="table-group-divider">
					{subjects?.map((subject) => (
						<tr key={subject.subject_id}>
							<td className="grade-table-t-first">
								{subject.subject_name}
							</td>
						</tr>
					))}
				</tbody>
			</table>

			{generalData.if_admin && (
				<>
					{editSubjects ?
						<div style={{
							marginBottom: "8px"
						}}>
							<textarea
							       style={{
								       width: "290px",
								       height: "120px",
								       minHeight: "60px",
								       maxHeight: "120px",
								       textAlign: "left",
								       verticalAlign: "top",
							       }}
							       placeholder="Enter user subject ids separated with comma and no whitespace (e.g. 1,2,3). List should contain all subjects, not just new ones."
							       id="subject-id-input"
							/>

						</div>
						: ""}

					<button className="btn btn-primary"
					        style={{
						        padding: "4px 16px",
						        fontSize: "16px",
						        marginBottom: "16px"
					        }}
					        onClick={() => {
								if (editSubjects) {
									// save changes
									let subjectIdElement = document.getElementById("subject-id-input") as HTMLInputElement;
									let subjectIds = subjectIdElement ? subjectIdElement.value : '';

									// send request
									setUserSubjects(
										generalData.username,
										subjectIds
									)

									setEditSubjects(false);
									setTimeout(() => {
										window.location.reload();
									}, 500);
								} else {
									setEditSubjects(true);
								}
					        }}
					        >
						Edit subjects
					</button>

				</>
			)}

			<h4>
			</h4>

		</div>
	);
}

export default UserPage;