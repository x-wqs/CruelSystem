# Robotron

## Introduction

- Judicious network management
  - properly configured network - a prerequisite to higher-level network functions
  - Network management involves human interactions - highly risky
  - agile network management enables network evolve quickly
- Challenges in management plane
  - distributed configurations - translating high-level intent into distributed low-level device configurations is difficult
  - multiple omains
   - large internet-facing services (e.g. Facebook) are hosted on a networks of networks
   - each sub-network has unique characteristics
   - devices, topology and management tasks vary per sub-network
  - versioning
   - netowrk topology and routing design can change significantly over time in different parts of the network
   - multiple versions of networks need to be managed
  - dependency - configuring network devices involves handling tightly coupled parameters
  - vendor differences - networks consist of devices from different vendors
- Robotron - a system for managing large and dynamic production networks
  - configuration-as-code
  - validation - avoid config errors
  - extensibility

## Netowrk and use cases
- network management - keeping track of the state of network components (e.g. switches, IPs, circuits) during their life cycle
Users <-> POPs (point-ofpresence) <-> Backbone <-> data centers

### Point-of-presence

- POPs contain a multi-tiered network
  - Peering Routers (PRs)
  - Internet Service Providers (ISps)
  - Backbone Routers (BBs)
- POP management tasks
  - building a new POP - most comprehensive
  - provisioning new peering / transit circuits
  - adjusting link capacity
  - changing BGP config