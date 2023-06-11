'use client';
import React, { useEffect, useState } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import Link from 'next/link';
import styles from './Breadcrumbs.module.css';

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
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		if (router) {
			const linkPath = pathName.split('/');
			console.log(linkPath);
			linkPath.shift();

			const pathArray = linkPath.map((path, i) => {
				return {
					breadcrumb: path,
					href: '/' + linkPath.slice(0, i + 1).join('/'),
				};
			});

			setBreadcrumbs(pathArray);
			setLoading(false);
		}
	}, [router]);

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

	if (!breadcrumbs || breadcrumbs.length <= 1) {
		return <></>;
	}

	return (
		<nav aria-label='breadcrumbs' className={styles.breadcrumbsWrapper}>
			<ol className={styles.breadcrumbsList}>
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
	);
};

export default Breadcrumbs;
