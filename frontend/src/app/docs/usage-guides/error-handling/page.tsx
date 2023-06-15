import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

const UGErrorHandling = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Error Handling</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					When interacting with the Open Data Service (ODS) API, it's important
					to be aware of potential error responses that you may encounter. The
					API service follows standard HTTP status codes to indicate the success
					or failure of a request.
				</p>
			</section>
			<section aria-labelledby='Error_Codes'>
				<h2 id='Error_Codes'>Error Codes</h2>
				<p>
					Here are some common HTTP error responses that you may encounter when
					using the ODS API:
				</p>
				<ul>
					<li>
						<strong>400 Bad Request</strong>: This error indicates that the
						request sent to the API is invalid or improperly formatted. It may
						be due to missing required parameters, invalid data, or incorrect
						syntax. Check your request payload and ensure that it adheres to the
						API's specifications.
					</li>
					<li>
						<strong>401 Unauthorized</strong>: This error occurs when the
						request lacks valid authentication credentials or the provided
						credentials are invalid. Ensure that you include the API key in the
						"Authorization" header using the correct format.
					</li>
					<li>
						<strong>403 Forbidden</strong>: This error is returned when the
						requested resource is restricted and access is not allowed. It
						typically indicates that the authenticated user does not have
						sufficient privileges to access the requested data. Check your
						permissions and ensure that you have the necessary authorization to
						access the specific resource.
					</li>
					<li>
						<strong>404 Not Found</strong>: This error suggests that the
						requested resource could not be found on the server. It may occur if
						the endpoint or resource URL is incorrect or if the resource has
						been moved or deleted. Double-check the endpoint and ensure that it
						matches the API documentation.
					</li>
					<li>
						<strong>500 Internal Server Error</strong>: This error indicates
						that an unexpected error occurred on the server-side. It may be
						caused by a temporary issue or a bug in the API service. If you
						receive this error, it's recommended to contact the support team and
						provide them with relevant details to investigate and resolve the
						issue.
					</li>
				</ul>
			</section>
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					When interacting with the Open Data Service API, it's important to be
					aware of potential error responses that you may encounter. The API
					service follows standard HTTP status codes to indicate the success or
					failure of a request. Here are some common HTTP error responses that
					you may encounter when using the ODS API:
				</p>
				<p>
					These are just a few examples of HTTP error responses that you may
					encounter while using the ODS API. It's important to handle these
					errors gracefully in your applications, providing meaningful error
					messages to users and implementing appropriate fallback mechanisms
					where necessary. Refer to the API documentation for specific error
					codes and their corresponding meanings.
				</p>
				<p>
					By understanding and effectively handling these error responses, you
					can build robust and reliable integrations with the ODS API, ensuring
					a smooth experience for your users even in the face of potential
					errors.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/usage-guides/authentication'
					name='Authentication'
				/>
				<NavigationArrowRight link='/docs/reference' name='Reference' />
			</nav>
		</article>
	);
};

export default UGErrorHandling;
