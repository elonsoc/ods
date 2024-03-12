import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import Link from 'next/link';
import {
  NavigationArrowLeft,
  NavigationArrowRight,
} from '../_components/NavigationArrows/NavigationArrows';

const Resources = () => {
  return (
    <article className={styles.docsPageMainContent}>
      <header className={styles.docsPageMainHeader}>
        <h1>Resources</h1>
      </header>
      <section className={styles.introductionSection}>
        <p>
          The "Resources" section provides you with valuable information and
          additional references to enhance your experience with the ODS API.
          This section includes the following pages:
        </p>
      </section>
      <section>
        <ul>
          <li>
            <Link href='docs/resources/faq' className={styles.inlineLink}>
              Frequently Asked Questions (FAQ)
            </Link>
            : The "Frequently Asked Questions" page addresses common queries and
            provides answers to frequently encountered issues or concerns
            related to the ODS API. It covers a range of topics, including
            registration, authentication, rate limits, error handling, and more.
            Refer to the FAQ section for quick answers to common questions and
            to troubleshoot any challenges you may encounter.
          </li>
          <li>
            <Link
              href='docs/resources/contact-us'
              className={styles.inlineLink}
            >
              Contact Us
            </Link>
            : The "Contact Us" page serves as a direct channel of communication
            between you and our dedicated support team. If you have any
            technical questions, feedback, or need assistance with the API, we
            encourage you to reach out to us through our support channels. Our
            support team is committed to providing prompt and helpful responses
            to ensure a smooth and successful integration of the ODS API into
            your applications.
          </li>
          <li>
            <Link
              href='docs/resources/terms-of-use'
              className={styles.inlineLink}
            >
              Terms of Use
            </Link>
            : The "Terms of Use" page outlines the terms and conditions that
            govern your use of the ODS API. It provides important information
            regarding the rights and responsibilities of users, acceptable usage
            policies, data usage guidelines, intellectual property, and more. It
            is essential to review and understand the Terms of Use before
            utilizing the API to ensure compliance and a mutually beneficial
            relationship between you and the ODS service.
          </li>
        </ul>
        <p>
          By providing these additional resources, we aim to assist you in
          maximizing the benefits of the ODS API, addressing common questions or
          concerns, and establishing clear guidelines for usage and
          communication.
        </p>
        <p>
          Feel free to explore the pages within the "Resources" section for
          further information and guidance. We are committed to supporting your
          integration efforts and fostering a collaborative environment for
          utilizing the ODS API effectively and responsibly.
        </p>
      </section>
      <nav className={styles.arrowWrapper}>
        <NavigationArrowLeft
          link='/docs/reference/endpoints'
          name='Endpoints'
        />
        <NavigationArrowRight link='/docs/resources/faq' name='FAQ' />
      </nav>
    </article>
  );
};

export default Resources;
