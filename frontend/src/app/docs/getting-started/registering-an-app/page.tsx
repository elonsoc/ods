import React from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';
import TableOfContents from '../../_components/TableOfContents/TableOfContents';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

const GSRegisteringAnApp = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Registering an App</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					To access the ODS API and retrieve data about Elon University, you
					need to register an application and obtain an API key. This guide will
					walk you through the process of registering an app and obtaining the
					necessary information.
				</p>
			</section>
			<section aria-labelledby='Prerequisites'>
				<h2 id='Prerequisites'>Prerequisites</h2>
				<p>
					Before registering an app with the ODS API, make sure you have the
					following prerequisites:
				</p>
				<ul>
					<li>An active account on Open Data Service.</li>
				</ul>
			</section>
			<section aria-labelledby='Creating_an_Application'>
				<h2 id='Creating_an_Application'>Creating an Application</h2>
				<p>To register an application, follow these steps:</p>
				<ol>
					<li>Log in to your ODS account.</li>
					<li>Go to the "Apps" section of the top navigation bar.</li>
					<li>
						Locate the "Create an Application" button
						<ul>
							<li>
								If you do not current have any applications, click the "Register
								an Application" button on the center of your screen.
							</li>
							<li>
								If you do have applications, click the "Add App" button on the
								top right of your screen.
							</li>
						</ul>
					</li>
					<li>
						After the button is clicked, an "Add Application" modal will be
						displayed.
					</li>
					<li>
						Enter the appropriate required fields for your application:
						<ul>
							<li>Name</li>
							<li>Description</li>
							<li>Owners</li>
						</ul>
					</li>
					<li>Submit the form to create your application.</li>
				</ol>
			</section>
			<section aria-labelledby='Retrieving_Your_API_Key'>
				<h2 id='Retrieving_Your_API_Key'>Retrieving Your API Key</h2>
				<p>
					After successfully creating your application, you will need to
					retrieve your API key. Follow these steps:
				</p>
				<ol>
					<li>Go to the "Apps" section of the top navigation bar.</li>
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
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					By following the steps outlined in this guide, you can successfully
					register an application with the ODS API and retrieve your API key.
					This key is crucial for authenticating your API calls.
				</p>
				<p>
					Next, you can explore the Making API Calls section to learn how to use
					your API key to retrieve data from the ODS API.
				</p>
				<p>
					If you have any questions or encounter any issues during the
					registration process, feel free to reach out to our support team for
					assistance.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/getting-started/overview'
					name='Overview'
				/>
				<NavigationArrowRight
					link='/docs/getting-started/making-api-calls'
					name='Making API Calls'
				/>
			</nav>
		</article>
	);
};

export default GSRegisteringAnApp;
