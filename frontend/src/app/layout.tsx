import Navbar from '@/components/ui/Navbar/Navbar';
import styles from '@/styles/pages/layout.module.css';
import '../styles/global/globals.css';
import { Metadata } from 'next';
import { Raleway } from 'next/font/google';
import { AuthProvider } from '@/context/auth/auth';

const raleway = Raleway({ subsets: ['latin'] });

export const metadata: Metadata = {
	title: {
		default: 'Open Data Service',
		template: '%s | ODS',
	},
	keywords:
		"'data access, api provider, Elon University, open source, open data service, ods, elon'",
	description:
		'Elon ODS is an Open Data Service at Elon University that provides API keys for students to access data about Elon University. Our API provider service offers data about buildings, courses, and more. Register an application today to get started!',
};

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<AuthProvider>
			<html lang='en'>
				<head></head>
				<body className={raleway.className}>
					<div className={styles.container}>
						<Navbar />
						<main>{children}</main>
						{/* <Footer /> */}
					</div>
				</body>
			</html>
		</AuthProvider>
	);
}
