
# GO Clean Architecture

I have borrowed some of the GO standard layout which will be described below in details. 
For more info: https://github.com/golang-standards/project-layout


## Structure:

  

**cmd**: This directory contains your application's main entry points. Each subdirectory here can represent a different executable or command-line tool for your application. For example, you might have a `cmd/api` directory for your API server and a `cmd/cli` directory for a command-line interface.

  

**pkg**: The `pkg` directory is where you place packages (Go libraries) that are designed for reuse within your project but could potentially be used by other projects in the future. These packages should be relatively self-contained and independent of your application's business logic. So `pkg`is used for code that can be reused across different projects. If you have generic utilities or libraries that might be useful elsewhere, you can put them here.

  

**internal**: The `internal` directory is for packages that are specific to your project and should not be imported by other projects. It's a way to enforce encapsulation and limit access to certain parts of your codebase. So this is a directory meant for code that is specific to your application and not meant to be reused.

  

**configs**: Store your configuration files here. These files might include environment-specific configuration settings, database configuration, and other application settings.

  

**web**: If your application includes web assets (HTML templates, JavaScript files, CSS stylesheets, etc.), you can organize them in this directory.

## Clean Architecture:
The Clean Architecture, proposed by Robert C. Martin, focuses on separating concerns into layers with strict dependencies among them. It consists of several concentric circles, with the innermost circle representing the core business logic or domain, surrounded by layers like Use Cases, Interface Adapters, and Frameworks & Drivers. In Clean Architecture, the emphasis is on maintaining independence of the inner layers from the outer layers, allowing for easier testing, maintenance, and adaptability. 

The key point in Clean Architecture is that the inner layers (use cases/entities) should not depend on the outer layers (such as frameworks, databases, or UI). Instead, the dependencies should flow from the outer layers towards the inner layers, allowing for flexibility, testability, and easier maintenance.

**Terminology:** 
Some terminology may differ from team to team
1. Domain / Model / Entity
2. Use Cases / Services / Interactors
3. Presenters / Controllers / Delivery
4. Repository / Gateway



**Clean Flow:** 
The flow typically involves:


1.  **Controllers Depending on Frameworks**: It's common for controllers in Clean Architecture to depend on frameworks or libraries for handling HTTP requests or other external inputs. For example using GIN for handling HTTP requests is a practical implementation detail that fits within the boundaries of the Clean Architecture. The key is to isolate these dependencies within the outer layers and prevent them from leaking into the core business logic.
    
2.  **Controllers Depending on Services/Usecases**: This represents the flow where the controllers, after handling incoming requests using the framework, pass control to the services or use cases to execute the application-specific logic. This demonstrates the separation of concerns, with controllers acting as the entry point and delegating business logic to the appropriate layers.
    
3.  **Services/Usecases Depending on Models/Domain**: This aligns with the Clean Architecture principle where the core business logic resides in the inner layers. Services or use cases should indeed depend on the domain models or entities, as they encapsulate the essential business rules and data structures.
    
4.  **Services/Usecases Depending on Repository Interfaces**: Utilizing repository interfaces as abstractions is a core tenet of Clean Architecture. By depending on repository interfaces rather than concrete implementations, your use cases or services remain decoupled from specific data storage mechanisms or implementations. This allows for flexibility in swapping or modifying data access methods without affecting the business logic.


Although the layered architecture pattern does not specify the number and types of layers that must exist in the pattern, most layered architectures consist of four standard layers: presentation, business, persistence, and database. 


## Clean vs N-Tier:
Both **Clean Architecture** and **N-Tier** **Architecture** aim to organize software into layers for better maintainability, separation of concerns, and scalability. They share similarities in the sense that they advocate for dividing an application into distinct layers to manage complexity and improve maintainability.

In **Clean Architecture**, there's a strong emphasis on Dependency Inversion Principle (DIP) and Dependency Injection. Clean Architecture advocates for the use of abstractions/interfaces (loosely coupled and testable code) to define interactions with external resources (like databases, file systems, etc.), keeping the inner layers independent of the concrete implementations of these resources.

**N-Tier** Architecture, on the other hand, does not specifically enforce this level of abstraction. While it also divides the application into layers (presentation, business logic, data access), it might not inherently emphasize the use of interfaces to abstract data access.


## Clean Diagrams:
![93830264-afa9c480-fcaa-11ea-9589-7c5308c291f4](https://github.com/adisnuhic/go-clean/assets/17688087/1c031497-567f-4381-b4f6-1efa1c503e05)
![clean-arch](https://github.com/adisnuhic/go-clean/assets/17688087/629abc87-6862-450f-950e-92bb14b7b343)


## N-Tier Diagrams:
![1_mCAPI1Q7y7XfyrtB58DzkQ](https://github.com/adisnuhic/go-clean/assets/17688087/275cdf95-e9a5-4c7d-bed8-e9e3f4ad2bf0)








