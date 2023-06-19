import Link from 'next/link';
import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import Breadcrumbs from '../_components/Breadcrumbs/Breadcrumbs';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../_components/NavigationArrows/NavigationArrows';

const GettingStarted = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Getting Started</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					The "Getting Started" page provides a comprehensive introduction to
					using the Elon ODS API. It covers the following key points:
				</p>
				<ol>
					<li>
						<Link
							className={styles.inlineLink}
							href='/docs/getting-started/overview'
						>
							Overview
						</Link>
						: Users gain an understanding of the Open Data Service API and its
						purpose in providing access to data about Elon University. It
						highlights the benefits and possibilities of using the API in
						applications.
					</li>
					<li>
						<Link
							className={styles.inlineLink}
							href='/docs/getting-started/registering-an-app'
						>
							Registering an App
						</Link>
						: Users are guided through the process of creating and registering
						their application. This involves logging into their Elon ODS
						account, providing essential information such as app name and
						description, and submitting the registration form.
					</li>
					<li>
						<Link
							className={styles.inlineLink}
							href='/docs/getting-started/making-api-calls'
						>
							Making API Calls
						</Link>
						: Users learn how to make their first API call using their
						registered app and the provided API key. The guide covers the
						necessary endpoints, parameters, and authentication methods to
						successfully retrieve data from the Open Data Service API.
					</li>
				</ol>
				<p>
					By following the steps outlined in the "Getting Started" page, users
					can gain a solid foundation for integrating the Open Data Service API
					into their applications. It equips them with the necessary knowledge
					and steps to register an app, obtain an API key, and start making API
					calls to retrieve data about Elon University.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft link='/docs' name='Introduction' />
				<NavigationArrowRight
					link='/docs/getting-started/overview'
					name='Overview'
				/>
			</nav>
		</article>
	);
};

export default GettingStarted;
