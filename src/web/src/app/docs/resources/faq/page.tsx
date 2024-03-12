import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

const FAQ = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>FAQ</h1>
			</header>
			<section className={styles.introductionSection}>
				<ol>
					<li>
						<strong>
							Q: How do I register an application to obtain an API key?
						</strong>
						<ul>
							<li>
								A: To register an application and obtain an API key, follow the
								steps outlined in the "Registering an App" section of the
								documentation. It provides a detailed guide on creating an
								application and generating an API key for authentication.
							</li>
						</ul>
					</li>
					<li>
						<strong>Q: What authentication method does the API use?</strong>
						<ul>
							<li>
								A: The website uses Single Sign-On (SSO) through Elon
								University's authentication system. When making API calls,
								include the API key in the Authorization header as a form of
								authentication.
							</li>
						</ul>
					</li>
					<li>
						<strong>Q: How can I handle errors returned by the API?</strong>
						<ul>
							<li>
								A: When making API calls, it's important to handle potential
								errors appropriately. The "Error Handling" section in the Usage
								Guides provides information on common error responses and
								suggestions for error handling, such as using try-catch blocks
								or implementing appropriate error handling mechanisms in your
								application.
							</li>
						</ul>
					</li>
					<li>
						<strong>Q: What data format does the API use for responses?</strong>
						<ul>
							<li>
								A: The API primarily uses the JSON (JavaScript Object Notation)
								format for responses. It provides a structured and easily
								parseable format for retrieving data from the API. The "Data
								Format" section in the Reference provides further details on the
								JSON structure used in API responses.
							</li>
						</ul>
					</li>
				</ol>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft link='/docs/resources' name='Resources' />
				<NavigationArrowRight
					link='/docs/resources/contact-us'
					name='Contact Us'
				/>
			</nav>
		</article>
	);
};

export default FAQ;
