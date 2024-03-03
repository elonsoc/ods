import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
	const accessToken = request.cookies.get('ods_login_cookie_nomnom');
	const refreshToken = request.cookies.get('ods_refresh_cookie_nomnom');
	if (!accessToken && !refreshToken) {
		return NextResponse.redirect(new URL('/login', request.url));
	}
}

export const config = {
	matcher: ['/apps/:path*', '/apps', '/api/applications/:function*'],
};
