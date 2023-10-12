# GO REST API standard layout
Golang standard layout is mostly used for small to mid size projects. Some advantages of using standard layout are:
-   Suitable for small to medium-sized projects.

-   Emphasizes Go's package-centric design, often referred to as "package-oriented" development. It organizes your code based on packages rather than strictly enforcing a layered architecture.    

-   Works well for projects where simplicity and quick development are priorities.
    
-   This approach encourages a more modular and package-centric organization of your code.
    

-   You can still implement separation of concerns by using package-level encapsulation and defining clear interfaces within your packages.

## Structure:

**cmd**: This directory contains your application's main entry points. Each subdirectory here can represent a different executable or command-line tool for your application. For example, you might have a `cmd/api` directory for your API server and a `cmd/cli` directory for a command-line interface.

**pkg**: The `pkg` directory is where you place packages (Go libraries) that are designed for reuse within your project but could potentially be used by other projects in the future. These packages should be relatively self-contained and independent of your application's business logic.

**internal**: The `internal` directory is for packages that are specific to your project and should not be imported by other projects. It's a way to enforce encapsulation and limit access to certain parts of your codebase.

**configs**: Store your configuration files here. These files might include environment-specific configuration settings, database configuration, and other application settings.