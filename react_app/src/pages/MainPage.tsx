import React, {useState} from "react";
import "../styles/main_page.css";
import SubjectsMenu from "../components/SubjectsMenu";
import {grades, subjects} from "../scripts/data";
import GradeTable from "../components/GradeTable";

const MainPage: React.FC = () => {
	const [activeSubject, setActiveSubject] = useState(subjects[0]);

	return (
		<div className="main-page">

			<div className="left-menu">
				<SubjectsMenu
					onSubjectClick={(subject) => setActiveSubject(subject)}
					/>
			</div>

			<div className="right-content">

				<h1>{activeSubject.name}</h1>

				<p style={{width:"50%"}}>
					{   activeSubject.name + " — " +
						activeSubject.description +
						" It is read by " + activeSubject.professor +
						" and last updated " + activeSubject.last_updated + " day(s) ago."
					}
				</p>

				<GradeTable grades={
					grades.filter((grade) => grade.subject_id === activeSubject.subject_id)
				}/>

			</div>

		</div>
	);
}

export default MainPage;