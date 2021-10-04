- FBNet models components -> typed objects, data attributes -> value fields, association attributes -> relationship fields
- Value fields contain object data
- relationship fields contain typed references to other objects

#### Desired vs Derived
- FBNet models partition into two distinct groups
 - Desired
 - Derived
- Desired models capture the desired network state, maintained by network engineers with a set of specialized tools
 - To make changes, modify the data to describe the updated network design
 - data used to drive the generation of device configs
- Derived models reflect the current operation network state
 - data populated based on real-time collection from network devices
- use case: anomaly detection
 - differences between data could imply expected or unexpected deviation from planned network design

### APIs
#### Read APIs
`get<ObjectType>(fields, query)`
`fields` - a list of value fields relative to the object of the given type
`query`
- criteria the returned list of objects must match
- made of _expressions_
- expression form `<field> <op> <rvalue>`
	- `field` - local or indirect value field to compare to
	- `op` - comparison operator
	- `rvalue` - a list of values to compare against

#### Write APIs
High-level operations that add, update, or delete multiple objects

### Architecture and implementation

#### Storage Layer

- Persistent object store - MySQL
- Each FBNet model <-> dataase table
- Each column <-> field in the model
- Each row <-> object

#### Service Layer
- Read and write APIs are exposed as language-independent Thrift remote procedure calls (RPC)