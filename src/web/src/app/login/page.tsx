'use client';
import { useEffect } from 'react';
import { configuration } from '@/config/Constants';
const BACKEND_URL = configuration.url.BACKEND_API_URL;

const Login = () => {
	useEffect(() => {
		window.location.replace(BACKEND_URL);
	});
	return null;
};

export default Login;
