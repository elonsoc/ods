/** @type {import('next').NextConfig} */
const nextConfig = {
	experimental: {
		appDir: true,
	},
	output: 'standalone',
	async redirects() {
		return [
			{
				source: '/saml/acs',
				destination: 'http://api.ods.elon.edu/saml/acs',
				permanent: true,
				basePath: false,
			},
		];
	},
};

module.exports = nextConfig;
