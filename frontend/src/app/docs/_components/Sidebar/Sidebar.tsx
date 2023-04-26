import { Url } from 'next/dist/shared/lib/router/router';
import Link from 'next/link';
import React from 'react';
import styles from './Sidebar.module.css';

const Sidebar = () => {
	return (
		<nav className={styles.sidebarContainer}>
			<p>
				<strong>Docs</strong>
			</p>
			<ol>
				<NavLink title={'Introduction'} link={'docs/introduction'} />
				<li>
					<strong>Getting Started</strong>
				</li>
				<NavLink title={'Overview'} link={'docs/getting_started'} />
				<NavLink
					title={'Registering an App'}
					link={'docs/registering_applications'}
				/>
				<NavLink title={'Making API Calls'} link={'docs/making_api_calls'} />
				<li>
					<strong>Usage Guides</strong>
				</li>
				<NavLink title={'Authentication'} link={'docs/authentication'} />
				<NavLink title={'Rate Limits'} link={'docs/rate_limits'} />
				<NavLink title={'Error Handling'} link={'docs/error_handling'} />
				<li>
					<strong>Reference</strong>
				</li>
				<NavLink title={'Data Formats'} link={'docs/data_formats'} />
				<li>
					<strong>Resources</strong>
				</li>
				<NavLink title={'FAQ'} link={'/faq'} />
				<NavLink title={'Contact Us'} link={'/contact_us'} />
				<NavLink title={'Terms of Use'} link={'/terms_of_use'} />
			</ol>
		</nav>
	);
};

interface NavLinkProps {
	title: String;
	link: Url;
}

const NavLink = ({ title, link }: NavLinkProps) => {
	return (
		<li>
			<Link href={link}>
				<strong>{title}</strong>
			</Link>
		</li>
	);
};

export default Sidebar;
