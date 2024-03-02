import { NextResponse } from 'next/server';
import { redirect } from 'next/navigation';
import { config } from '@/config/Constants';
import { cookies } from 'next/headers';

const BACKEND_URL = config.url.BACKEND_API_URL;

async function fetchWithAutoRefresh(url: string, options: RequestInit) {
	let response = await fetch(url, options);
	if (response.status === 401) {
		const refreshSuccessful = await refreshToken();
		if (!refreshSuccessful) {
			console.error('Failed to refresh token or no refresh token available.');
			return null;
		}
		response = await fetch(url, options);
	}
	return response;
}

async function refreshToken() {
	const refreshToken = cookies().get('ods_refresh_cookie_nomnom')?.value;
	if (!refreshToken) {
		return false;
	}

	const requestOptions = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'X-Refresh-Token': refreshToken,
		},
	};

	const response = await fetch(`${BACKEND_URL}/refresh`, requestOptions);
	return response.ok;
}

export async function GET() {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	if (!login_cookie) {
		return new NextResponse();
	}

	const requestOptions: RequestInit = {
		cache: 'no-store',
		headers: { Cookie: `${login_cookie.name}=${login_cookie.value}` },
	};

	const response = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/`,
		requestOptions
	);
	if (!response) {
		redirect('/login');
	}

	const applicationJSON = await parsePotentialJSON(response);
	return NextResponse.json(applicationJSON);
}

export async function POST(request: Request) {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	if (!login_cookie) {
		return new NextResponse();
	}

	const options: any = {
		method: 'POST',
		duplex: 'half',
		headers: {
			'Content-Type': 'application/json',
			credentials: 'include',
			Cookie: `${login_cookie.name}=${login_cookie.value}`,
		},
		body: request.body,
		cache: 'no-store',
	};

	const response = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/`,
		options
	);
	if (!response) {
		redirect('/login');
	}

	const data = await parsePotentialJSON(response);
	return NextResponse.json(data);
}

async function parsePotentialJSON(res: Response) {
	try {
		return await res.json();
	} catch (error) {
		console.error('Failed to parse JSON:', error);
		return res.body;
	}
}
