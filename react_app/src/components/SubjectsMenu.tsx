import React, {useEffect, useState} from "react";
import SubjectsMenuItem from "./SubjectsMenuItem";
import {Subject} from "../scripts/data";
import {getHomeData} from "../scripts/home";

interface SubjectsMenuProperties {
	onSubjectClick: (subject: Subject) => void;
	subjects: Subject[];
}

const SubjectsMenu: React.FC<SubjectsMenuProperties> = ({onSubjectClick, subjects}) => {
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
						name={subject.subject_name}
						description={subject.subject_description}
						is_active={subject.subject_id === activeSubject.subject_id}
						on_click={() => activeSubjectIdChange(subject)}
					/>
				))}

			</div>
		</>
	)
};

export default SubjectsMenu;