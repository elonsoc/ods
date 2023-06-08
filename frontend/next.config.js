/** @type {import('next').NextConfig} */
const nextConfig = {
	experimental: {
		appDir: true,
	},
	output: 'standalone',
	async redirects() {
		return [
			{
				source: '/docs/getting-started',
				destination: '/docs/getting-started/overview',
				permanent: true,
			},
		];
	},
};

module.exports = nextConfig;
