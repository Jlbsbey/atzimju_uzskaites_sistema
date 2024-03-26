import Cookies from 'js-cookie';

export const setAuthCookie = (authKey: string, expiryTime: string) => {
	let formattedExpiryTime = new Date(expiryTime);
	Cookies.set('auth_key', authKey, { expires: formattedExpiryTime });
};

export const getAuthCookie = () => {
	return Cookies.get('auth_key');
};

export const removeAuthCookie = () => {
	Cookies.remove('auth_key');
};
