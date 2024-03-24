import bcrypt from 'bcryptjs';
import {getAuthCookie} from "./cookies";

export const encryptRequest = (request: string): string => {
	const authCookie = getAuthCookie();
	if (!authCookie) {
		return "";
	}

	const base= request + authCookie;
	return hashString(base);
}

const hashString = (input: string): string => {
	const saltRounds = 10; // Number of salt rounds for hashing
	return bcrypt.hashSync(input, saltRounds);
};
