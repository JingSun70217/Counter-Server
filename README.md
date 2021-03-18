

# Introduction
This Counter server is a REST server developed using Go programming language. The server includes the following features:

-	Counter service.
-	Increment a counter.
-	Decrement a counter.
-	Return the value of the counter when asked for.<br>

A Docker image `sunjing70217/counter` (uploaded in DockerHub) is created with the binary of the server. <br>
The server could be deployed on a local Kubernetes.<br>
To work in a production environment, tests were designed for the server.

`/src/main.go` is the main file of this project. <br>
`/src/main_test.go` is the test file of the main functions. <br>
`/Dockerfile` is the Dockerfile for creating Docker images automatically.<br>
`/counter_kube.yml` is the YAML file for creating Kubernetes Deployment and Service.







# Preparation
This server has been tested in virtual machine of Ubuntu20.04 in host system Windows 10 with following setups. <br>
-	Go 1.16.2
-	Docker 19.03.8
-	Kubernetes 1.20
-	Minikube 1.18.1



# Run application in a Kubernetes cluster
Please navigate to `/Counter-Server-main/` then start the following steps:

### Deploy to Kubernetes cluster
Deploy the app to Kubernetes cluster, well, to minikube.<br>
``` $ minikube start``` <br>

![minikube start](https://github.com/JingSun70217/Counter-Server/blob/main/img/1_minikube_start.png)

### Create a Deployment and a Service

The online Docker image `sunjing70217/counter` will be pulled.<br>

``` $ kubectl create -f counter_kube.yml``` 


### Check the status of Deployment and Service with the minikube dashboard (Optional)

``` $ minikube dashboard``` <br>

![dashboard](https://github.com/JingSun70217/Counter-Server/blob/main/img/2_dashboard.png)

### Access the service
``` $ minikube service counter-server ```<br>

The home page will pop up: <br>

![Home Page](https://github.com/JingSun70217/Counter-Server/blob/main/img/3_homePage.png)

Navigate to different functions.<br>

**1. Return the value of the counter when asked for.** The counter number will be returned when navigating to `/getCounter`.<br>
![getCounter](https://github.com/JingSun70217/Counter-Server/blob/main/img/4_getCounter.png)


**2. Increment a counter** when navigating to `/increaseCounter`.<br>

![increaseCounter](https://github.com/JingSun70217/Counter-Server/blob/main/img/5_inCounter.png)

The number of counters will be increased by 1 when reloading the current page.<br>



**3. Decrement a counter when** navigating to `/decreaseCounter`.<br>

![decreaseCounter](https://github.com/JingSun70217/Counter-Server/blob/main/img/6_deCounter.png)

The number of counters will be decreased by 1 when reloading the current page.

# Testing
The `main_test.go` file in `/src/` tests the connection status to port and checks the printed strings for each function.  <br>

Running Test: ``` $ go test```  under `/src/`. <br>

If all tests are passed, the “PASS” result will be visualized.



