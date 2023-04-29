function handleSubmit(event: any) {
	event.preventDefault();
}

const ApplicationSettingsForm = () => {
	return <form onSubmit={handleSubmit}></form>;
};
