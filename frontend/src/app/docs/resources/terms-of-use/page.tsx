import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import { NavigationArrowLeft } from '../../_components/NavigationArrows/NavigationArrows';

const TermsOfUse = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Terms of Use</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					Welcome to the Open Data Service (ODS) API! Before accessing or using
					the ODS API, please carefully read and understand the following terms
					and conditions. By accessing or using the API, you agree to be bound
					by these terms of use. If you do not agree with any part of these
					terms, please refrain from using the API.
				</p>
			</section>
			<section aria-labelledby='API_Usage'>
				<h2 id='API_Usage'>1. API Usage</h2>
				<p>
					1.1. Authorization: You are granted a non-exclusive, non-transferable
					right to access and use the ODS API solely for the purpose of
					integrating it into your applications or services.
				</p>
				<p>
					1.2. API Key: You are responsible for obtaining a valid API key
					through the registration process. Your API key is personal and
					confidential. Do not share your API key with any unauthorized parties.
				</p>
				<p>
					1.3. Acceptable Use: You agree to use the ODS API in a manner
					consistent with applicable laws, regulations, and the intended purpose
					of the API. You shall not engage in any activities that may disrupt,
					damage, or interfere with the API or its associated systems.
				</p>
				<p>
					1.4. Prohibited Use: You shall not use the ODS API to create or
					distribute any applications or services that are illegal, harmful, or
					offensive. You shall not use the API to create or distribute any
					applications or services that violate the rights of any third party.
				</p>
			</section>
			<section aria-labelledby='Data_Usage'>
				<h2 id='Data_Usage'>2. Data Usage</h2>
				<p>
					2.1. Data Ownership: The data provided through the ODS API remains the
					property of Elon University. You may only use the data for the
					purposes specified by the API documentation and within the scope of
					your authorized access.
				</p>
				<p>
					2.2. Data Protection: You shall handle any data obtained through the
					API in accordance with applicable data protection laws and
					regulations. Do not disclose, sell, or misuse the data obtained
					through the API.
				</p>
				<p>
					2.3. Data Attribution: You shall include the following attribution
					statement in any application or service that uses data obtained
					through the API: “This application/service uses data provided by Elon
					University through the Open Data Service API.”
				</p>
			</section>
			<section aria-labelledby='Intellectual_Property'>
				<h2 id='Intellectual_Property'>3. Intellectual Property</h2>
				<p>
					3.1. Ownership: Elon University retains all rights, title, and
					interest in the ODS API, including any associated documentation, code,
					and intellectual property.
				</p>
				<p>
					3.2. Attribution: When using the ODS API, you must provide proper
					attribution to Elon University as the source of the data. The
					attribution guidelines are provided in the API documentation.
				</p>
				<p>
					3.3. Modification: You shall not modify, adapt, or alter the ODS API
					in any way.
				</p>
				<p>
					3.4. Reverse Engineering: You shall not reverse engineer, decompile,
					or disassemble the ODS API.
				</p>
			</section>
			<section aria-labelledby='Limitation_of_Liability'>
				<h2 id='Limitation_of_Liability'>4. Limitation of Liability</h2>
				<p>
					4.1. Warranty: The ODS API is provided on an "as is" and "as
					available" basis. Elon University does not warrant that the API will
					be error-free, uninterrupted, or secure. Your use of the API is at
					your own risk.
				</p>
				<p>
					4.2. Indemnification: You agree to indemnify and hold Elon University,
					its employees, and affiliates harmless from any claims, damages, or
					losses arising out of your use of the API or violation of these terms
					of use.
				</p>
				<p>
					4.3. Limitation of Liability: Elon University shall not be liable for
					any direct, indirect, incidental, special, consequential, or exemplary
					damages arising out of your use of the API or violation of these terms
					of use.
				</p>
			</section>
			<section aria-labelledby='Modifications_and_Termination'>
				<h2 id='Modifications_and_Termination'>
					5. Modifications and Termination
				</h2>
				<p>
					5.1. Modifications: Elon University reserves the right to modify or
					update these terms of use at any time. Any changes will be effective
					immediately upon posting on the API documentation website. It is your
					responsibility to review the terms periodically.
				</p>
				<p>
					5.2. Termination: Elon University may, at its discretion, suspend or
					terminate your access to the API if you violate these terms of use or
					engage in any unauthorized or abusive activities.
				</p>
			</section>
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>6. Conclusion</h2>
				<p>
					If you have any questions or concerns regarding these terms of use,
					please refer to the "Contact Us" section of the API documentation for
					support channels.
				</p>
				<p>
					By accessing and using the ODS API, you acknowledge that you have
					read, understood, and agreed to these terms of use.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft
					link='/docs/resources/contact-us'
					name='Contact Us'
				/>
			</nav>
		</article>
	);
};

export default TermsOfUse;
