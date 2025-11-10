# ArgoCD End-to-End Deployment Runbook

## Steps

### 1. Create Autopilot Cluster
Create autopilot cluster with the following configuration:
- Private IP
- Regular release channel
- No IP connection
- DNS connection
- No backups
- Auto updates paused

### 2. SSH into Cloud Shell
```bash
gcloud cloud-shell ssh
```

### 3. Authenticate to GCP Project
```bash
gcloud auth login
```
Paste the provided validation token when prompted.

### 4. Install ArgoCD
Follow instructions at: https://argo-cd.readthedocs.io/en/stable/getting_started/

**Important:** Run the following command first to use a GCP managed endpoint for public connectivity, otherwise the install will time out:
```bash
gcloud container clusters get-credentials <clustername> --region <regionname>
```

### 5. Access ArgoCD UI
Do NOT enable port forwarding despite documentation instructions as that will prevent connectivity.

Access ArgoCD at: `https://<IP_ADDR>:8080`

ArgoCD uses port 8080.

### 6. Recovery from Broken ArgoCD Installation
If you break ArgoCD by applying a manifest or ConfigMap, reinstall using the previous steps after deleting the configmap:
```bash
kubectl delete configmap argo -n argo
```

### 7. Update User Password
```bash
argocd account update-password --account <name>
```

### 8. Add New User Account
When generating a new token, add the user to the master ConfigMap at the top:
```yaml
apiVersion: v1
data:
  accounts.justin: apiKey,login  # Name of account - need BOTH values
  admin.enabled: "true"           # Admin account - disable after new account is set
```

### 9. Grant Admin Permissions to New Account
Edit the RBAC ConfigMap:
```bash
kubectl -n argo edit configmap argocd-rbac-cm
```

Add the following to the bottom of the ConfigMap:
```yaml
policy.csv: |
  g, justin, role:admin
```
- `g` is group binding which maps the role to the user
- `p` is permissions for the role if needing to edit that

Reload the ConfigMap:
```bash
kubectl -n argo edit configmap argocd-rbac-cm
```

### 10. Reload Default ConfigMap
```bash
kubectl rollout restart deploy argocd-server -n argo
```

### 11. Disable Admin Account
Validate you can login with the new account, then update the ConfigMap to disable the admin account. Validate you can no longer login with admin.

API keys are used to programmatically manage ArgoCD.

### 12. Build Docker Image
Build, validate tag, and push Docker image.

### 13. Create Application
Create app job for each application with the following settings:
- **App name**: Name of app
- **Project name**: Default
- **Sync Policy**: Automatic & Self Heal
- **Repo Name**: GitHub repo name
- **Rev**: HEAD
- **Path**: Name of folder in main (Argo)
- **Cluster URL**: `https://kubernetes.default.svc` (default)
- **Namespace**: argo

Find namespaces on cluster:
```bash
kubectl get namespaces
```

### 14. Fix Permissions Issues
Permissions issues are likely once ArgoCD is first deployed. Follow these steps to correct:

Validate the service account exists from the install:
```bash
kubectl get clusterrole | grep argocd
```
Look for `argocd-application-controller`.

Check the rules:
```bash
kubectl get clusterrole argocd-application-controller -o yaml
```
Should see wildcards.

Make sure set to the correct namespace:
```bash
kubectl get clusterrolebinding argocd-application-controller -o yaml | grep namespace
```

Probably in the wrong namespace, so update by deleting and recreating:
```bash
kubectl delete clusterrolebinding argocd-application-controller

kubectl create clusterrolebinding argocd-application-controller \
  --clusterrole=argocd-application-controller \
  --serviceaccount=argo:argocd-application-controller
```

Reload the app controller:
```bash
kubectl -n argo rollout restart deployment argocd-application-controller
```

Open a new tab and you should have permissions to create a new app pipeline.

### 15. Validate Deployment
Validate the job syncs, the pipeline completed showing healthy and synced, and confirm via logs from the pod the expected job output indicating RC=0:
```bash
kubectl get pods -n <namespace>
```
Should see completed 0/1.
```bash
kubectl logs <pod> -n <namespace>
```
Check the logs to validate.
