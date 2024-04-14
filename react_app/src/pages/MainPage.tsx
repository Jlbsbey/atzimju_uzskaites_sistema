import React, {useEffect, useState} from "react";
import "../styles/main_page.css";
import SubjectsMenu from "../components/SubjectsMenu";
import {Mark, Subject, User} from "../scripts/data";
import GradeTable from "../components/GradeTable";
import SimpleBar from 'simplebar-react';
import 'simplebar-react/dist/simplebar.min.css';
import {getHomeData} from "../scripts/home";
import MarkEditor from "../components/MarkEditor";

const MainPage: React.FC = () => {
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
	const [marks, setMarks] = useState<Mark[]>(
		[
			{
				mark_id: 0,
				subject_id: 0,
				professor_id: "",
				professor_name: "",
				student_id: 0,
				student_name: "",
				value: 0,
				create_date: "",
				edit_date: ""
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

	useEffect(() => {
		getHomeData().then((data) => {
			setSubjects(data.content.subjects);
			setActiveSubject(data.content.subjects[0]);
			setActiveMarks(data.content.marks.filter((mark: Mark) =>
				mark.subject_id === activeSubject.subject_id
			));
			setMarks(data.content.marks);
			setUsers(data.content.users);
			setGeneralData(
				{
					loaded: true,
					user_id: data.content.user_id,
					role: data.content.role
				}
			)
		});
	}, []);

	function changeActiveSubject(subject: Subject) {
		setActiveSubject(subject);
		let newActiveMarks = marks.filter((mark: Mark) =>
			mark.subject_id === subject.subject_id
		)
		setActiveMarks(newActiveMarks);
		console.log(newActiveMarks)
	}

	const [activeSubject, setActiveSubject] = useState(subjects[0]);
	const [activeMarks, setActiveMarks] = useState<Mark[]>([]);

	const [overlayAddMarkActive, setOverlayMarkActive] = useState(false);
	const [overlayEditMarkActive, setOverlayEditMarkActive] = useState(-1);

	return (
		<>
			<div className="main-page">

				<div className="left-menu">
					<SimpleBar
						style={{ maxHeight: "100%", height: "100%", width: "100%"}}
					>
						<SubjectsMenu
							onSubjectClick={(subject) => changeActiveSubject(subject)}
							subjects={subjects}
						/>
					</SimpleBar>
				</div>

				<div className="right-content">

					<h1>{activeSubject.subject_name}</h1>

				{generalData.loaded && (
					<>
						<p style={{width:"50%"}}>
							{   activeSubject.subject_name + " — " +
								activeSubject.subject_description
							}
						</p>

						{overlayAddMarkActive && (
							<MarkEditor
								mode="add"
							/>
						)}
						{overlayEditMarkActive !== -1 && (
							<MarkEditor
								mode="edit"
								student={marks.find((mark) => mark.mark_id === overlayEditMarkActive)?.student_name}
							/>
						)}

						{generalData.role === "professor" && (
							<button className="btn btn-primary"
							        style={{
										padding: "4px 16px",
								        fontSize: "16px",
								        marginBottom: "16px"
									}}
							        onClick={() => {
										if (overlayEditMarkActive === -1) {
											if (overlayAddMarkActive) {
												// submit add
												setOverlayMarkActive(false);
											} else {
												setOverlayMarkActive(true);
												setOverlayEditMarkActive(-1);
											}
										} else {
											// submit edit
											setOverlayEditMarkActive(-1);
										}
									}}
							>
								{overlayEditMarkActive === -1 ? "Add mark" : "Edit mark"}
							</button>
						)}

						<GradeTable marks={activeMarks}
						            mode={generalData.role}
						            overlayEditMarkActive={overlayEditMarkActive}
						            setOverlayAddMarkActive={setOverlayMarkActive}
						            setOverlayEditMarkActive={setOverlayEditMarkActive}
						/>
					</>
				)}

				</div>
			</div>

		</>
	);
}

export default MainPage;