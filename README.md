# qpid-dispatch-go-tests
Qpid Dispatch Go Test suite for k8s.

# Goal

Propose a sample test suite that can be used to create test scenarios that run on
Kubernetes and can be also used to run in other platforms like OpenShift.
 
This initial sample defines a structure for organizing test suites (ideally _pkg_ will be
kept in a separate repository from the _test_ suites), a framework that is capable of
deploying operators (modified version of qdr-operator e2e test framework), use multiple
kubernetes contexts, run tests in parallel, setting up and tearing down namespaces, run
client applications.

# How it works

# Pre-requisites
You must have a running cluster and you must be logged in to your cluster using
an account that has been granted with cluster admin role (because test suite will
create and remove cluster level resources, like namespaces, roles and CRDs).

The `KUBECONFIG` variable must be set and referring to a kubernetes config that has
the credentials and contexts related with the cluster to be used.

**_Note:_** If no context is specified, the framework will use the current-context
set in your KUBECONFIG file. The framework allows using multiple contexts as well,
and this is a feature that can be used to define tests that run across clusters.
 
## Lifecycle

1. Once a new instance of the Framework is created, a new namespace is created and the
qdr-operator (only for now) is deployed to the new namespace (along with all its 
dependant resources). This happens for every test spec, so they can run 
in parallel, independently from each other.

2. After the Framework has been initialized, the setup of your current test suite has
to be  done, creating the resources needed by your suite (Interconnect, secrets, clients).

3. The test specs, within a test suite, will be executed in any order and they
can also be executed in parallel, which helps reducing the overall time taken by your
test suite (optionally).

4. After each test spec completes, a teardown phase is executed, removing all
resources created (including removal of the generated namespace).

The suite has been defined using Ginkgo (BDD Go test framework). Further info can be 
found at: https://onsi.github.io/ginkgo/.

# Structure of sample test suites

This initial sample (draft for discussion) has a "_test_" directory that contains a few 
test suites nested to it. Each test suite is executed sequentially by Ginkgo, but the 
Specs within each suite can be randomly executed in parallel (optionally).

**_Note:_** The goal is that the final model will provide a framework (_pkg_ directory)
in one repository and the test suites in distinct repositories.

At the _./test_ directory, we have a file named _test_base.go_ which provides the _Initialize_
method and the _init_ method. The _Initialize_ method must be called from your test suite once,
as it will enforce Ginkgo to also invoke the _init_ (which gets invoked first).

Nested to the _test_ directory, we have two other directories:
- _smoke_
- _integration_ (dummy just to illustrate the structure)

Inside the _smoke_ directory, we have two distinct test suites, one called _oneinterior_ and the
other called _twointerior_.

The directory tree should have the following files:
1. A file named something like: "*__suite_test.go_"
   
   This file is used as the entry point to trigger the execution of your test specs, for the related test suite.
   Looking at it you will see that it invokes the Initialize() method defined at the _test_base.go_.

2. _setup.go_

   Provides the logic to create the Framework instance(s) that will be used by the related test suite, and the
   _BeforeEach_ and _AfterEach_ executions that are coordinated by Ginkgo to perform setup and teardown of your
   scenario.
   
3. Go files that **Describe** your test specifications. Examples:
   3. one_interior_deployment.go
   3. one_interior_messaging.go

# Sample test suites

## oneinterior - One Interior router in a single namespace

### Scenario
Uses a minimalist deployment of a single Interconnect node, using interior mode deployed to a single namespace.

### Setup
An instance of the Framework is created at the _BeforeEach_ phase (coordinated by Ginkgo), which starts
the creation of a new (temporary) namespace and the deployment of the qdr-operator into it.

Once the Framework is ready, in the _JustBeforeEach_ phase, the Interconnect (_InterconnectSpec_) resource is created
and deployed using the default context (of other provided through CLI).

It waits till the deployment is available or till some error is returned.
 
### Test specs
Now that the scenario for your test suite is up and running, Ginkgo will run all specs available for the given 
test suite. The order is not guaranteed and that is why it is important to perform the setup and teardown appropriately,
as it also helps when running specs (within the same suite) in parallel.

* **Query routers in the network on each pod (one_interior_deployment.go)**

   This spec will run `qdmanage` tool (through `kubectl`) inside all Pods related with your
   deployment and validate that the number of router nodes (entities) returned match
   the expected number for the deployed network.
   
* **Exchange messages through router mesh (one_interior_messaging.go)**

   Before running each spec, Client containers (using cli-proton-python) will be deployed (one sender and one receiver),
   and those clients will exchange 100 small messages through the deployed router.
   
   This simple test demonstrates an approach for running client applications and validating results.
   
### Teardown
At the _setup.go_ file for the _oneinterior_ suite, the AfterEach method is invoked to clean up resources created,
by the Framework instance, after each spec completes. Otherwise, a suite defining multiple specs would consume lots of
resources till it finishes.

## twointerior - Two Interior Routers against distinct namespaces

### Scenario
Defines a network of routers deployed against distinct namespaces, within the same cluster (context).

### Setup
Two instances of the Framework are created against the same cluster, but each instance will have resources
deployed to their own namespaces.

Once Framework instances are ready, an Interconnect instance is deployed at the "Framework 1" and another instance
is deployed to "Framework 2", and this last one defined an inter-router connector that is set to the service created
by the instance deployed at "Framework 1".

### Test specs
* **Query routers in the network on each pod**

   Similarly to the other suite, this test validates the number of router nodes available in the network,
   by parsing results from `qdmanage`.
   
   Although, in this suite, it expects **two** routers to show up in the list.
   
* **Exchanges anycast messages using multiple clients against the distributed mesh**

   This spec creates 2 senders and 2 receivers connected against each router (against the two distinct namespaces),
   and these 8 clients will exchange 4000 AnyCast messages (256 bytes) through the formed network.
   
### Teardown
Once each spec is done, the _AfterEach_ method of each Framework instance is called to free up resources.

# Running the sample suite

Before you can run the sample test suites, there are a few pre-requisites that must be fulfilled.
  
## Pre-requisites
* Setup your go environment (install go, set GOPATH, ...)
* Install `kubectl`
* Have a running Kubernetes cluster you can use (or start your own cluster)
* export KUBECONFIG variable
* Log into your cluster or setup your contexts (if not yet done)
* Install ginkgo (see: https://onsi.github.io/ginkgo/)

    ```shell script
    $ go get github.com/onsi/ginkgo/ginkgo
    $ go get github.com/onsi/gomega/...
   ```

## Executing tests

Once your cluster is up and running, you can run the test suites by executing:

1. Run all test suites

`ginkgo -v -r ./test`

2. Run a specific test suite

`ginkgo -v -r ./test/smoke/oneinterior/`

3. Run tests against a specific context

`ginkgo -v -r ./test -- --context mycontext`

4. Run tests and save JUnit results to a given directory

`ginkgo -v -r ./test -- --report-dir ./results`
