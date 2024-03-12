import Link from 'next/link';
import React from 'react';
import styles from '@/styles/pages/application.module.css';

const BackLink = () => {
	return (
		<Link href='/apps' className={styles.backLink}>
			<svg
				className={styles.leftArrow}
				xmlns='http://www.w3.org/2000/svg'
				viewBox='0 0 24 24'
			>
				<title>Back</title>
				<path d='M10.05 16.94V12.94H18.97L19 10.93H10.05V6.94L5.05 11.94Z' />
			</svg>
			Back to Apps
		</Link>
	);
};

export default BackLink;
