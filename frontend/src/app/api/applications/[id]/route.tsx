import { NextResponse } from 'next/server';
import { UserAppInformation } from '../application.d';
import applications from '../data.json';

export async function GET(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<NextResponse> {
	const id = params.id;
	let application = applications.filter((app) => app.id === parseInt(id));
	return NextResponse.json(application);
}

export async function PUT(
	request: Request,
	{
		params,
	}: {
		params: { id: string };
	}
): Promise<NextResponse> {
	const body: UserAppInformation = await request.json();
	const id = params.id;
	for (let i = 0; i < applications.length; i++) {
		if (applications[i].id.toString() === id) {
			applications[i] = { ...applications[i], ...body };
		}
	}

	return NextResponse.json(applications);
}
