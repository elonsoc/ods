import { Metadata } from 'next';
import Sidebar from './_components/Sidebar/Sidebar';
import styles from '@/styles/pages/docs/docs.module.css';
import Breadcrumbs from './_components/Breadcrumbs/Breadcrumbs';

export const metadata: Metadata = {
	title: 'Docs',
};

export default function Layout({ children }: { children: React.ReactNode }) {
	return (
		<div className={styles.docsContainer}>
			<Sidebar />
			<div className={styles.docsMain}>{children}</div>
		</div>
	);
}
