import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';
import CodeCopyable from '@/components/CodeCopyable/CodeCopyable';

const Endpoints = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Endpoints</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					Endpoints in the API represent specific resources or functionalities
					that you can interact with. Each endpoint corresponds to a unique URL
					that you can use to perform various operations and retrieve specific
					data.
				</p>
			</section>
			<section aria-labelledby='Structure_of_an_Endpoint'>
				<h2 id='Structure_of_an_Endpoint'>Structure of an Endpoint</h2>
				<p>
					An endpoint typically consists of two main components: the base URL
					and the endpoint path. The base URL is the root URL of the API, and
					the endpoint path is the specific route that identifies the desired
					resource or operation. When combined, they form the complete URL to
					access a particular endpoint.
				</p>
				<p>For example, consider the following example endpoint:</p>
				<CodeCopyable code='https://api.example.com/v1/users' />
				<p>
					In this example, https://api.example.com is the base URL, and{' '}
					<span className='inline-code'>/v1/users</span> is the endpoint path.
					Together, they form the complete URL to access the "users" endpoint.
				</p>
			</section>
			<section aria-labelledby='Endpoint_Functionality'>
				<h2 id='Endpoint_Functionality'>Endpoint Functionality</h2>
				<p>
					Each endpoint in an API serves a specific purpose or provides access
					to a particular resource. For example, an API may have endpoints for
					retrieving user information, creating new records, searching for data,
					or performing calculations.
				</p>
				<p>
					The functionality of an endpoint is usually documented in the API
					documentation, which provides details about the available endpoints,
					their usage, and the expected request/response formats.
				</p>
			</section>
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					Understanding endpoints is crucial for effectively utilizing an API.
					By identifying the available endpoints, their functionalities, the
					associated HTTP methods, and any required authentication, you can
					interact with the API to retrieve data, create resources, update
					information, or perform other operations based on the API's design and
					capabilities.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/reference/data-format'
					name='Data Format'
				/>
				<NavigationArrowRight
					link='/docs/reference/endpoints/buildings_v1'
					name='Buildings v1'
				/>
			</nav>
		</article>
	);
};

export default Endpoints;
