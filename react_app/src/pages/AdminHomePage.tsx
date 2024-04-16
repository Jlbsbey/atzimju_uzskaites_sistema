import React, {useEffect, useState} from "react";
import "../styles/main_page.css";
import {Subject, User} from "../scripts/data";
import SimpleBar from 'simplebar-react';
import 'simplebar-react/dist/simplebar.min.css';
import SubjectsMenuItem from "../components/SubjectsMenuItem";
import AdminEditor from "../components/AdminEditor";
import {getAdminData} from "../scripts/admin";
import {addUser} from "../scripts/addUser";
import {addSubject} from "../scripts/addSubject";
import internal from "node:stream";

const AdminHomePage: React.FC = () => {
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
	const [students, setStundets] = useState<User[]>(
		[
			{
				user_id: 0,
				username: ""
			}
		]
	);
	const [professors, setProfessors] = useState<User[]>(
		[
			{
				user_id: 0,
				username: ""
			}
		]
	);
	const [generalData, setGeneralData] = useState({
		loaded: false,
		user_id: 0,
		role: ""
	});

	function loadData() {
		getAdminData().then((data) => {
			setSubjects(data.content.subjects);
			setStundets(data.content.students);
			setProfessors(data.content.professors);
			console.log(students);
			setGeneralData(
				{
					loaded: true,
					user_id: data.content.user_id,
					role: data.content.role
				})
		});
	}

	const [adminMode, setAdminMode] = useState("subjects");
	const [overlayButtonMode, setOverlayButtonMode] = useState("");
	const [editSubject, setEditSubject] = useState<number>(0);

	useEffect(() => {
		loadData();
	}, [overlayButtonMode]);

	return (
		<>
			<div className="main-page">

				<div className="left-menu">
					<SimpleBar
						style={{maxHeight: "100%", height: "100%", width: "100%"}}
					>
						<div className="subject_menu_component list-group">
							<SubjectsMenuItem
								name="Subjects"
								description="View and edit subjects."
								is_active={adminMode === "subjects"}
								on_click={() => setAdminMode("subjects")}
							/>
							<SubjectsMenuItem
								name="Students"
								description="View and edit students."
								is_active={adminMode === "students"}
								on_click={() => setAdminMode("students")}
							/>
							<SubjectsMenuItem
								name="Professors"
								description="View and edit professors."
								is_active={adminMode === "professors"}
								on_click={() => setAdminMode("professors")}
							/>
						</div>
					</SimpleBar>
				</div>

				<div className="right-content">

				<h1>
						{
							adminMode.charAt(0).toUpperCase() + adminMode.slice(1)
						}
					</h1>

					{generalData.loaded && (
						<>

							{overlayButtonMode !== "" && (
								<AdminEditor
									mode={overlayButtonMode}
								/>
							)}

							{(
								<button className="btn btn-primary"
								        style={{
									        padding: "4px 16px",
									        fontSize: "16px",
									        marginBottom: "16px"
								        }}
								        onClick={() => {
									        if (adminMode === "subjects") {

												if (editSubject !== -1) {
													if (overlayButtonMode === "subjects") {
														// submit add
														let subjectElement = document.getElementById("subject-input") as HTMLInputElement;
														let descriptionElement = document.getElementById("description-input") as HTMLInputElement;
														let subjectString = subjectElement ? subjectElement.value : '';
														let descriptionString = descriptionElement ? descriptionElement.value : '';

														addSubject(
															subjectString,
															descriptionString
														)

														loadData();

														setOverlayButtonMode("");
														setEditSubject(-1);
													} else {
														setOverlayButtonMode("subjects");
														setEditSubject(-1);
													}
												} else {
													let subjectElement = document.getElementById("subject-input") as HTMLInputElement;
													let descriptionElement = document.getElementById("description-input") as HTMLInputElement;
													let subjectString = subjectElement ? subjectElement.value : '';
													let descriptionString = descriptionElement ? descriptionElement.value : '';

													// submit

													setOverlayButtonMode("");
													setEditSubject(-1);
												}


									        } else if (adminMode === "students") {
										        if (overlayButtonMode === "students") {
											        // submit add
											        let nameElement = document.getElementById("name-input") as HTMLInputElement;
											        let surnameElement = document.getElementById("surname-input") as HTMLInputElement;
											        let emailElement = document.getElementById("email-input") as HTMLInputElement;
											        let passwordElement = document.getElementById("password-input") as HTMLInputElement;
											        let avatarLinkElement = document.getElementById("avatar-link-input") as HTMLInputElement;

											        let nameString = nameElement ? nameElement.value : '';
											        let surnameString = surnameElement ? surnameElement.value : '';
											        let emailString = emailElement ? emailElement.value : '';
											        let passwordString = passwordElement ? passwordElement.value : '';
											        let avatarLinkString = avatarLinkElement ? avatarLinkElement.value : '';

											        addUser(
												        nameString,
												        surnameString,
												        emailString,
												        passwordString,
												        avatarLinkString,
												        "student"
											        )

											        loadData();

											        setOverlayButtonMode("");
										        } else {
											        setOverlayButtonMode("students");
											        setEditSubject(-1);
										        }
									        } else if (adminMode === "professors") {
										        if (overlayButtonMode === "professors") {
											        // submit add
											        let nameElement = document.getElementById("name-input") as HTMLInputElement;
											        let surnameElement = document.getElementById("surname-input") as HTMLInputElement;
											        let emailElement = document.getElementById("email-input") as HTMLInputElement;
											        let passwordElement = document.getElementById("password-input") as HTMLInputElement;
											        let avatarLinkElement = document.getElementById("avatar-link-input") as HTMLInputElement;

											        let nameString = nameElement ? nameElement.value : '';
											        let surnameString = surnameElement ? surnameElement.value : '';
											        let emailString = emailElement ? emailElement.value : '';
											        let passwordString = passwordElement ? passwordElement.value : '';
											        let avatarLinkString = avatarLinkElement ? avatarLinkElement.value : '';

											        addUser(
												        nameString,
												        surnameString,
												        emailString,
												        passwordString,
												        avatarLinkString,
												        "professor"
											        )
											        loadData();

											        setOverlayButtonMode("");
										        } else {
											        setOverlayButtonMode("professors");
											        setEditSubject(-1);
										        }
									        }
								        }}
								>
									Add {adminMode.slice(0, -1)}
								</button>
							)}

							<table className="grade-table table table-sm table-striped"
							       style={{minWidth: "600px"}}
							       >
								<thead>
								<tr>
									{adminMode === "subjects" && (
										<>
											<th className="grade-table-t-first">
												Id
											</th>
											<th>
												{adminMode.charAt(0).toUpperCase() + adminMode.slice(1, -1)}
											</th>
											<th>
												Description
											</th>
										</>
									)}

									{adminMode === "students" && (
										<>
											<th className="grade-table-t-first">
												{adminMode.charAt(0).toUpperCase() + adminMode.slice(1, -1)}
											</th>
										</>
									)}

									{adminMode === "professors" && (
										<>
											<th className="grade-table-t-first">
												{adminMode.charAt(0).toUpperCase() + adminMode.slice(1, -1)}
											</th>
										</>
									)}
								</tr>
								</thead>
								<tbody className="table-group-divider">
								{adminMode === "subjects" && (
									subjects?.map((subject) => (
										<tr key={subject.subject_id}>
											<td className="grade-table-t-first">
												{subject.subject_id}
											</td>
											<td>
												{subject.subject_name}
											</td>
											<td>
												{subject.subject_description}
											</td>
										</tr>
									)))}

								{adminMode === "students" && (
									students?.map((student) => (
										<tr key={student.user_id}>
										<td className="grade-table-t-first">
											<a href={"/user?id=" + student.user_id}>
												{student.name}
											</a>
										</td>
									</tr>
								)))}

								{adminMode === "professors" && (
									professors?.map((professor) => (
										<tr key={professor.user_id}>
											<td className="grade-table-t-first">
												<a href={"/user?id=" + professor.user_id}>
													{professor.name}
												</a>
											</td>
										</tr>
									)))}
								</tbody>
							</table>
						</>
					)}

				</div>
			</div>

		</>
	);
}

export default AdminHomePage;