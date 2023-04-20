'use client';

import styles from '@/styles/home.module.css';
import Link from 'next/link';
import { useState } from 'react';
import { Raleway } from 'next/font/google';

const raleway = Raleway({ subsets: ['latin'] });

export default function Home() {
	const [validUser, setValidUser] = useState<boolean>(false);
	return (
		<div className={styles.container}>
			<header>
				<h2 className={styles.elonSubheading}>Elon University</h2>
			</header>
			<h1 className={styles.introTitle}>Open Data Service</h1>
			<p className={styles.heroDescription}>
				Access data about Elon University's buildings, courses, and more through
				our API. Register your application for an API key and start building
				innovative applications for the Elon community.
			</p>

			<Link
				href={validUser ? '/apps' : '/denied'}
				className={`${raleway.className} ${styles.loginButton}`}
			>
				Get Started
			</Link>
			<Link href='/' className={styles.learnMoreLink}>
				Learn More{' '}
				<svg
					xmlns='http://www.w3.org/2000/svg'
					viewBox='0 0 24 24'
					className={styles.rightArrowSVG}
				>
					<title>arrow-right-thin</title>
					<path d='M14 16.94V12.94H5.08L5.05 10.93H14V6.94L19 11.94Z' />
				</svg>
			</Link>
			<div className={styles.validUserInput}>
				<input
					type='checkbox'
					id='validUser'
					onChange={() => setValidUser(!validUser)}
				></input>
				<label htmlFor='validUser'>Valid User</label>
			</div>
		</div>
	);
}
