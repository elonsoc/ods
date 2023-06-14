import React, { Suspense } from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';
import TableOfContents from '../../_components/TableOfContents/TableOfContents';
import CodeCopyable from '@/components/CodeCopyable/CodeCopyable';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

export default function GSOverviewPage() {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Overview</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					This guide will walk you through the process of getting started with
					our API, from logging in to retrieving your API key and making API
					calls. By following these steps, you'll be able to access the rich
					data about Elon University and build powerful applications.
				</p>
			</section>
			<section aria-labelledby='Prerequisites'>
				<h2 id='Prerequisites'>Prerequisites</h2>
				<p>
					Before you begin using the Open Data Service API, make sure you have
					the following prerequisites:
				</p>
				<ul>
					<li>Basic knowledge of HTTP requests and JSON data format.</li>
					<li>An active account on Open Data Service.</li>
				</ul>
			</section>
			<section aria-labelledby='Logging_In'>
				<h2 id='Logging_In'>Logging In</h2>
				<p>
					To access the ODS API, you need to log in to your account. Follow
					these steps:
				</p>
				<ol>
					<li>
						Locate the "Log In" button at the top right corner of the page or
						click the "Getting Started" button on the homepage.
					</li>
					<li>Sign in through Elon Universrity using SSO. </li>
					<li>
						If you have not recently logged in to any Elon University's
						services, enter your credentials &#40;e.g. username and
						password&#41; and log in with two-factor authentication.
					</li>
				</ol>
			</section>
			<section aria-labelledby='Creating_an_Application'>
				<h2 id='Creating_an_Application'>Creating an Application</h2>
				<p>
					Once you are logged in, you need to create an application to obtain an
					API key. Follow these steps:
				</p>
				<ol>
					<li>
						Go to the "Apps" link on your navigation bar on the top right of
						your page.
					</li>
					<li>
						If you are not logged in, you will be prompted to log in first
						before you have the permission to create an application.
					</li>
					<li>
						If you are logged in then you will see one of two options:{' '}
						<ul>
							<li>
								If you do not current have any applications, click the "Register
								an Application" button.
							</li>
							<li>
								If you do have applications, click the "Add App" button on the
								top right of your screen.
							</li>
						</ul>
					</li>
					<li>
						Fill in the required information, including the application name,
						description, owners, and any other relevant details.
					</li>
					<li>Submit the form to create the application.</li>
				</ol>
			</section>
			<section aria-labelledby='Retrieving_Your_API_Key'>
				<h2 id='Retrieving_Your_API_Key'>Retrieving Your API Key</h2>
				<p>
					After creating your application, you can retrieve your API key. Follow
					these steps:
				</p>
				<ol>
					<li>
						Go to the "Apps" link on your navigation bar on the top right of
						your page.
					</li>
					<li>
						Locate the application you created and click on it to view its
						details.
					</li>
					<li>
						{' '}
						In the application details, you will find your API key. Make sure to
						copy and securely store it.
					</li>
				</ol>
			</section>
			<section aria-labelledby='Making_API_Calls'>
				<h2 id='Making_API_Calls'>Making API Calls</h2>
				<p>
					With your API key in hand, you are ready to start making API calls to
					retrieve data from Elon University. Here's an example of a basic API
					call:
				</p>
				<CodeCopyable
					code='curl -X GET "http://api.ods.elon.edu/v1/buildings/" -H
							"Authorization: Bearer [YOUR_API_KEY]"'
				/>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/getting-started'
					name='Getting Started'
				/>
				<NavigationArrowRight
					link='/docs/getting-started/registering-an-app'
					name='Registering an App'
				/>
			</nav>
		</article>
	);
}
