import { UserAppInformation } from '@/app/apps/_components/UserApp/UserApp';
import type { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(
	req: NextApiRequest,
	res: NextApiResponse<UserAppInformation>
) {
	if (req.method !== 'POST') {
		res.status(405);
		return;
	}

	const endpoint = 'http://localhost:5319/applications';

	const options = {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({ help: 'me' }),
	};

	const data = await fetch(endpoint, options);
	const dataJSON = await data.json();

	res.status(200);
}
