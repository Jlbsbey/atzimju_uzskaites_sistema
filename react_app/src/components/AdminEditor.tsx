import React from "react";
import '../styles/mark_editor.css';

interface AdminEditorProperties {
	mode: string;
}

const AdminEditor: React.FC<AdminEditorProperties> = ({
	mode
}) => {
	return (
		<div className="overlay_mark">
			<h3>
				Add {mode.slice(0, -1)}
			</h3>

			<table>
				{mode === "subjects" && (
					<>
						<tr>
							<td>
								Subject name
							</td>
							<td>
								<input id="subject-input" type="text"/>
							</td>
						</tr>
						<tr>
							<td style={{paddingRight: "20px"}}>
								Subject description
							</td>
							<td>
								<input id="description-input" type="text"/>
							</td>
						</tr>
					</>
				)}
				{(mode === "students" || mode === "professors") && (
					<>
						<tr>
							<td>
								Name
							</td>
                            <td>
                                <input id="name-input" type="text"/>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                Surname
                            </td>
                            <td>
                                <input id="surname-input" type="text"/>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                Email
                            </td>
                            <td>
                                <input id="email-input" type="text"/>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                Password
                            </td>
                            <td>
                                <input id="password-input" type="password"/>
                            </td>
                        </tr>
                        <tr>
                            <td style={{paddingRight: "20px"}}>
                                Avatar link
                            </td>
                            <td>
                                <input id="avatar-link-input" type="text"/>
                            </td>
                        </tr>
                    </>
				)}


			</table>

		</div>
	)
};

export default AdminEditor