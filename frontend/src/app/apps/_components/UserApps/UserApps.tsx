'use client';

import { ApplicationSimple } from '@/app/api/applications/application.d';
import UserApp, { AppInfo } from '../UserApp/UserApp';
import Link from 'next/link';

const UserApps = ({ applications }: { applications: ApplicationSimple[] }) => {
	return (
		<>
			{applications.map((app: ApplicationSimple) => (
				<Link href={`apps/${app.id}`} key={app.id}>
					<UserApp info={app} />
				</Link>
			))}
		</>
	);
};

export default UserApps;
