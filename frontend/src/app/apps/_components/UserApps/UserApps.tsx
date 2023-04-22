'use client';

import { useState } from 'react';
import UserApp, { AppInfo } from '../UserApp/UserApp';
import AddAppModal from '../UserAppModal/AddAppModal';
import Link from 'next/link';

export interface InformationDetails {
	name: string;
	description: string;
	owners: string;
	teamName: string;
}

const UserApps = ({ applications, handleSubmit }: any) => {
	const [apps, setApps] = useState<InformationDetails[]>([]);
	const hasApps = applications.length;

	async function handleAddApp(appInfo: InformationDetails) {
		// 	// const { name, description, owners, teamName } = appInfo;
		handleSubmit(appInfo);
	}

	return (
		<>
			{applications.map((app: AppInfo) => (
				<Link href={`apps/${app.id}`} key={app.id}>
					<UserApp info={app} />
				</Link>
			))}
		</>
	);
};

export default UserApps;
