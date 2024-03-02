import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
	let cookie = request.cookies.get('ods_login_cookie_nomnom');
	if (!cookie) {
		return NextResponse.redirect(new URL('/login', request.url));
	}
}

export const config = {
	matcher: ['/apps/:path*', '/apps', '/api/applications/:function*'],
};
