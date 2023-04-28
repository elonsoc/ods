# ODS Backend

This README concerns the backend of ODS. For information about the frontend, see [frontend/README.md](frontend/README.md). For information about the entire project, see [README.md](README.md).

The backend is written in Go and leverages a dependency injection pattern to support the use of external services. This might be a bit jarring at the beginning but it is a very powerful pattern that allows us to easily swap out services for testing and development. In fact, its necessary to use this because we couldn't mock without it... at least not easily.
