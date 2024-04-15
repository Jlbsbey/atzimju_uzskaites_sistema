import React from "react";
import '../styles/mark_editor.css';

interface GradeTableProperties {
	student?: string;
	mode: "add" | "edit";
}

const MarkEditor: React.FC<GradeTableProperties> = ({
	                                                     student,
	                                                     mode
                                                    }) => {
	return (
		<div className="overlay_mark">
			<h2>
				{mode === "add" ? "Add mark" : "Edit mark"}
			</h2>

			<table>
				<tr>
					<th style={{paddingRight: "20px"}}>
						Student
					</th>
					<th>
						{mode === "edit" ? student :
							<input type="text"/>
						}
					</th>
				</tr>
				<tr>
					<td>
						Mark
					</td>
					<td>
						<input type="number" id="mark-input"/>
					</td>
				</tr>
			</table>

		</div>
	)
};

export default MarkEditor;