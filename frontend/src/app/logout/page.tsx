'use client';
import { useEffect } from 'react';
import { configuration } from '@/config/Constants';
const BACKEND_URL = configuration.url.BACKEND_API_URL;

const Logout = () => {
	useEffect(() => {
		window.location.replace(BACKEND_URL + '/logout');
	});
	return null;
};

export default Logout;
