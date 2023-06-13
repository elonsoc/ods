import React from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';
import TableOfContents from '../../_components/TableOfContents/TableOfContents';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

const GSMakingAPICalls = () => {
	return (
		<div className={styles.docsPageMainContent}>
			<h1>Making API Calls</h1>
			<p>
				Once you have registered your application and obtained your API key, you
				are ready to start making API calls to retrieve data from the ODS API.
				This guide will walk you through the process of constructing API
				requests, including required parameters and authentication, to retrieve
				the desired data about Elon University.
			</p>
			<h2 id='Overview_of_API_Endpoints'>Overview of API Endpoints</h2>
			<p>
				Before making API calls, it's essential to familiarize yourself with the
				available endpoints provided by the ODS API. Each endpoint represents a
				specific data resource, such as buildings, courses, or other relevant
				information about Elon University. Refer to the API documentation or the
				"Reference" section for a comprehensive list of endpoints and their
				functionalities.
			</p>
			<h2 id='Contructing_API_Requests'>Constructing API Requests</h2>
			<p>
				To make an API call, you need to construct a request using the
				appropriate HTTP method (GET) and provide the necessary parameters. The
				required parameters vary depending on the endpoint and the specific data
				you want to retrieve. Consult the API documentation or the "Reference"
				section for detailed information on each endpoint's required parameters.
			</p>
			<h2 id='Authentication'>Authentication</h2>
			<p>
				To authenticate your API requests, you need to include your API key in
				the request headers. The API key serves as your credentials to access
				the protected resources. Ensure that you include the API key in the
				"Authorization" header with the appropriate authentication method, such
				as "Bearer" or "API Key." Failure to authenticate correctly will result
				in an unauthorized response from the API.
			</p>
			<h2 id='Handling_Responses'>Handling Responses</h2>
			<p>
				When you make an API call, you will receive a response from the ODS API.
				The response will contain the requested data or provide relevant
				information about the success or failure of the request. It's essential
				to understand the different response codes and their meanings, such as
				200 for a successful request or 404 for a resource not found. Handle
				responses appropriately in your application logic to provide a seamless
				user experience.
			</p>
			<h2 id='Rate_Limits'>Rate Limits</h2>
			<p>
				The ODS API imposes rate limits to ensure fair usage and maintain the
				performance and stability of the system. Rate limits restrict the number
				of API calls you can make within a specific time period. Refer to the
				API documentation or the "Reference" section for information on the
				specific rate limits applicable to your account or application. Ensure
				that you adhere to these limits to avoid any disruptions in accessing
				the API.
			</p>
			<h2 id='Conclusion'>Conclusion</h2>
			<p>
				With the information provided in this guide, you have the knowledge and
				tools to start making API calls to retrieve data from the ODS API.
				Remember to consult the API documentation or the "Reference" section for
				specific details on available endpoints, required parameters,
				authentication methods, response handling, and rate limits.
			</p>
			<p>
				By effectively making API calls, you can unlock a wealth of information
				about Elon University and leverage it in your applications. If you
				encounter any difficulties or have further questions, feel free to reach
				out to our support team for assistance. Happy exploring and retrieving
				data with the ODS API!
			</p>

			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/getting-started/registering-an-app'
					name='Registering an App'
				/>
				<NavigationArrowRight link='/docs/usage-guides' name='Usage Guides' />
			</nav>
		</div>
	);
};

export default GSMakingAPICalls;
