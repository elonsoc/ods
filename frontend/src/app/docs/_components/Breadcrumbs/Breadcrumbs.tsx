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
		}
	}, [router]);

	if (!breadcrumbs || breadcrumbs.length <= 1) {
		return null;
	}

	return (
		<nav aria-label='breadcrumbs' className={styles.breadcrumbsWrapper}>
			<ol className={styles.breadcrumbsList}>
				{breadcrumbs.map((breadcrumb: PathInferface, i: any) => {
					return (
						<li key={breadcrumb.href}>
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
