import { NextResponse } from 'next/server';
import { RequestOptions, UserAppInformation } from './application.d';
import { configuration } from '@/config/Constants';
const BACKEND_URL = configuration.url.BACKEND_API_URL;

import { cookies } from 'next/headers';

export async function GET(request: Request): Promise<NextResponse> {
	let applicationJSON;
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	try {
		const res = await fetch(`${BACKEND_URL}/applications/`, {
			cache: 'no-store',
			credentials: 'include',
			headers: {
				Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
			},
		});
		applicationJSON = await res.json();
	} catch (error) {
		console.log(error);
		return new NextResponse();
	}
	return NextResponse.json(applicationJSON);
}

export async function POST(request: Request): Promise<NextResponse> {
	try {
		const login_cookie = cookies().get('ods_login_cookie_nomnom');
		const options: RequestOptions | any = {
			method: 'POST',
			duplex: 'half',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
			},
			body: request.body,
			cache: 'no-store',
		};
		const res = await fetch(`${BACKEND_URL}/applications/`, options);
		const data = await res.json();
		return NextResponse.json(data);
	} catch (error) {
		console.log(error);
		return new NextResponse();
	}
}
