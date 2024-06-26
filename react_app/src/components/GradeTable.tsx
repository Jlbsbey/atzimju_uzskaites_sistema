import React, {useEffect} from "react";
import '../styles/grade_table.css';
import {Mark} from "../scripts/data";

interface GradeTableProperties {
	marks: Mark[];
	mode: string;
	overlayEditMarkActive: number;
	setOverlayAddMarkActive: (active: boolean) => void;
	setOverlayEditMarkActive: (active: number) => void;
}

const GradeTable: React.FC<GradeTableProperties> = ({
	marks,
	mode,
	overlayEditMarkActive,
	setOverlayAddMarkActive,
	setOverlayEditMarkActive
}) => {
	let studentMode = mode == "student";

	let [trs, setTrs] = React.useState<JSX.Element[]>([]);

	useEffect(() => {
		let trs = marks.map((mark) => (
			<tr key={mark.mark_id}>
				<td className="grade-table-t-first">{mark.value}</td>
				<td>
					<a href={`/user?id=${studentMode ? mark.professor_id :  mark.student_id }`}>
						{studentMode ? mark.professor_name : mark.student_name}
					</a>
				</td>
				<td>{mark.create_date}</td>
				<td>{mark.edit_date !== mark.create_date ? mark.edit_date : "-"}</td>
				{!studentMode && (
					<td>
						<button className="btn btn-primary"
						        style={{
							        padding: "0 12px",
							        fontSize: "12px",
							        fontWeight: "bold"
						        }}
						        onClick={() => {
							        setOverlayAddMarkActive(false);
									console.log("mark id:" + mark.mark_id)
							        setOverlayEditMarkActive(mark.mark_id);
						        }}
						>
							Edit
						</button>
					</td>
				)}
			</tr>
		));

		setTrs(trs);
	}, [marks]);

	return (
		<table className="grade-table table table-sm table-striped">
			<thead>
				<tr>
					<th className="grade-table-t-first">Grade</th>
					<th>{studentMode ? "Professor" : "Student"}</th>
					<th>Created</th>
					<th>Last Updated</th>
					{!studentMode && <th>Edit</th>}
				</tr>
			</thead>

			<tbody className="table-group-divider">
				{trs}
			</tbody>
		</table>
	)
};

export default GradeTable;