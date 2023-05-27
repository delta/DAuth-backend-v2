# Dauth-Backend-V2

## Process of setting a route

1. Put your route in router (or create different router for grouping routes)
2. Make controller associated with that router
3. The controller will be using services that it'll take a parameter
   while constructing the controller, One controller should have one service
4. Now create that Service to be used by the controller
5. Service can take One or more repository for to implement that feature, Service
   Can also require config and other information that can be passed while creating
   the service
6. Create Repositories that might be required by the Service.
7. Repositories required DB connection which is already present

## Layers

Database Connection -> (Models + Repositories) -> Services -> Controllers -> Router -> Application
