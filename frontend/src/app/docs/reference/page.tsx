import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import Link from 'next/link';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../_components/NavigationArrows/NavigationArrows';

const Reference = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Reference</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					The "Reference" section of the ODS API documentation provides
					essential information to understand and utilize the API effectively.
					It consists of two main areas: "Data Format" and "Endpoints."
				</p>
			</section>
			<section>
				<ul>
					<li>
						<Link
							href='docs/reference/data-format'
							className={styles.inlineLink}
						>
							Data Format
						</Link>
						: The "Data Format" subsection explains the JSON format used in API
						responses. JSON (JavaScript Object Notation) is a widely adopted
						format for structuring and representing data. Understanding the JSON
						structure is crucial for parsing and extracting information from the
						API responses.
					</li>
					<li>
						<Link href='docs/reference/endpoints' className={styles.inlineLink}>
							Endpoints
						</Link>
						: The "Endpoints" subsection covers the various endpoints available
						in the ODS API. Each endpoint represents a specific resource or
						functionality provided by the API. The documentation provides an
						overview of the available endpoints, their descriptions, and the
						supported HTTP methods.
					</li>
				</ul>
				<p>
					The "Reference" section offers a comprehensive resource to understand
					the API's data format and explore the available endpoints, enabling
					developers to effectively interact with the API.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/usage-guides/error-handling'
					name='Error Handling'
				/>
				<NavigationArrowRight
					link='/docs/reference/data-format'
					name='Data Format'
				/>
			</nav>
		</article>
	);
};

export default Reference;
