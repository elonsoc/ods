import Navbar from '@/components/Navbar/Navbar';
import Footer from '@/components/Footer/Footer';
import { Inter } from 'next/font/google';
import styles from '@/styles/layout.module.css';
import '../styles/globals.css';
import { Metadata } from 'next';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
	title: 'Elon ODS',
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
		<html lang='en'>
			<head></head>
			<body>
				<div className={styles.container}>
					<Navbar />
					<main className={inter.className}>{children}</main>
					<Footer />
				</div>
			</body>
		</html>
	);
}
