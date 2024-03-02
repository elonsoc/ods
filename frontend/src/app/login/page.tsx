'use client';
import { useEffect } from 'react';
import { config } from '@/config/Constants';

const BACKEND_URL = config.url.BACKEND_API_URL;

const Login = () => {
	useEffect(() => {
		window.location.replace(BACKEND_URL);
	});
	return null;
};

export default Login;
