import { Metadata } from 'next';

export const metadata: Metadata = {
	title: 'Applications',
};

export default function Layout({ children }: { children: React.ReactNode }) {
	return <>{children}</>;
}
