'use client';

import Sidebar from './_components/Sidebar/Sidebar';
import styles from '@/styles/pages/docs/docs.module.css';
import Breadcrumbs from './_components/Breadcrumbs/Breadcrumbs';
import TableOfContents from './_components/TableOfContents/TableOfContents';
import { useEffect, useState } from 'react';
import { usePathname } from 'next/navigation';

export default function Layout({ children }: { children: React.ReactNode }) {
	const pathName = usePathname();
	const [mobileSidebarActive, setMobileSidebarActive] = useState(false);

	const toggleMobileSidebar = () => {
		setMobileSidebarActive(!mobileSidebarActive);
	};

	useEffect(() => {
		setMobileSidebarActive(false);
	}, [pathName]);

	return (
		<div className={styles.docsContainer}>
			<Sidebar
				mobileSidebarActive={mobileSidebarActive}
				toggleSidebar={toggleMobileSidebar}
			/>
			<div className={styles.docsMain}>
				<Breadcrumbs toggleSidebar={toggleMobileSidebar} />
				{children}
				<TableOfContents />
			</div>
		</div>
	);
}
