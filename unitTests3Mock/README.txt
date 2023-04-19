=== Introduction ===

In this exercise you are going to rework existing code, so testing single components separately becomes easy.

The application architecture uses 3 layers: service, entity and storage.

Service is the interface to the user, where input data enters the program and the results are being returned.

Entity contains the structs of simple objects to represent input and output data. It may contain business logic confined
within one entity. But higher logic concerning different entities typically lives in the service layer.

Storage is the layer that offers persistence of entities.

The way this project is structured right now, unit testing the service layer separately without running real storage code
is impossible. Let's do something about that!

=== Tasks ===

Divide each unit test into 3 sections: //arrange, //act and //assert.

1. Create a unit test for the filter function. So far, no refactoring of the existing code is necessary, as the entity
layer is at the bottom and has no dependencies. This is just to get you warmed up.

2. Create a unit test for service.GetOld(). Unit tests must not involve layers and components other than the one to be
tested. You will have to refactor the code in order to decouple the service layer from the storage layer. This will
allow to test the service function separately without involving real storage. Use inversion of control and a mockup
storage struct.

3. Working on existing code might make you stumble into broken code. Run the unit tests in person_test.go.
Can you fix what is broken?