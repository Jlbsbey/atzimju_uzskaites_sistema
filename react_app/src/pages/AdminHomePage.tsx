import React, {useEffect, useState} from "react";
import "../styles/main_page.css";
import {Mark, Subject, User} from "../scripts/data";
import SimpleBar from 'simplebar-react';
import 'simplebar-react/dist/simplebar.min.css';
import SubjectsMenuItem from "../components/SubjectsMenuItem";
import AdminEditor from "../components/AdminEditor";
import {getAdminData} from "../scripts/admin";
import {addUser} from "../scripts/addUser";
import {addSubject} from "../scripts/addSubject";

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
	const [users, setUsers] = useState<User[]>(
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
			setUsers(data.content.users);
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
										        } else {
											        setOverlayButtonMode("subjects");
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
												        "professor"
											        )

											        loadData();

											        setOverlayButtonMode("");
										        } else {
											        setOverlayButtonMode("students");
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
										        }
									        }
								        }}
								>
									Add {adminMode.slice(0, -1)}
								</button>
							)}
						</>
					)}

				</div>
			</div>

		</>
	);
}

export default AdminHomePage;