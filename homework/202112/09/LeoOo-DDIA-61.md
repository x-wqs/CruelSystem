# Datalog
A subset of Prolog.

More general than triple-store. A predicate `(subject, object)` is written.

Key idea is 'rules'. 'Rules' can tell the database about new predicates.

Rules can be derived from data or from other rules. 

Rules can also refer to other rules lie a recursive function.


# Evolving of Data model

Hierarchical (tree) -> Relational -> NonRelational: Document and Graph.

Document and graph data do not enforce a schema for the data they store, which can make it easier to adapt applications to changing requirements.

Explicit schema: enforced on write. Implicit schema: handled on read.

