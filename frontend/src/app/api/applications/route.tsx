import { NextResponse } from 'next/server';
import { UserAppInformation } from './application.d';
import applications from './data.json';
import { config } from '@/config/Constants';
const CURRENT_URL = config.url.BACKEND_API_URL;

// ---------MOCK
// export async function POST(request: Request): Promise<NextResponse> {
// 	const body: UserAppInformation = await request.json();
// 	applications.push({
// 		id: applications.length ? applications[applications.length - 1].id + 1 : 1,
// 		...body,
// 	});
// 	return NextResponse.json(body);
// }

// export async function GET(request: Request): Promise<NextResponse> {
// 	const { searchParams } = new URL(request.url);
// 	const id = searchParams.get('id');
// 	if (id) {
// 		let application = applications.filter((app) => app.id === parseInt(id));
// 		return NextResponse.json(application);
// 	}

// 	return NextResponse.json(applications);
// }

// ----

export async function GET(request: Request): Promise<NextResponse> {
	let applicationJSON;
	try {
		const res = await fetch(`${CURRENT_URL}/applications`, {
			headers: {
				'Content-Type': 'application/json',
				credentials: 'include',
			},
		});
		applicationJSON = await res.json();
	} catch (error) {
		console.log('There was an error', error);
	}
	return NextResponse.json(applicationJSON || []);
}

export async function POST(request: Request): Promise<NextResponse> {
	const endpoint = `${CURRENT_URL}/applications`;

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			credentials: 'include',
		},
		body: request.body, // parsed automatically to an object as of Next.js v12
	};

	const res = await fetch(endpoint, options);
	const data = await res.json();
	return NextResponse.json(data);
}
