import { type NextRequest, NextResponse } from 'next/server';
import { configuration } from '@/config/Constants';
import { cookies } from 'next/headers';
import { fetchWithAutoRefresh } from '@/actions/token';
const BACKEND_URL = configuration.url.BACKEND_API_URL;

export async function GET(
	request: NextRequest,
	{
		params,
	}: {
		params: { id: string };
	}
) {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	const id = params.id;

	const options: RequestInit = {
		method: 'GET',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
		},
		cache: 'no-cache',
	};
	const res = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/${id}`,
		options
	);

	if (!res || res.status === 401) {
		return NextResponse.redirect(new URL('/login', request.url));
	}

	const applicationJSON = await res.json();
	return NextResponse.json(applicationJSON);
}

export async function PUT(
	request: NextRequest,
	{
		params,
	}: {
		params: { id: string };
	}
) {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	const id = params.id;
	const options: RequestInit | any = {
		method: 'PUT',
		duplex: 'half',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
		},
		body: await request.json(),
		cache: 'no-store',
	};
	const res = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/${id}`,
		options
	);
	return NextResponse.json(res);
}

export async function DELETE(
	request: NextRequest,
	{
		params,
	}: {
		params: { id: string };
	}
) {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	const id = params.id;
	const options: any = {
		method: 'DELETE',
		duplex: 'half',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
		},
		body: request.body,
		cache: 'no-store',
	};
	const res = await fetchWithAutoRefresh(
		`${BACKEND_URL}/applications/${id}`,
		options
	);
	return NextResponse.json(res);
}
