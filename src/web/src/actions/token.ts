'use server';

import { configuration } from '@/config/Constants';
import { cookies } from 'next/headers';

const BACKEND_URL = configuration.url.BACKEND_API_URL;

export async function fetchWithAutoRefresh(
	url: string,
	options: RequestInit | any
) {
	if (options.body && typeof options.body !== 'string') {
		options.body = JSON.stringify(options.body);
	}

	let response = await fetch(url, options);

	if (response.status === 401) {
		const tokens = await refreshTokens();
		options.headers = {
			...options.headers,
			Cookie: `ods_login_cookie_nomnom=${tokens?.access_token}`,
		};
		response = await fetch(url, options);
	}
	return response;
}

export async function refreshTokens() {
	const refreshToken = cookies().get('ods_refresh_cookie_nomnom')?.value;
	if (!refreshToken) {
		console.error('No refresh token available.');
		return null;
	}

	const requestOptions: RequestInit = {
		method: 'POST',
		cache: 'no-store',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			'X-Refresh-Token': refreshToken,
		},
	};

	const response = await fetch(`${BACKEND_URL}/refresh`, requestOptions);
	if (!response.ok) {
		console.error('Failed to refresh token:', response.statusText);
		return null;
	}

	const { access_token, refresh_token } = await response.json();

	cookies().set('ods_login_cookie_nomnom', access_token, {
		expires: new Date(new Date().getTime() + 5 * 60 * 1000),
		secure: true,
		httpOnly: true,
		path: '/',
	});
	cookies().set('ods_refresh_cookie_nomnom', refresh_token, {
		expires: new Date(new Date().getTime() + 7 * 24 * 60 * 60 * 1000),
		secure: true,
		httpOnly: true,
		path: '/',
	});

	return { access_token, refresh_token };
}
