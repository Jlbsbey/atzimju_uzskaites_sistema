import React from "react";
import '../styles/subject_menu_component_item.css';

interface SubjectsMenuItemProperties {
	name: string;
	description: string;
	last_updated: string;
	professor: string;
	is_active?: boolean;
	on_click: () => void;
}

const SubjectsMenuItem: React.FC<SubjectsMenuItemProperties> = ({
	name,
	description,
	last_updated,
	professor,
	is_active = false,
	on_click,
}) => {
	return (
		<a
			onClick={on_click}
		    className={
				"subject_menu_component_item list-group-item list-group-item-action" +
				(is_active ? " active" : "")
			}
		>
			<div className="d-flex w-100 justify-content-between">
				<h5 className="mb-1">
					{name}
				</h5>
				<small>
					{last_updated + " day(s) ago"}
				</small>
			</div>
			<p className="mb-1">
				{description}
			</p>
			<small>
				{professor}
			</small>
		</a>
	)
};

export default SubjectsMenuItem;