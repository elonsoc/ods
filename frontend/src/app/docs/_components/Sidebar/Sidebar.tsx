'use client';

import Link from 'next/link';
import React, { useState } from 'react';
import styles from './Sidebar.module.css';
import { usePathname } from 'next/navigation';
import { configuration } from '@/config/Constants';

const URL = configuration.url.API_URL;

const Sidebar = ({
	mobileSidebarActive,
	toggleSidebar,
}: {
	mobileSidebarActive: boolean;
	toggleSidebar: () => void;
}) => {
	return (
		<aside
			className={`${styles.sidebar} ${
				mobileSidebarActive ? styles.sidebarActive : ''
			}`}
		>
			<nav className={styles.sidebarContainer} aria-label='Related Topics'>
				<header>
					<p className={styles.sidebarHeader}>
						<strong>
							<Link href='docs'>Docs</Link>
						</strong>
					</p>
					<button
						className={styles.closeButton}
						type='button'
						onClick={toggleSidebar}
					>
						<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24'>
							<title>close</title>
							<path d='M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z' />
						</svg>
					</button>
				</header>
				<ul className={styles.sidebarLinks}>
					<li className={styles.sectionHeader}>
						<p>
							<strong className={styles.sublistHeader}>
								<Link href='docs/getting-started'>Getting Started</Link>
							</strong>
						</p>
						<ul className={styles.sublist}>
							<NavLink
								title={'Overview'}
								link={'docs/getting-started/overview'}
							/>
							<NavLink
								title={'Registering an App'}
								link={'docs/getting-started/registering-an-app'}
							/>
							<NavLink
								title={'Making API Calls'}
								link={'docs/getting-started/making-api-calls'}
							/>
						</ul>
					</li>
					<li className={styles.sectionHeader}>
						<p>
							<strong className={styles.sublistHeader}>
								<Link href='docs/usage-guides'>Usage Guides</Link>
							</strong>
						</p>
						<ul className={styles.sublist}>
							<NavLink
								title={'Authentication'}
								link={'docs/usage-guides/authentication'}
							/>
							{/* <NavLink
								title={'Rate Limits'}
								link={'docs/usage-guides/rate-limits'}
							/> */}
							<NavLink
								title={'Error Handling'}
								link={'docs/usage-guides/error-handling'}
							/>
						</ul>
					</li>
					<li className={styles.sectionHeader}>
						<p>
							<strong className={styles.sublistHeader}>
								<Link href='docs/reference'>Reference</Link>
							</strong>
						</p>
						<ul className={styles.sublist}>
							<NavLink
								title={'Data Format'}
								link={'docs/reference/data-format'}
							/>
							<NavLinkDropdown
								title={'Endpoints'}
								link={'docs/reference/endpoints'}
								sublinks={[
									{
										title: 'Buildings v1',
										link: 'docs/reference/endpoints/buildings_v1',
									},
								]}
							/>
						</ul>
					</li>
					<li className={styles.sectionHeader}>
						<p>
							<strong className={styles.sublistHeader}>
								<Link href='docs/resources'>Resources</Link>
							</strong>
						</p>
						<ul className={styles.sublist}>
							<NavLink title={'FAQ'} link={'docs/resources/faq'} />
							<NavLink
								title={'Contact Us'}
								link={'docs/resources/contact-us'}
							/>
							<NavLink
								title={'Terms of Use'}
								link={'docs/resources/terms-of-use'}
							/>
						</ul>
					</li>
				</ul>
			</nav>
		</aside>
	);
};

interface NavLinkProps {
	title: string;
	link: string;
}

const NavLink = ({ title, link }: NavLinkProps) => {
	const pathName = usePathname();
	const path = pathName.replace(URL, '');
	return (
		<li
			className={`${styles.docLink} ${path == '/' + link ? styles.active : ''}`}
		>
			<Link href={link}>{title}</Link>
		</li>
	);
};

interface NavLinkDropdownProps {
	title: string;
	link: string;
	sublinks: NavLinkProps[];
}

const NavLinkDropdown = ({ title, link, sublinks }: NavLinkDropdownProps) => {
	const [dropdownActive, setDropdownActive] = useState(false);
	const pathName = usePathname();
	const path = pathName.replace(URL, '');
	return (
		<li>
			<span
				className={`${styles.dropdownLinkWrapper} ${
					path == '/' + link ? styles.active : ''
				}`}
			>
				<Link href={link}>{title}</Link>
				<button
					className={`${styles.dropdownButton} ${
						dropdownActive ? styles.activeDropdown : ''
					}`}
					type='button'
					onClick={() => setDropdownActive(!dropdownActive)}
				>
					<svg
						className={styles.dropdownSVG}
						xmlns='http://www.w3.org/2000/svg'
						viewBox='0 0 24 24'
					>
						<title>Dropdown</title>
						<path d='M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z' />
					</svg>
				</button>
			</span>
			{dropdownActive && (
				<ul className={styles.sublistLinks}>
					{sublinks.map((sublink, index) => (
						<li
							key={index}
							className={path == '/' + sublink.link ? styles.activeSub : ''}
						>
							<Link href={sublink.link}>{sublink.title}</Link>
						</li>
					))}
				</ul>
			)}
		</li>
	);
};

export default Sidebar;
