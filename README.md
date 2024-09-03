# Just a test website with a simple service

* Copyright 2024, Keith Paterson, All rights reserved.

* This is quick and dirty and missing some error handling.  Maybe a lot of error handling.
* I'm making serious assumptions about how a resource is set up,
  which may fail if/when I try adding more resources.

# How to access it from k8s

* First, know that this was configured for docker-k8s, not actual k8s
* Build the container and then run the deployment
  * `./build/build.sh service -d`
  * `./deploy/deploy.sh`
* Access it via `http://localhost:32754`
  * yes, only http.  https is more involved to setup/configure and for a local run I don't need that.

## Notes about Ngnix (which is also required)
* install this after starting docker-k8s
* follow instructions here:
  https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-manifests/
  * don't bother with AWS or GCS/Azure steps
* Find the port by querying the svc in the nginx namespace
  * `kubectl get ns` will get you the namespace
  * `kubectl describe svc nginx-ingress --namespace=nginx-ingress`
  And you can use other `kubectl` commands to find other things
* __NOTE__: This only works for http, setting up https would require additional auth related configuraiton
  which I'm not going to do for now.
  
