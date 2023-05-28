import React from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';

const UGRateLimits = () => {
	return (
		<>
			<Breadcrumbs />
			<div className={styles.docsPageMainContent}>
				<h1>Rate Limits</h1>
			</div>
		</>
	);
};

export default UGRateLimits;
