'use client';

import React, { FormEvent, useState } from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import {
	NavigationArrowLeft,
	NavigationArrowRight,
} from '../../_components/NavigationArrows/NavigationArrows';

const ContactUs = () => {
	const [state, setState] = useState({
		subject: '',
		message: '',
	});

	const handleInputChange = (event: any): void => {
		const { name, value } = event.currentTarget;
		setState((prevInfo) => ({
			...prevInfo,
			[name]: value,
		}));
	};

	return (
		<article className={styles.docsPageMainContent}>
			<header className={styles.docsPageMainHeader}>
				<h1>Contact Us</h1>
			</header>
			<section>
				<p>
					We value your feedback, questions, and suggestions, and are committed
					to providing you with the assistance you need.
				</p>
				<p>
					If you encounter any issues, have technical questions, or require
					additional information related to the ODS API, we encourage you to get
					in touch with us. Our dedicated support team is here to help and
					ensure a smooth and productive experience with the API.
				</p>
			</section>
			<form method='POST'>
				<div className={styles.inputWrapper}>
					<label htmlFor='subject'>
						Subject <span className={styles.requiredRed}>*</span>
					</label>
					<input
						type='text'
						id='subject'
						name='subject'
						value={state.subject}
						onChange={handleInputChange}
						required={true}
					></input>
				</div>
				<div className={styles.inputWrapper}>
					<label htmlFor='message'>
						Message<span className={styles.requiredRed}>*</span>
					</label>
					<textarea
						id='message'
						name='message'
						value={state.message}
						onChange={handleInputChange}
						required={true}
					></textarea>
				</div>
				<button className={styles.submitContactFormButton} type='submit'>
					Submit
				</button>
			</form>
			<nav className={styles.arrowWrapper}>
				<NavigationArrowLeft link='/docs/resources/faq' name='FAQ' />
				<NavigationArrowRight
					link='/docs/resources/terms-of-use'
					name='Terms of Use'
				/>
			</nav>
		</article>
	);
};

export default ContactUs;
