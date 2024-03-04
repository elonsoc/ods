'use client';

import CodeBlockContainer from '@/components/CodeBlockContainer/CodeBlockContainer';
import { HOME_PAGE_COURSE_SAMPLE } from '@/components/CodeBlockContainer/code';
import styles from '@/styles/pages/home.module.css';
import Link from 'next/link';
import { configuration } from '@/config/Constants';
import { useAuth } from '@/context/auth/auth';

export default function Home() {
	const { isAuthenticated } = useAuth();
	return (
		<div className={styles.container}>
			<div className={styles.heroContainer}>
				<h2 className={`${styles.subheading} ${styles.elonSubheading}`}>
					Elon University
				</h2>
				<h1 className={styles.introTitle}>Open Data Service</h1>
				<p className={styles.heroDescription}>
					Access data about Elon University's buildings, courses, and more
					through our API. Register your application for an API key and start
					building innovative applications for the Elon community.
				</p>
				{isAuthenticated ? (
					<Link href={'/apps'} className={styles.button}>
						Get Started
					</Link>
				) : (
					<Link
						href={configuration.url.BACKEND_API_URL}
						rel='noopener noreferrer'
						className={styles.button}
					>
						Log In with Elon
					</Link>
				)}
				<Link href='/docs' className={styles.learnMoreLink}>
					Learn More{' '}
					<svg
						xmlns='http://www.w3.org/2000/svg'
						viewBox='0 0 24 24'
						className={styles.rightArrowSVG}
					>
						<title>Forward</title>
						<path d='M14 16.94V12.94H5.08L5.05 10.93H14V6.94L19 11.94Z' />
					</svg>
				</Link>
			</div>
			<section className={styles.accessibleDataSection}>
				<div className={styles.dataSectionInfo}>
					<h3 className={`${styles.subheading} ${styles.dataFormatSubheading}`}>
						Data Formats
					</h3>
					<h2 className={styles.dataFormatHeading}>Easily Accessible Data</h2>
					<p className={styles.dataFormatDescription}>
						Open Data Service provides a simple and easy way to access
						comprehensive data about Elon University through our API.
						<br />
						<br />
						With just a few lines of code, you can retrieve information about
						buildings, courses, and more in a format that's easy to integrate
						into your applications.
						<br />
						<br />
						Our API is designed to be developer-friendly, with standardized
						endpoints and data formats that make it simple to get the data you
						need. Whether you're a student, researcher, or developer, our API
						provides a powerful tool for accessing data about Elon University.
						<br />
						<br />
						Click the "Getting Started" button below to learn more about how to
						get started with our API and start accessing the data you need.
					</p>
					<Link
						href='docs/reference/data-format'
						className={`${styles.button} ${styles.dataFormatButton}`}
					>
						Data Format
					</Link>
				</div>
				<CodeBlockContainer text={HOME_PAGE_COURSE_SAMPLE} codeType='JSON' />
			</section>
		</div>
	);
}
