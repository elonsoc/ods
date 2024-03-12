import { type NextRequest, NextResponse } from 'next/server';
import { configuration } from '@/config/Constants';
import { cookies } from 'next/headers';
import { fetchWithAutoRefresh } from '@/actions/token';

const BACKEND_URL = configuration.url.BACKEND_API_URL;

export async function GET(request: NextRequest) {
	const accessToken = cookies().get('ods_login_cookie_nomnom');

	const requestOptions: RequestInit = {
		method: 'GET',
		cache: 'no-store',
		credentials: 'include',
		headers: {
			Cookie: `${accessToken?.name}=${accessToken?.value}`,
			'Content-Type': 'application/json',
		},
	};

	const response = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/`,
		requestOptions
	);

	if (!response || response.status === 401) {
		return NextResponse.redirect(new URL('/login', request.url));
	}

	const applicationJSON = await response.json();
	return NextResponse.json(applicationJSON);
}

export async function POST(request: NextRequest) {
	const accessToken = cookies().get('ods_login_cookie_nomnom');
	const options: RequestInit | any = {
		method: 'POST',
		duplex: 'half',
		credentials: 'include',
		cache: 'no-store',
		headers: {
			'Content-Type': 'application/json',
			Cookie: `${accessToken?.name}=${accessToken?.value}`,
		},
		body: await request.json(),
	};

	const response = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/`,
		options
	);

	if (!response || response.status === 401) {
		return NextResponse.redirect(new URL('/login', request.url));
	}

	const data = await response.json();
	return NextResponse.json(data);
}
