/* this is a bit confusing—API_URL and BACKEND_API_URL doesn't make much sense.

API_URL is just the URL that the service exists. With nextjs, our frontend calls
<url>/api/* for api calls in the backend. The "backend api url" is used exclusively on the frontend's backend
and, honestly, we can remove it entirely from the frontend by making it an api call.

The reason I'm a bit confused is because we can call the frontend's api using just /api/* and not require
the full api_url.
*/
interface Config {
	url: {
		API_URL: string;
		BACKEND_API_URL: string;
	};
}
const prod: Config = {
	url: {
		API_URL: `https://ods.elon.edu`,
		BACKEND_API_URL:
			process.env.NEXT_PUBLIC_BACKEND_API_URL || 'https://api.ods.elon.edu',
	},
};
const dev: Config = {
	url: {
		API_URL: `http://localhost:3001`,
		BACKEND_API_URL: 'http://localhost:3000',
	},
};

export const configuration = process.env.NODE_ENV == `development` ? dev : prod;
