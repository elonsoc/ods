import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import Breadcrumbs from '../_components/Breadcrumbs/Breadcrumbs';
import TableOfContents from '../_components/TableOfContents/TableOfContents';
import Link from 'next/link';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../_components/NavigationArrows/NavigationArrows';

const UsageGuides = () => {
	return (
		<>
			<Breadcrumbs />
			<div className={styles.docsPageMainContent}>
				<header>
					<h1>Usage Guides</h1>
				</header>
				<p>
					The "Usage Guides" section provides detailed information and best
					practices for utilizing the Open Data Service API effectively. Whether
					you're a student, researcher, or developer, these guides will help you
					understand and leverage the various aspects of the ODS API to access
					data about Elon University.
				</p>
				<ol>
					<li>
						<Link
							className={styles.inlineLink}
							href='docs/usage-guides/authentication'
						>
							Authentication
						</Link>
						: Authentication is a crucial aspect of working with the ODS API.
						This guide will walk you through the various authentication methods
						supported by the API. It will provide step-by-step instructions on
						how to authenticate your API requests, ensuring secure and
						authorized access to the data. Understanding and implementing
						authentication correctly is essential for protecting sensitive
						information and maintaining data integrity.
					</li>
					<li>
						<Link
							className={styles.inlineLink}
							href='docs/usage-guides/rate-limits'
						>
							Rate Limits
						</Link>
						: To ensure fair usage and maintain optimal performance, the ODS API
						enforces rate limits. This guide will explain how rate limits work,
						including the specific limits applicable to your account or
						application. You will learn strategies to effectively manage and
						optimize your API usage within these limits, preventing any
						disruptions in accessing the data. Understanding rate limits will
						help you design efficient and scalable applications that make the
						most out of the ODS API.
					</li>
					<li>
						<Link
							className={styles.inlineLink}
							href='docs/usage-guides/error-handling'
						>
							Error Handling
						</Link>
						: While working with APIs, encountering errors is inevitable. This
						guide will provide an overview of the different types of errors you
						might encounter when interacting with the ODS API and how to handle
						them gracefully. It will cover common error codes, error response
						formats, and best practices for error handling in your applications.
						Understanding and properly handling errors will enhance the
						reliability and robustness of your integrations with the ODS API.
					</li>
				</ol>
				<p>
					The "Usage Guides" section equips you with the knowledge and best
					practices for utilizing the ODS API to its full potential. By
					following the guides on authentication, rate limits, and error
					handling, you can ensure secure access, efficient usage, and effective
					error management while retrieving data about Elon University.
				</p>
				<nav className={styles.arrowWrapper}>
					<NavigationArrowLeft
						link='/docs/getting-started/making-api-calls'
						name='Making API Calls'
					/>
					<NavigationArrowRight
						link='/docs/usage-guides/authentication'
						name='Authentication'
					/>
				</nav>
			</div>
		</>
	);
};

export default UsageGuides;
