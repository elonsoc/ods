import { NextResponse } from 'next/server';
import { UserAppInformation } from '../application.d';
import applications from '../data.json';
import { config } from '@/config/Constants';
import { cookies } from 'next/headers';
const BACKEND_URL = config.url.BACKEND_API_URL;

// --------- MOCK
// export async function GET(
// 	request: Request,
// 	{
// 		params,
// 	}: {
// 		params: { id: string };
// 	}
// ): Promise<NextResponse> {
// 	const id = params.id;
// 	let application = applications.filter((app) => app.id === parseInt(id));
// 	return NextResponse.json(application);
// }

// export async function PUT(
// 	request: Request,
// 	{
// 		params,
// 	}: {
// 		params: { id: string };
// 	}
// ): Promise<NextResponse> {
// 	const body: UserAppInformation = await request.json();
// 	const id = params.id;
// 	for (let i = 0; i < applications.length; i++) {
// 		if (applications[i].id.toString() === id) {
// 			applications[i] = { ...applications[i], ...body };
// 		}
// 	}

// 	return NextResponse.json(applications);
// }

// -------------

export async function GET(
  request: Request,
  {
    params,
  }: {
    params: { id: string };
  }
): Promise<NextResponse> {

  const login_cookie = cookies().get("ods_login_cookie_nomnom")
  if (login_cookie === undefined) {
    // this happens when, somehow, they're here yet they don't have
    // a login cookie :)
    return new NextResponse();
  }

  const id = params.id;
  let applicationJSON;
  try {
    const res = await fetch(`${BACKEND_URL}/applications/${id}`, {
      headers: {
        'Content-Type': 'application/json',
        credentials: 'include',
        Cookie: `${login_cookie.name}=${login_cookie.value}`
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
  const login_cookie = cookies().get("ods_login_cookie_nomnom")
  if (login_cookie === undefined) {
    // this happens when, somehow, they're here yet they don't have
    // a login cookie :)
    return new NextResponse();
  }

  const id = params.id;
  const options: any = {
    method: 'PUT',
    duplex: 'half',
    headers: {
      'Content-Type': 'application/json',
      credentials: 'include',
      Cookie: `${login_cookie.name}=${login_cookie.value}`
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

  const login_cookie = cookies().get("ods_login_cookie_nomnom")
  if (login_cookie === undefined) {
    // this happens when, somehow, they're here yet they don't have
    // a login cookie :)
    return new NextResponse();
  }

  const id = params.id;
  const options: any = {
    method: 'DELETE',
    duplex: 'half',
    headers: {
      'Content-Type': 'application/json',
      credentials: 'include',
      Cookie: `${login_cookie.name}=${login_cookie.value}`
    },
    body: request.body,
    cache: 'no-store',
  };

  const res = await fetch(`${BACKEND_URL}/applications/${id}`, options);

  return new NextResponse();

}
