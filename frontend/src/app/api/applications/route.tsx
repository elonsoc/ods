import { NextResponse } from 'next/server';
import { UserAppInformation } from './application.d';
import applications from './data.json';
import { config } from '@/config/Constants';
const BACKEND_URL = config.url.BACKEND_API_URL;

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
	console.log(`Fetching to url ${BACKEND_URL}/applications`);
	try {
		const res = await fetch(`${BACKEND_URL}/applications`);
		applicationJSON = await res.json();
	} catch (error) {
		console.log('oops, an error:', error)
		return new NextResponse()
	}
	console.log('applicationJSON', applicationJSON);
	return NextResponse.json(applicationJSON)
}

export async function POST(request: Request): Promise<NextResponse> {

	const options = {
		method: 'POST',
		duplex: 'half',
		headers: {
			'Content-Type': 'application/json',
			credentials: 'include',
		},
		body: request.body, // parsed automatically to an object as of Next.js v12
	};

	const res = await fetch(`${BACKEND_URL}/applications`, options);
	const data = await res.json();
	console.log(data)
	return NextResponse.json(data)
}
