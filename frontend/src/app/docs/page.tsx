import { Metadata } from 'next';
import React from 'react';
import styles from '@/styles/pages/docs/docs.module.css';
import { NavigationArrowRight } from './_components/NavigationArrows/NavigationArrows';

export const metadata: Metadata = {
  title: 'Docs',
  description: 'Documentation for the Open Data Service API',
};

export default function DocsPage() {
  return (
    <article className={styles.docsPageMainContent}>
      <header className={styles.docsPageMainHeader}>
        <h1>Introduction</h1>
      </header>
      <section className={styles.introductionSection}>
        <p>
          Welcome to the Open Data Service API documentation! This comprehensive
          guide provides detailed information on how to integrate and utilize
          the ODS API to access data about Elon University. Whether you're a
          student, researcher, or developer, this documentation will help you
          harness the power of our API to retrieve information about buildings,
          courses, and more.
        </p>
      </section>
      <section aria-labelledby='Documentation_Structure'>
        <h2 id='Documentation_Structure'>Documentation Structure</h2>
        <p>The documentation is organized into the following sections:</p>
        <ul>
          <li>
            {' '}
            <strong>Getting Started</strong>
            <ul>
              <li>
                The "Getting Started" section offers a step-by-step guide to
                help you quickly set up and begin using the Open Data Service
                API. It covers prerequisites, registration of your application,
                retrieving your API key, and making your first API call. If
                you're new to the API, this is the perfect place to begin.
              </li>
            </ul>
          </li>

          <li>
            {' '}
            <strong>Usage Guides</strong>
            <ul>
              <li>
                The "Usage Guides" section provides detailed guides on how to
                authenticate your API requests, understand and manage rate
                limits, and effectively handle errors returned by the API. These
                guides will ensure that you can securely access the data you
                need, efficiently manage your API usage, and gracefully handle
                any errors that may occur.
              </li>
            </ul>
          </li>
          <li>
            {' '}
            <strong>Reference</strong>
            <ul>
              <li>
                The "Reference" section serves as a comprehensive reference
                guide for the Open Data Service API. It contains detailed
                information about available endpoints, request parameters, and
                response formats. This section is invaluable when you need
                precise details about the API's functionalities and how to
                interact with it programmatically.
              </li>
            </ul>
          </li>
          <li>
            {' '}
            <strong>Resources</strong>
            <ul>
              <li>
                The "Resources" section provides additional helpful resources to
                assist you in your API integration. It includes FAQs, contact
                information for our support team, and terms of use. This section
                ensures that you have access to the necessary resources to
                overcome any challenges and make the most out of the API.
              </li>
            </ul>
          </li>
        </ul>
      </section>
      <section aria-labelledby='Getting_Started'>
        <h2 id='Getting_Started'>Getting Started</h2>
        <p>
          If you're new to the ODS API, we recommend starting with the "Getting
          Started" section. It will guide you through the initial setup process
          and provide you with the essential knowledge to begin integrating the
          API into your applications.
        </p>
      </section>
      <nav className={styles.arrowWrapper}>
        <NavigationArrowRight
          link='/docs/getting-started'
          name='Getting Started'
        />
      </nav>
    </article>
  );
}
