import React from 'react';
import Breadcrumbs from './_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';

export default function DocsPage() {
	return (
		<>
			<Breadcrumbs />

			<div className={styles.docsPageMainContent}>
				<h1>Introduction</h1>
				<p>
					Lorem ipsum dolor sit, amet consectetur adipisicing elit. Quidem sint
					soluta in doloribus iure provident ipsam sequi nam. Quas aliquam
					impedit voluptatum. Modi reprehenderit error iste quae similique harum
					sed!
				</p>
				<p>
					Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maiores,
					porro itaque sint accusantium omnis sed et odit, quia eius esse
					facilis sapiente tenetur dignissimos natus praesentium, nihil mollitia
					quo.
				</p>
			</div>
		</>
	);
}
