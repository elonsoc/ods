'use client';
import React, { useEffect, useState } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import Link from 'next/link';
import styles from './Breadcrumbs.module.css';
import Sidebar from '../Sidebar/Sidebar';

interface PathInferface {
	breadcrumb: string;
	href: string;
}

const convertBreadcrumb = (string: string) => {
	return string
		.replace(/-/g, ' ')
		.replace(/oe/g, 'ö')
		.replace(/ae/g, 'ä')
		.replace(/ue/g, 'ü')
		.replace(/\w\S*/g, function (txt) {
			return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
		});
};

const Breadcrumbs = () => {
	const router = useRouter();
	const pathName = usePathname();
	const [breadcrumbs, setBreadcrumbs] = useState<PathInferface[]>([
		{ breadcrumb: '', href: '' },
	]);
	const [sidebarActive, setSidebarActive] = useState(false);
	const [loading, setLoading] = useState(true);

	const generateBreadcrumbs = () => {
		const linkPath = pathName.split('/');
		linkPath.shift();

		const pathArray = linkPath.map((path, i) => {
			return {
				breadcrumb: path,
				href: '/' + linkPath.slice(0, i + 1).join('/'),
			};
		});

		setBreadcrumbs(pathArray);
		setLoading(false);
	};

	useEffect(() => {
		if (router) {
			generateBreadcrumbs();
		}
	}, [pathName]);

	if (loading) {
		return (
			<nav aria-label='breadcrumbs' className={styles.breadcrumbsWrapper}>
				<ol className={styles.breadcrumbsList}>
					<li className={`${styles.crumb} ${styles.loadingCrumb}`}>Loading</li>
					<li className={`${styles.crumb} ${styles.loadingCrumb}`}>
						Loading Links...
					</li>
					<li className={`${styles.crumb} ${styles.loadingCrumb}`}>Loading</li>
				</ol>
			</nav>
		);
	}

	if (!breadcrumbs || breadcrumbs.length <= 0) {
		return <></>;
	}

	return (
		<>
			<nav aria-label='breadcrumbs' className={styles.breadcrumbsWrapper}>
				<ol className={styles.breadcrumbsList}>
					<li className={styles.expandSidebarLi}>
						<button
							type='button'
							className={styles.expandSidebarButton}
							onClick={() => setSidebarActive(!sidebarActive)}
						>
							<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
								<title>Expand Sidebar</title>
								<path d='M6,2H18A2,2 0 0,1 20,4V20A2,2 0 0,1 18,22H6A2,2 0 0,1 4,20V4A2,2 0 0,1 6,2M6,8V16H10V8H6Z' />
							</svg>
						</button>
					</li>
					{breadcrumbs.map((breadcrumb: PathInferface, i: any) => {
						return (
							<li className={styles.crumb} key={breadcrumb.href}>
								<Link href={breadcrumb.href}>
									{convertBreadcrumb(breadcrumb.breadcrumb)}
								</Link>
							</li>
						);
					})}
				</ol>
			</nav>
		</>
	);
};

export default Breadcrumbs;
