import { NextResponse } from 'next/server';
import { configuration } from '@/config/Constants';
import { cookies } from 'next/headers';
import { fetchWithAutoRefresh } from '@/actions/token';

const BACKEND_URL = configuration.url.BACKEND_API_URL;

export async function GET(request: Request): Promise<NextResponse> {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	try {
		const options: RequestInit = {
			headers: {
				'Content-Type': 'application/json',
				Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
			},
			credentials: 'include',
			cache: 'no-store',
		};
		const res = await fetchWithAutoRefresh(
			`${BACKEND_URL}/login/status`,
			options
		);
		return NextResponse.json({ isAuthenticated: res.status === 200 });
	} catch (error) {
		console.error('There was an error', error);
		return NextResponse.json({ isAuthenticated: false });
	}
}
