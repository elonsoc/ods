'use client';

import UserApp, { AppInfo } from '../UserApp/UserApp';
import Link from 'next/link';

const UserApps = ({ applications }: any) => {
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
