## Troubleshooting Callout
Two distinct code debugging events logged thus far.

The following logs document troubleshooting and validation events related to Go map handling and input validation.

## ArgoCD
**Summary:** Argo installation, functionality and validation.

- 2025-11-09 Install Argo on GKE cluster.
![ACD1-1](../Docs_Screenshots/ACD1-1.jpg)

- 2025-11-09 Install Argo CLI on GKE cluster.
![ACD1-2](../Docs_Screenshots/ACD1-2.jpg)

- 2025-11-09 Patch Argo to GCP provided loadbalancer, then obtain the GUI external IP.  
![ACD1-3](../Docs_Screenshots/ACD1-3.jpg)

- 2025-11-09 Validate ArgoCD GUI comes up externally, login, change the admin password after obtaining the default one-time use initial password.
![ACD1-4](../Docs_Screenshots/ACD1-4.jpg)

- 2025-11-09 Disable and validate that the admin account cannot login to the GUI.
![ACD1-5](../Docs_Screenshots/ACD1-5.jpg)

## Code Debugging
**Summary:** Code Debugging and validation - what went wrong and what was fixed.

- ## 2025-11-05 TROUBLESHOOTING: HOF variables were mismatched which was breaking the map, as they were not uniformly defined in a manner each function could call during code execution when provided with user input.
![CD1-1](../Docs_Screenshots/CD1-1.jpg)

- 2025-11-05 Code validation to ensure code is prompting user for inputting then properly returning the expected integers.
![CD1-2](../Docs_Screenshots/CD1-2.jpg)

- ## 2025-11-06 TROUBLESHOOTING: Error handling in code failing due to map of main user input function included too many data types: float64 and string. Corrected by adding HOF later in the code to an if-else statement which returns the desired string which handles erroneous input from the enduser.
![CD1-3](../Docs_Screenshots/CD1-3.jpg)

- 2025-11-07 Added additional functionality - Thank you to the user after the app returns the total sales tax for their purchase, added as a defer function. 
![CD1-5](../Docs_Screenshots/CD1-5.jpg)

- 2025-11-07 Smoke test of code once container is built using the -it flag. (Numbering out of order due to GitHub image caching bug).  
![CD1-4](../Docs_Screenshots/CD1-4.jpg)

## Result: Smoketest code is functional -> Dockerize code -> Add additional functionality to code rebuild image -> Smoke test code runs once container is built using the -it flag.

## CI/CD Pipeline Validation
**Summary:** Push to GCP Artifact registry; pipeline operation, validation, troubleshooting.

- 2025-11-09 Build both Docker images for each of the apps for Argo pipeline.
![PL1-1](../Docs_Screenshots/PL1-1.jpg)

- 2025-11-09 Tag and push each of the Docker images to GCP Artifact Registry.
![PL1-2](../Docs_Screenshots/PL1-2.jpg)

- 2025-11-09 Validate each of the Docker images is now in GCP Artifact Registry.
![PL1-3](../Docs_Screenshots/PL1-3.jpg)

- ## 2025-11-09 TROUBLESHOOT: Both of the app manifests are in the same root level folder in GitHub which is breaking one of the pipelines. Deleted pipeline, created child folders in GitHub for each pipeline, then recreated individual apps in Argo.
![PL1-4](../Docs_Screenshots/PL1-4.jpg)

- 2025-11-09 Validate each pipeline is synced, one the app is healthy the other it is degraded which is expected and intended behavior. 
![PL1-5](../Docs_Screenshots/PL1-5.jpg)

- 2025-11-09 Validate UTC Go job pod completed and returned an RC=0. Also validated expected output is returned in the pod log.
![PL1-6](../Docs_Screenshots/PL1-6.jpg)

- 2025-11-09 Succesful pipeline validation of UTC Go jobpipeline.
![PL1-7](../Docs_Screenshots/PL1-7.jpg)

- 2025-11-09 This pipeline is intentionally in a failing state, as the code requires user input, in a container would require exec -it, which will not function as intended in a Kubernetes pipeline, as the control plane will prevent the pod from creating. It is provided as a side by side comparion to the UTC pod which ran from end-to-end with an RC=0 (Good run).
![PL1-8](../Docs_Screenshots/PL1-8.jpg)

## Result: Docker image validated pushed from local machine to GCP project, confirming authentication (via API) and succesful push

## GKE
**Summary:** Cluster creation, administation and pipeline smoketesting/validation.

- 2025-11-07 Create GKE autopilot cluster and validate up and running.
![GKE1-1](../Docs_Screenshots/GKE1-1.jpg)
