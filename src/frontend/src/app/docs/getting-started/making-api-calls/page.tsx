import React from 'react';
import Breadcrumbs from '../../_components/Breadcrumbs/Breadcrumbs';
import styles from '@/styles/pages/docs/docs.module.css';
import TableOfContents from '../../_components/TableOfContents/TableOfContents';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';
import CodeCopyable from '@/components/CodeCopyable/CodeCopyable';
import CodeBlockContainer from '@/components/CodeBlockContainer/CodeBlockContainer';
import { QUERY_PARAMETER_SECTION } from '@/components/CodeBlockContainer/code';

const GSMakingAPICalls = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Making API Calls</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					Once you have registered your application and obtained your API key,
					you are ready to start making API calls to retrieve data from the ODS
					API. This guide will walk you through the process of constructing API
					requests, including required parameters and authentication, to
					retrieve the desired data about Elon University.
				</p>
			</section>
			<section aria-labelledby='Overview_of_API_Endpoints'>
				<h2 id='Overview_of_API_Endpoints'>Overview of API Endpoints</h2>
				<p>
					Before making API calls, it's essential to familiarize yourself with
					the available endpoints provided by the API. Each endpoint represents
					a specific data resource, such as buildings, courses, or other
					relevant information about Elon University. Refer to the API
					documentation or the "Reference" section for a comprehensive list of
					endpoints and their functionalities.
				</p>
			</section>
			<section aria-labelledby='Constructing_API_Requests'>
				<h2 id='Contructing_API_Requests'>Constructing API Requests</h2>
				<p>
					To make an API call, you need to construct a request using the
					appropriate HTTP method (GET) and provide the necessary parameters.
					The required parameters vary depending on the endpoint and the
					specific data you want to retrieve. Consult the API documentation or
					the "Reference" section for detailed information on each endpoint's
					required parameters.
				</p>
			</section>
			<section aria-labelledby='Query_String_Structure'>
				<h2 id='Query_String_Structure'>Query String Structure</h2>
				<p>
					The query string is a part of the URL that follows the "?" character
					and contains key-value pairs separated by "&" symbols. It allows you
					to pass parameters to the API and customize the behavior of your
					request. Here's an example of a query string in the context of the ODS
					API:
				</p>
				<CodeCopyable
					code={`https://api.ods.elon.edu/locations/v1/buildings?location=Main%20Campus
`}
				/>
				<p>
					In the above example, the query string includes two parameters:
					location and capacity. The values for these parameters are
					URL-encoded. In this case, the location parameter is set to "Main
					Campus" and the capacity parameter is set to "100". The URL-encoded
					format replaces spaces with "%20".
				</p>
			</section>
			<section aria-labelledby='Including_Query_Parameters'>
				<h2 id='Including_Query_Parameters'>Including Query Parameters</h2>
				<p>
					When making GET requests to the ODS API, you can include query
					parameters to filter or refine the data you receive. These parameters
					are appended to the base URL as key-value pairs in the query string.
					Here's an example of a GET request to retrieve buildings with specific
					parameters:
				</p>
				<CodeCopyable
					code={`GET /locations/v1/buildings?location=Main%20Campus&capacity=100
`}
				/>
				<p>
					In this example, the request is made to the /locations/v1/buildings
					endpoint, and the query parameters location and capacity are included
					to narrow down the results. The API will respond with a list of
					buildings that match the specified criteria.
				</p>
			</section>
			<section aria-labelledby='Encoding_Query_Parameters'>
				<h2 id='Encoding_Query_Parameters'>Encoding Query Parameters</h2>
				<p>
					When including query parameters in your API calls, it's important to
					properly encode the parameter values to ensure correct interpretation
					by the server. Encoding is necessary to handle special characters,
					spaces, and other reserved characters that may cause issues in the
					URL. Most programming languages provide built-in functions or
					libraries to handle URL encoding.
				</p>
				<p>
					For example, in JavaScript, you can use the{' '}
					<code className='code'>encodeURIComponent()</code> function to encode
					query parameter values:
				</p>
				<CodeBlockContainer
					text={QUERY_PARAMETER_SECTION}
					codeType='JavaScript'
				/>
				<p>
					By properly encoding the query parameter values, you can ensure that
					your API requests are correctly interpreted by the server, especially
					when dealing with special characters or spaces.
				</p>
			</section>
			<section aria-labelledby='Authentication'>
				<h2 id='Authentication'>Authentication</h2>
				<p>
					To authenticate your API requests, you need to include your API key in
					the request headers. The API key serves as your credentials to access
					the protected resources. Ensure that you include the API key in the
					"Authorization" header with the appropriate authentication method.
					Failure to authenticate correctly will result in an unauthorized
					response from the API.
				</p>
			</section>
			<section aria-labelledby='Handling_Responses'>
				<h2 id='Handling_Responses'>Handling Responses</h2>
				<p>
					When you make an API call, you will receive a response from the ODS
					API. The response will contain the requested data or provide relevant
					information about the success or failure of the request. It's
					essential to understand the different response codes and their
					meanings, such as 200 for a successful request or 404 for a resource
					not found. Handle responses appropriately in your application logic to
					provide a seamless user experience.
				</p>
			</section>
			{/* <section aria-labelledby='Rate_Limits'>
				<h2 id='Rate_Limits'>Rate Limits</h2>
				<p>
					The ODS API imposes rate limits to ensure fair usage and maintain the
					performance and stability of the system. Rate limits restrict the
					number of API calls you can make within a specific time period. Refer
					to the API documentation or the "Reference" section for information on
					the specific rate limits applicable to your account or application.
					Ensure that you adhere to these limits to avoid any disruptions in
					accessing the API.
				</p>
			</section> */}
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					With the information provided in this guide, you have the knowledge
					and tools to start making API calls to retrieve data from the ODS API.
					Remember to consult the API documentation or the "Reference" section
					for specific details on available endpoints, required parameters,
					authentication methods, response handling, and rate limits.
				</p>
				<p>
					By effectively making API calls, you can unlock a wealth of
					information about Elon University and leverage it in your
					applications. If you encounter any difficulties or have further
					questions, feel free to reach out to our support team for assistance.
					Happy exploring and retrieving data with the ODS API!
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/getting-started/registering-an-app'
					name='Registering an App'
				/>
				<NavigationArrowRight link='/docs/usage-guides' name='Usage Guides' />
			</nav>
		</article>
	);
};

export default GSMakingAPICalls;
