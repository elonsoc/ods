import { NextResponse } from 'next/server';
import { configuration } from '@/config/Constants';
import { cookies } from 'next/headers';
const BACKEND_URL = configuration.url.BACKEND_API_URL;

export async function GET(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<NextResponse> {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	const id = params.id;
	let applicationJSON;
	try {
		const res = await fetch(`${BACKEND_URL}/applications/${id}`, {
			method: 'GET',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json',
				Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
			},
			cache: 'no-cache',
		});
		applicationJSON = await res.json();
	} catch (error) {
		console.log('There was an error', error);
		return NextResponse.json({ error: error });
	}
	return NextResponse.json(applicationJSON);
}

export async function PUT(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<Response> {
	const login_cookie = cookies().get('ods_login_cookie_nomnom');
	const id = params.id;
	const options: any = {
		method: 'PUT',
		duplex: 'half',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			Cookie: `${login_cookie?.name}=${login_cookie?.value}`,
		},
		body: request.body,
		cache: 'no-store',
	};
	const res = await fetch(`${BACKEND_URL}/applications/${id}`, options);
	return new NextResponse();
}

export async function DELETE(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<Response> {
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

	const res = await fetch(`${BACKEND_URL}/applications/${id}`, options);
	return new NextResponse();
}
