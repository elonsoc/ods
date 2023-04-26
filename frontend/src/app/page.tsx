'use client';

import styles from '@/styles/pages/home.module.css';
import Link from 'next/link';
import { useState } from 'react';

export default function Home() {
	const [validUser, setValidUser] = useState<boolean>(false);
	return (
		<div className={styles.container}>
			<h2 className={`${styles.subheading} ${styles.elonSubheading}`}>
				Elon University
			</h2>
			<h1 className={styles.introTitle}>Open Data Service</h1>
			<p className={styles.heroDescription}>
				Access data about Elon University's buildings, courses, and more through
				our API. Register your application for an API key and start building
				innovative applications for the Elon community.
			</p>

			<Link href={validUser ? '/apps' : '/denied'} className={styles.button}>
				Get Started
			</Link>
			<Link href='/' className={styles.learnMoreLink}>
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
			<div className={styles.validUserInput}>
				<input
					type='checkbox'
					id='validUser'
					onChange={() => setValidUser(!validUser)}
				></input>
				<label htmlFor='validUser'>Valid User</label>
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
						href='/'
						className={`${styles.button} ${styles.dataFormatButton}`}
					>
						Data Formats
					</Link>
				</div>
				<div className={styles.jsonContainer}>
					<p className={styles.jsonHeader}>
						<strong>JSON</strong>
					</p>
					<pre className={styles.jsonCode}>
						<code>
							{`{
  course_id: 1234,
  name: Computer Science II,
  description: This course continues the study of object-oriented programming with an emphasis on graphical user interfaces, event handling...,
  instructor: Ryan Mattfeld,
  schedule: [
	{
	  day: Monday,
	  start_time: 12:30 PM,
	  end_time: 1:40 PM
	},
	{
	  day: Wednesday,
	  start_time: 12:30 PM,
	  end_time: 1:40 PM
	},
  ],
  rotation : {
	  semesters: [Fall, Spring]
	  periodical: yearly
  },
  location: Mooney Building, Room 202,
  credits: 4,
  prerequisites: [CSC 1300]
}`}
						</code>
					</pre>
				</div>
			</section>
		</div>
	);
}
