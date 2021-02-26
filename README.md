
# Camunda App

Complete example of Application that use [Camunda BPM](https://camunda.com/products/camunda-bpm/) to accompany part 6 of my blog series [My Journey with Camunda](https://medium.com/@arifsetiawan/my-journey-with-camunda-toc-3030da004511). It is recommended to go through blog series to understand what we want to achieve here.

This repo has 3 services and uses Golang as backend and Vue as frontend.

- Auth proxy for JWT based auth for Camunda REST API 
- External task handler
- Company Portal (called Nexus) with example Leave Request workflow. This app is built with [CoreUI free Vue.js Admin Template](https://coreui.io/vue/). Two main functions for the Portal are:
    - Create request that will trigger workflow execution
    - Working with user tasks

Note that Send Email task is doing nothing, but it should be easy to add email sender function using Golang smtp for example.

# Development

## 1. Deploy and setup Camunda

See [Deploy and setup Camunda](Camunda.md)

## 2. Run proxy and external tasks

```
make clean && make
```

Open new tab to run proxy
```
make -C deploy run-proxy
```

Open new tab to run external task handler
```
make -C deploy run-external-task
```

## 3. Run nexus

```
cd nexus
```
