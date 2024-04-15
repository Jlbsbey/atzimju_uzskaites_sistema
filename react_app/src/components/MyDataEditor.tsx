import React from "react";
import '../styles/mark_editor.css';

const MyDataEditor = () => {
	return (
		<div className="overlay_mark">
			<h2>
				Change my data
			</h2>

			<table>
				<tr>
					<th>
						New email
					</th>
					<td>
						<input type="text" id="new-email-input"/>
					</td>
				</tr>
				<tr>
					<th>
						Old password
					</th>
					<td>
						<input type="password" id="old-password-input"/>
					</td>
				</tr>
				<tr>
					<th>
						New password
					</th>
					<td>
						<input type="password" id="new-password-input"/>
					</td>
				</tr>
			</table>

		</div>
	)
};

export default MyDataEditor;