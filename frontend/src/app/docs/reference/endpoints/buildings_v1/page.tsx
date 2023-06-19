import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import CodeCopyable from '@/components/CodeCopyable/CodeCopyable';
import CodeBlockContainer from '@/components/CodeBlockContainer/CodeBlockContainer';
import { BUILDING_RESPONSE_SAMPLE } from '@/components/CodeBlockContainer/code';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '@/app/docs/_components/NavigationArrows/NavigationArrows';

const BuildingsV1 = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Buildings v1</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					The version one "buildings" endpoint in the Open Data Service (ODS)
					API provides access to information about buildings at Elon University.
					This endpoint allows you to retrieve details such as building names,
					locations, capacities, and other relevant data.
				</p>
			</section>
			<section aria-labelledby='Retrieving_Buildings'>
				<h2 id='Retrieving_Buildings'>Retrieving Buildings</h2>
				<p>
					To retrieve a list of all buildings, you can make a GET request to the
					following endpoint:
				</p>
				<CodeCopyable code='https://api.ods.elon.edu/v1/buildings' />
				<p>
					This request will return a response containing an array of building
					objects. Each building object represents a specific building at Elon
					University and includes various attributes, such as:
				</p>
				<ul>
					<li>
						<span className='inline-code'>id</span>: Unique identifier for the
						building.
					</li>
					<li>
						<span className='inline-code'>name</span>: Name of the building.
					</li>
					<li>
						<span className='inline-code'>floors</span>: An array of floors
						objects that the building contains.
					</li>
					<li>
						<span className='inline-code'>location</span>: Location or campus
						where the building is situated.
					</li>
					<li>
						<span className='inline-code'>address</span>: Address of the
						building.
					</li>
					<li>
						<span className='inline-code'>type</span>: Type of the building.
					</li>
				</ul>
			</section>
			<section aria-labelledby='Retrieving_Specific_Buildings_by_ID'>
				<h2 id='Retrieving_Specific_Buildings_by_ID'>
					Retrieving Specific Buildings by ID
				</h2>
				<p>
					In addition to retrieving a list of all buildings, you can also
					retrieve information about a specific building by its ID. To do this,
					you can make a GET request to the following endpoint:
				</p>
				<CodeCopyable code='https://api.ods.elon.edu/v1/buildings/{buildingID}' />
				<p>
					Replace <span className='inline-code'>&#123;buildingID&#125;</span> in
					the endpoint URL with the actual ID of the building you want to
					retrieve. For example, to retrieve information about a building with
					an ID of "123456", you would make a GET request to:
				</p>
				<CodeCopyable code='https://api.ods.elon.edu/v1/buildings/123456' />
				<p>
					This request will return a response containing the building object
					with the specified ID. You can then use the retrieved data to display
					detailed information about the building in your application.
				</p>
			</section>
			<section aria-labelledby='Building_Object_Structure'>
				<h2 id='Building_Object_Structure'>Building Object Structure</h2>
				<p>
					Each building object returned by the "buildingsv1" endpoint follows a
					consistent structure. Here's an example representation of a building
					object:
				</p>
				<CodeBlockContainer text={BUILDING_RESPONSE_SAMPLE} codeType='JSON' />
			</section>
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					The "buildings v1" endpoint of the ODS API allows you to access
					comprehensive information about buildings at Elon University. By
					making requests to this endpoint, you can retrieve a list of all
					buildings or retrieve specific buildings by their IDs to obtain
					specific buildings based on your requirements. Utilize the data
					provided by the "buildings v1" endpoint to enhance your applications
					with accurate and up-to-date information about Elon University
					buildings.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/reference/endpoints'
					name='Endpoints'
				/>
				<NavigationArrowRight link='/docs/resources' name='Resources' />
			</nav>
		</article>
	);
};

export default BuildingsV1;
