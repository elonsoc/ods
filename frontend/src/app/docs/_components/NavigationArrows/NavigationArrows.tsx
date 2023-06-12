import Link from 'next/link';
import React from 'react';

import styles from './NavigationArrows.module.css';

interface NavigationArrowProps {
	link: string;
	name: string;
}

export const NavigationArrowLeft = ({ link, name }: NavigationArrowProps) => {
	return (
		<Link href={link} className={`${styles.arrowLink} ${styles.left}`}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
				<title>Previous</title>
				<path d='M10.05 16.94V12.94H18.97L19 10.93H10.05V6.94L5.05 11.94Z' />
			</svg>
			{name}
		</Link>
	);
};

export const NavigationArrowRight = ({ link, name }: NavigationArrowProps) => {
	return (
		<Link href={link} className={`${styles.arrowLink} ${styles.right}`}>
			{name}
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
				<title>Next</title>
				<path d='M14 16.94V12.94H5.08L5.05 10.93H14V6.94L19 11.94Z' />
			</svg>
		</Link>
	);
};
