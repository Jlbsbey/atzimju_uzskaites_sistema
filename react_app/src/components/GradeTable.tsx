import React from "react";
import '../styles/grade_table.css';
import {Grade} from "../scripts/data";

interface GradeTableProperties {
	grades: Grade[];
}

const GradeTable: React.FC<GradeTableProperties> = ({
	grades
}) => {
	return (
		<table className="grade-table table table-sm table-striped">
			<thead>
				<tr>
					<th className="grade-table-t-first">Grade</th>
					<th>Professor</th>
					<th>Created</th>
					<th>Last Updated</th>
				</tr>
			</thead>

			<tbody className="table-group-divider">
				{grades.map((grade) => (
					<tr key={grade.id}>
						<td className="grade-table-t-first">{grade.value}</td>
						<td>{grade.professor}</td>
						<td>{grade.created_unix}</td>
						<td>{grade.last_updated_unix}</td>
					</tr>
				))}
			</tbody>
		</table>
	)
};

export default GradeTable;