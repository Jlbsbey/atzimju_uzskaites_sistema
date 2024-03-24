import Cookies from 'js-cookie';

export const setAuthCookie = (authKey: string) => {
	const expiryTime = new Date();
	expiryTime.setTime(expiryTime.getTime() + 60 * 60 * 1000); // 60 minutes from now

	Cookies.set('auth_key', authKey, { expires: expiryTime });
};

export const getAuthCookie = () => {
	return Cookies.get('auth_key');
};

export const removeAuthCookie = () => {
	Cookies.remove('auth_key');
};
