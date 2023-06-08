import React from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';

export default function GSOverviewPage() {
	return (
		<>
			<Breadcrumbs />
			<div className={styles.docsPageMainContent}>
				<h1>Getting Started</h1>
				<p>
					Lorem ipsum dolor sit, amet consectetur adipisicing elit. Eum qui
					ipsum quasi accusantium laborum enim repellat commodi quod? Eligendi
					delectus pariatur odio exercitationem voluptatem. Asperiores aliquid
					et hic optio possimus.
				</p>
				<p>
					Lorem ipsum, dolor sit amet consectetur adipisicing elit. Nobis, nulla
					ad sapiente nostrum eaque labore veritatis ab dolor, laboriosam
					incidunt, similique quaerat itaque quia ullam possimus nisi quae in
					repellat.
				</p>
			</div>
		</>
	);
}
