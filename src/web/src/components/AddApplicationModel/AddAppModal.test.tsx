import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom';
import userEvent from '@testing-library/user-event';
import AddAppModal from './AddAppModal';

describe('AddAppModal', () => {
	const onAdd = jest.fn();
	const onClose = jest.fn();

	const renderComponent = () => {
		return render(<AddAppModal onAdd={onAdd} onClose={onClose} />);
	};

	it('should render the modal correctly', () => {
		renderComponent();
		expect(screen.getByText('Application Registration')).toBeInTheDocument();
		expect(screen.getByLabelText('Name *')).toBeInTheDocument();
		expect(screen.getByLabelText('Description *')).toBeInTheDocument();
		expect(screen.getByLabelText('Owners *')).toBeInTheDocument();
		expect(screen.getByRole('button', { name: 'Add' })).toBeInTheDocument();
	});

	it('should call onAdd when the form is submitted with valid inputs', async () => {
		renderComponent();
		const nameInput = screen.getByLabelText('Name *') as HTMLInputElement;
		const descriptionInput = screen.getByLabelText(
			'Description *'
		) as HTMLInputElement;
		const ownersInput = screen.getByLabelText('Owners *') as HTMLInputElement;
		const teamNameInput = screen.getByLabelText(
			'Team Name'
		) as HTMLInputElement;
		const addButton = screen.getByRole('button', { name: 'Add' });

		const name = 'Test App';
		const description = 'A test app for Jest';
		const owners = 'Test User';
		const teamName = 'Test Team';

		await userEvent.type(nameInput, name);
		await userEvent.type(descriptionInput, description);
		await userEvent.type(ownersInput, owners);
		await userEvent.type(teamNameInput, teamName);
		fireEvent.click(addButton);

		expect(onAdd).toHaveBeenCalledWith({
			name,
			description,
			owners,
			teamName,
		});
	});

	it('should call onClose when the Close button is clicked', () => {
		renderComponent();
		const closeButton = screen.getByRole('button', { name: 'close' });

		fireEvent.click(closeButton);

		expect(onClose).toHaveBeenCalledWith(false);
	});
});
