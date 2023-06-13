import { Metadata } from 'next';
import Sidebar from './_components/Sidebar/Sidebar';
import styles from '@/styles/pages/docs/docs.module.css';
import Breadcrumbs from './_components/Breadcrumbs/Breadcrumbs';
import TableOfContents from './_components/TableOfContents/TableOfContents';

export const metadata: Metadata = {
	title: 'Docs',
	description: 'Documentation for the Open Data Service API',
};

export default function Layout({ children }: { children: React.ReactNode }) {
	return (
		<div className={styles.docsContainer}>
			<Sidebar />
			<div className={styles.docsMain}>
				<Breadcrumbs />
				{children}
				<TableOfContents />
			</div>
		</div>
	);
}
