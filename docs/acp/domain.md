---
date: 2023-07-27

---

SourceHub's Access Control Policy (ACP) module enables SourceHub's users to access control their applications.

The module implements an Access Control Engine inspired in Google's Zanzibar Authorization Service.


Access Control Paradigm
========================

The main use case SourceHub was designed to support is a Discretionary Access Control model.

In discretionary access control object owners have autonomy to create and delete objects at will, 
furtheremore they are free to grant access to any actor they see fit.


Domain Model
============

Defines the vocabulary used across the application.


## Policy

A `Policy` models the high level access control rules for an usecase.
Any SourceHub user can create a `Policy`.

A Policy usually is created to match an application domain, the resources and relations in it closely reflect the relationships in an application, making them closely coupled to individual applications.

Each Policy has associated to it a set of relationships.
Relationships are created by Actors.

Note on Policy id: the id generation algorithm hashes the policy along with the creator account sequence number, this scheme guarantees that uniqueness and determinism.

## Relationship

A Relationship is an ACP module entity which declares a relation between an Object and a Subject.

## Object

An object represents a generic entity to be access controlled.
Objects are defined by an Id, an opaque identifier which is assumed to be unique accross a Policy.
Objects are associated to a Policy Resource.

The ACP module makes no assumption and does not care about the id scheme used within its policies.


## Subject

A Subject is the entity refered by a Relationship.

Subjects can either be an Actor or an Actor Set.


## Actor

An Actor is a generic term referring to any entity that owns Objects within a Policy.
Actors are represented through a DID.

Effectively actor can represent human users, bots or abstract entities such as corporations.


## Actor Set

An Actor Set represents a set of Actors.
This abstraction is a convenient way of group actors in a relationship.

Actor Sets are defined by an object and a relation.

The intuition about this representation is that every actor related to the object through the actorset relation is part of the actor set.

Actor Sets are recursive, meaning that an actor set can point to another actor set.

## Resource

A Resource defines a namespace for grouping objects.
Resources defines a set of relations and permissions.

## Relation

Relation is a core concept of the access control model.
They are named entities and represent a generic coupling between an Object and an Actor as part of a Relationship.

Relations often represent things like: ownership, membership, containment or roles such as reader.

Each Relation also define a set of Value Restrictions and Managed Relations.


## Value Restriction

## Managed Relation

## Permission

A Permission is used to represent the right of performing an operation to an Object.
Permissions are dynamically evaluated according to its Permission Expression.

eg. Bob wants to perform the operation READ on object bob.txt

## Permission Expression


## Check
