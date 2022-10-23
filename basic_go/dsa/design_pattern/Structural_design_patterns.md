# Structural design patterns

> last updated: 2022/10/23

## intro

Structural design patterns `describe the relationships between the entities`. They are used to form large structures using classes and objects.

## Adapter

The adapter pattern provides a wrapper with an interface required by the API client to link incompatible types and `act as a translator between the two types`.

The adapter uses the interface of a class to be a class with another compatible interface.

The `dependency inversion principle can be adhered to by using the adapter pattern`, when a class defines its own interface to the next level module interface implemented by an adapter class

`Delegation` is the other principle used by the adapter pattern.

The adapter pattern comprises the target, adaptee, adapter, and client:

- `Target` is the interface that the client calls and invokes methods on the adapter and adaptee.
- The `client` wants the incompatible interface implemented by the adapter.
- The adapter translates the incompatible interface of the adaptee into an interface that the client wants.

## Bridge

`Bridge decouples the implementation from the abstraction`. The abstract base class can be subclassed to provide different implementations and `allow implementation details to be modified easily`.

The bridge patterns allow the implementation details to change at runtime.

The bridge pattern demonstrates the principle, `preferring composition over inheritance`.

Composition maintains a **_has-a_** relationship with the implementation, instead of an **_is-a_** relationship.
