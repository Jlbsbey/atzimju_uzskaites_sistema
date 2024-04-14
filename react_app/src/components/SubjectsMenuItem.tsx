import React from "react";
import '../styles/subject_menu_component_item.css';

interface SubjectsMenuItemProperties {
	name: string;
	description: string;
	is_active?: boolean;
	on_click: () => void;
}

const SubjectsMenuItem: React.FC<SubjectsMenuItemProperties> = ({
	name,
	description,
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
			</div>
			<p className="mb-1">
				{description}
			</p>
		</a>
	)
};

export default SubjectsMenuItem;