import { POST, GET } from './route';
import { NextResponse, NextRequest } from 'next/server';

const applications = [
	{
		id: 1,
		name: 'Application Name',
		description:
			'Aggregates course information from various educational institutions or online course providers using APIs and displays it in a user-friendly format, allowing users to browse and compare different courses, and also enabling them to enroll in courses directly through the platform.',
		owners: 'John Doe, Jane Doe',
		teamName: 'Team name',
	},
	{
		id: 2,
		name: 'CourseFinder',
		description:
			'A web application that utilizes APIs to provide comprehensive course information to users, including course descriptions, schedules, and pricing.',
		owners: 'Mark Johnson',
		teamName: 'CourseFinder Team',
	},
	{
		id: 3,
		name: 'CourseHub',
		description:
			'A mobile app that aggregates course information from multiple sources, including online course providers and universities, allowing users to easily search and enroll in courses from their mobile devices.',
		owners: 'Sarah Lee',
		teamName: 'CourseHub Development Team',
	},
	{
		id: 4,
		name: 'Learnly',
		description:
			'An e-learning platform that utilizes APIs to provide users with a personalized learning experience, offering curated course recommendations and progress tracking features.',
		owners: 'Adam Smith, Mary Johnson',
		teamName: 'Learnly Development Team',
	},
];

describe('API endpoints', () => {
	test('POST should add new application to applications array', async () => {
		const request = new NextRequest('/api', {
			method: 'POST',
			body: JSON.stringify({
				id: 5,
				name: 'New Application',
				description: 'A new application',
				owners: 'Jane Doe',
				teamName: 'Team name',
			}),
			headers: {
				'Content-Type': 'application/json',
			},
		});

		const response = await POST(request);

		expect(response).toEqual(
			NextResponse.json({
				id: 5,
				name: 'New Application',
				description: 'A new application',
				owners: 'Jane Doe',
				teamName: 'Team name',
			})
		);

		expect(applications).toContainEqual({
			id: 5,
			name: 'New Application',
			description: 'A new application',
			owners: 'Jane Doe',
			teamName: 'Team name',
		});
	});

	test('GET should return all applications if no id provided', async () => {
		const request = new NextRequest('/api', {
			method: 'GET',
		});

		const response = await GET(request);

		expect(response).toEqual(NextResponse.json(applications));
	});

	test('GET should return empty array if no applications found', async () => {
		const request = new NextRequest('/api?id=5', {
			method: 'GET',
		});

		const response = await GET(request);

		expect(response).toEqual(NextResponse.json([]));
	});

	test('GET should return application with provided id', async () => {
		const request = new Request('/api?id=2', {
			method: 'GET',
		});

		const response = await GET(request);

		expect(response).toEqual(
			NextResponse.json([
				{
					id: 2,
					name: 'CourseFinder',
					description:
						'A web application that utilizes APIs to provide comprehensive course information to users, including course descriptions, schedules, and pricing.',
					owners: 'Mark Johnson',
					teamName: 'CourseFinder Team',
				},
			])
		);
	});
});
