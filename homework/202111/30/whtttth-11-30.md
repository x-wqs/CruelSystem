# LinkedIn Latency Insights

## Introduction

A stream processing system developed to gain latency insights across different flows

## Event selection

First identify the flow we want to measure and select start and endpoints on the flow. We will create StartTrackingEvent and EndTrackingEvent at these points and process al of them. All messages will be emit to messagequeue like kafka and to be processed.

## Stream Processing

Use Samza (stream processing freamework) to consume events from kafka and perform compuation and statistics.

## Join Events

We need to join a start event with its corresponding end event, so we have Partitioner and Joiner which parittions the event as a single push on the same process and calculates the latency of matched events.

## Partitioner

Ensure start and end event to be within a single host. The evetns are originally consumed randly. Both events keeps common data like a key such that they can be bartiioned based on and be consumed by same Joiner later.

## Joiner

Calculating the latency between start and end events. It maintains a in memory cache of all start events and end events, then it checks the matching counter parts and calculate the latency.