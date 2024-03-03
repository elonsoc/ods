import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';
import CodeCopyable from '@/components/CodeCopyable/CodeCopyable';

const UGAuthentication = () => {
	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Authentication</h1>
			</header>
			<section className={styles.introductionSection}>
				<p>
					Authentication is a crucial step in interacting with our API. This
					guide will walk you through the authentication process, ensuring
					secure access to the data about Elon University. Open Data Service
					utilizes a single sign-on (SSO) approach to be able to register for
					applications combined with an API key for authentication.
				</p>
			</section>
			<section aria-labelledby='Single_Sign-On'>
				<h2 id='Single_Sign-On'>Single Sign-On (SSO)</h2>
				<p>
					ODS leverages the Elon University Single Sign-On (SSO) system for user
					authentication. To access the API or an API key, you need to have a
					valid Elon University account. By utilizing your Elon University
					credentials, you can authenticate and gain authorized access to the
					API resources. This SSO approach ensures a seamless and secure
					authentication process.
				</p>
			</section>
			<section aria-labelledby='Including_API_Key_in_Requests'>
				<h2 id='Including_API_Key_in_Requests'>
					Including API Key in Requests
				</h2>
				<p>
					In addition to the SSO authentication, you also need to include your
					API key in the requests to authenticate your API calls. After logging
					in via SSO and obtaining your API key, you should include it in the
					"Authorization" header of your API requests. The API key serves as a
					unique identifier for your application and grants you access to the
					protected resources.
				</p>
				<p className='code'>Authorization: YOUR_API_KEY</p>
				<p>
					Ensure that you replace `YOUR_API_KEY` with the actual API key
					obtained from your registered application.
				</p>
			</section>
			<section aria-labelledby='Authentication_Example'>
				<h2 id='Authentication_Example'>Authentication Example</h2>
				<p>
					Here's an example of an API request using cURL, including the API key
					in the header:
				</p>
				<CodeCopyable
					code={`curl -X GET https://api.ods.elon.edu/locations/v1/buildings/ -H "Authorization: [YOUR_API_KEY]" | json_pp`}
				/>
				<p>Make sure to replace `YOUR_API_KEY` with your actual API key.</p>
			</section>
			<section aria-labelledby='Conclusion'>
				<h2 id='Conclusion'>Conclusion</h2>
				<p>
					Follow the instructions provided in this guide to ensure proper
					authentication when interacting with the API. By authenticating your
					API calls including the API key in the 'Authorization' header, you can
					securely access the data about Elon University and unlock its full
					potential. If you have any further questions or require additional
					support, our team is here to assist you throughout the authentication
					process. We are committed to providing a seamless and secure
					experience for all users of the API.
				</p>
			</section>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft link='/docs/usage-guides' name='Usage Guides' />
				{/* <NavigationArrowRight
					link='/docs/usage-guides/rate-limits'
					name='Rate Limits'
				/> */}
				<NavigationArrowRight
					link='/docs/usage-guides/error-handling'
					name='Error Handling'
				/>
			</nav>
		</article>
	);
};

export default UGAuthentication;
