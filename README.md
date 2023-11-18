# Design Philosophies

## Project Structure
This section outlines the structure of my project, which follows a package-oriented design to organize code and promote maintainability. The project is divided into several packages, each serving a specific purpose and [layer](#layers) in my application.

## Layers

### App

- **Responsibility**: Manages start-up and shutdown procedures, accepts and validates external input and output.
- **Subpackages**:
  - **Services**: Provides specific services for the application.
    - **API**: Handles REST API endpoints.
  - **Tooling**: Contains various tooling packages.
    - **Logfmt**: Handles log formatting.
    - **Init**: Initialize, migrate and seed database.

### Business

- **Responsibility**: Represents the core business logic of the application.
- **Subpackages**:
  - **Core**: Offer APIs for the business layer. Each subpackage under Core defines the data the project works with.
    - **Store**: Data access layer responsible for CRUD operations for the data.
  - **Database**: Manages database-related operations, including but not limited to managing database driver and database migration tools.
  - **Web**: Includes packages specific to the web application.
    - **Middleware**: Defines middleware for the web layer.

### Foundation

- **Responsibility**: Houses packages that are foundational to ALL businesses. Logging should not occur within the Foundation layer, and packages within Foundation should not import each other.
- **Subpackages**:
  - **Logger**: Helper functions to use `log/slog` package.
  - **Web**: Extension to the [CHI router](https://github.com/go-chi/chi) for extended functionality.

## Notes

- **Core Package**: The `Core` package within the `Business` layer serves as the entry point for the business layer and should be imported by other packages in app layer that need business-related functionality.
- **Stores Package**: Packages within the `Stores` subpackage should not import each other. If a package needs to access multiple entities, use the `Core` package instead.
- **Logging**: Logging should not occur within the `Foundation` layer. 

By adhering to this structure, I aim to create a maintainable and scalable codebase that promotes code clarity and separation of concerns.
