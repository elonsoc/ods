import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';
import CodeBlockContainer from '@/components/CodeBlockContainer/CodeBlockContainer';
import {
	REQUEST_PAYLOAD_SAMPLE,
	RESPONSE_PAYLOAD_SAMPLE,
} from '@/components/CodeBlockContainer/code';

const RDataFormat = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Data Format</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					The Open Data Service API primarily uses the JSON (JavaScript Object
					Notation) data format for both request payloads and response data.
					JSON is a lightweight and widely supported data interchange format
					that provides a simple and human-readable structure. It is well-suited
					for representing structured data, making it an ideal choice for
					exchanging information with the ODS API.
				</p>
			</section>
			<section>
				<h2 id='Request_Payloads'>Request Payloads</h2>
				<p>
					When sending data to the ODS API, you will typically include a request
					payload in the body of your API requests. The request payload should
					be formatted as a JSON object. It allows you to pass parameters,
					filters, or additional data necessary for the API to process your
					request.
				</p>
				<p>
					For example, when creating a new object such as a building or a
					course, you can construct a JSON object that includes the relevant
					properties and their corresponding values:
				</p>
				<CodeBlockContainer text={REQUEST_PAYLOAD_SAMPLE} codeType='JSON' />
				<p>
					Ensure that you adhere to the specific request payload requirements
					and validation rules outlined in the API documentation for each
					endpoint.
				</p>
			</section>
			<section aria-labelledby='Response_Data'>
				<h2 id='Response_Data'>Response Data</h2>
				<p>
					The data returned by the ODS API in response to your requests will
					also be in JSON format. The response data will typically be an object
					or an array of objects, depending on the nature of the request. Each
					object represents a resource or an entity and contains various
					attributes and their corresponding values.
				</p>
				<p>
					Here's an example response for a request to retrieve information about
					a specific building:
				</p>
				<CodeBlockContainer text={RESPONSE_PAYLOAD_SAMPLE} codeType='JSON' />
				<p>
					You can access the properties of the response data by parsing the JSON
					and extracting the relevant values. The structure of the response data
					may vary depending on the endpoint and the specific resource being
					requested.
				</p>
			</section>
			<section aria-labelledby='Libraries and Tools'>
				<h2 id='Libraries_and_Tools'>Libraries and Tools</h2>
				<p>
					Working with JSON in your applications is made easier with the
					availability of libraries and tools across different programming
					languages. These libraries provide functions and utilities to parse
					JSON, generate JSON, and manipulate JSON data structures.
				</p>
				<p>
					Depending on your programming language of choice, you can leverage
					libraries such as <span className='inline-code'>json</span> in Python,{' '}
					<span className='inline-code'>Json.NET</span> in C#,{' '}
					<span className='inline-code'>json-simple</span> in Java, or the
					built-in JSON support in JavaScript to handle JSON data effortlessly.
				</p>
			</section>
			<section
				aria-labelledby='Handling_JSON_in_Your_Application
'
			>
				<h2 id='Handling_JSON_in_Your_Application'>
					Handling JSON in Your Application
				</h2>
				<p>
					To interact with the ODS API effectively, your application should be
					equipped to handle JSON data. This involves parsing JSON responses,
					validating JSON structures, and serializing/deserializing JSON
					objects. Refer to the documentation of your chosen programming
					language or framework to learn more about working with JSON in your
					application.
				</p>
				<p>
					The JSON format's simplicity, readability, and wide adoption make it a
					suitable choice for exchanging data with the ODS API. By understanding
					the structure of JSON request payloads and response data, as well as
					leveraging relevant libraries and tools, you can seamlessly integrate
					the API into your applications and efficiently work with the data
					provided by the ODS API.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft link='/docs/reference' name='Reference' />
				<NavigationArrowRight
					link='/docs/reference/endpoints'
					name='Endpoints'
				/>
			</nav>
		</article>
	);
};

export default RDataFormat;
