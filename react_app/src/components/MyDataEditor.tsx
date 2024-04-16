import React from "react";
import '../styles/mark_editor.css';

interface MyDataEditorProperties {
	isAdmin: boolean;
}

const MyDataEditor: React.FC<MyDataEditorProperties> = ({isAdmin}) => {
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

				{!isAdmin && (
					<tr>
						<th>
							Old password
						</th>
						<td>
							<input type="password" id="old-password-input"/>
						</td>
					</tr>
				)}

				{isAdmin && (
					<>
						<tr>
							<th>
								New name
							</th>
							<td>
								<input type="text" id="new-name-input"/>
							</td>
						</tr>
						<tr>
							<th>
								New surname
							</th>
							<td>
								<input type="text" id="new-surname-input"/>
							</td>
						</tr>
					</>
				)}

				<tr>
					<th style={{paddingRight: "16px"}}>
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