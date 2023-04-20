import { NextResponse } from 'next/server';
import { UserAppInformation } from './application.d';

export async function POST(request: Request) {
	const body: UserAppInformation = await request.json();
	return NextResponse.json(body);
	// const endpoint = 'http://localhost:1337/applications';

	// const options = {
	// 	method: 'POST',
	// 	headers: {
	// 		'Content-Type': 'application/json',
	// 	},
	// 	body: request.body, // parsed automatically to an object as of Next.js v12
	// };

	// const res = await fetch(endpoint, options);
	// const data = await res.json();
	// return NextResponse.json(data);
}
