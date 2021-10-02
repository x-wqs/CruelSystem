### Backbone
- backbone network provides transport among POPs and DCs via optional transport links
- PRs and DRs as edge nodes to set up label-switched paths via BBs
- Augment long-haul capacity across the backbone network with circuit additions

## Robotron overview
- Top-down network management - model high-level operator intent, from which low-level vendor-specifc configs are generated, deployed and monitored

Network design -> config generation -> deployment -> monitoring

*FBNet*
- central repository for information, implemented as an object store. each network componened as an object
- object data and associations - attributes
- single source of truth

*network design*
- translating high-level netowrk design into changes to FBNet objects

*Config generations*
- builds vendor-specific device configs based on object states
- highly vendor- and model-dependent

*Deployment*
- Deploy to network devices
- Changes are deployed in small phases

*Monitoring*
- ensure no deviation from its desired state

## FBNet: MOdeling the network

- The vendor-agnostic, network-wide abstraction layer
- models and stores various netowrk device attributes and network-level attributes and topology descriptions
- design goals
  - data models should be simple and comprehensive
  - data models should be easy to extend and maintain over time
- Provides APIs for applications to query data and make changes

### Data model

#### Object value, and relationship
- network: physical and logical components
- attributes to store component data and associations between components