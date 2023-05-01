import { NextResponse } from 'next/server';
import { UserAppInformation } from './application.d';
import applications from './data.json';

export async function POST(request: Request): Promise<NextResponse> {
	const body: UserAppInformation = await request.json();
	applications.push({
		id: applications.length ? applications[applications.length - 1].id + 1 : 1,
		...body,
	});
	return NextResponse.json(body);
}

export async function GET(request: Request): Promise<NextResponse> {
	const { searchParams } = new URL(request.url);
	const id = searchParams.get('id');
	if (id) {
		let application = applications.filter((app) => app.id === parseInt(id));
		return NextResponse.json(application);
	}

	return NextResponse.json(applications);
}

// Possible GET code for our backend
// ----
// const { searchParams } = new URL(request.url);
// const id = searchParams.get('id');
// const res = await fetch(`http://localhost:1337/applications/${id}`, {
// 	headers: {
// 		'Content-Type': 'application/json',
//      'credentials': 'include',
// 	},
// });
// const application = await res.json();

// return NextResponse.json({ application });

// POST
// ----
// const endpoint = 'http://localhost:1337/applications';

// const options = {
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/json',
//      'credentials': 'include',
// 	},
// 	body: request.body, // parsed automatically to an object as of Next.js v12
// };

// const res = await fetch(endpoint, options);
// const data = await res.json();
// return NextResponse.json(data);