import React, {useState} from "react";
import SubjectsMenuItem from "./SubjectsMenuItem";
import {Subject, subjects} from "../scripts/data";

interface SubjectsMenuProperties {
	onSubjectClick: (subject: Subject) => void;
}

const SubjectsMenu: React.FC<SubjectsMenuProperties> = ({onSubjectClick}) => {
	const [activeSubject, setActiveSubject] = useState(subjects[0]);

	const activeSubjectIdChange = (subject: Subject) => {
		setActiveSubject(subject);
		onSubjectClick(subject);
	}

	return (
		<>
			<div className="subject_menu_component list-group">

				{subjects.map((subject, index) => (
					<SubjectsMenuItem
						key={index}
						name={subject.name}
						description={subject.description}
						last_updated={subject.last_updated}
						professor={subject.professor}
						is_active={subject.subject_id === activeSubject.subject_id}
						on_click={() => activeSubjectIdChange(subject)}
					/>
				))}

			</div>
		</>
	)
};

export default SubjectsMenu;