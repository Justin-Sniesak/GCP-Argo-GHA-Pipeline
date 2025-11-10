## 🧩 GCP-Argo-GHA-Pipeline

End-to-end CI/CD lab automating Go builds through Docker, GitHub Actions, and GCP Artifact Registry → GKE with ArgoCD.

Completed 11 / 09 / 2025 · Total Build Time: 7 hours

---

## 🚀 Project Overview

This lab extends the Cloud-Native Infrastructure Lab, focusing on CI/CD design, GitOps automation, and Kubernetes deployment.
It demonstrates a fully reproducible workflow for containerized Go applications —
build → test → publish → deploy → validate — leveraging modern DevOps toolchains.

Category	Tools & Services
- Languages	Go
- Containers	Docker
- CI/CD	GitHub Actions · ArgoCD
- Cloud	GCP (Artifact Registry · GKE · Secrets Manager)
- Security	IAM · Secret Manager 

---

## 🧩 Core Objectives (Completed)

Milestone	Description
- CI/CD Pipeline Implementation	Automated Go builds via Docker and GitHub Actions → published images to GCP Artifact Registry.
- ArgoCD Integration	Deployed ArgoCD on a GKE Autopilot cluster with private IPs and validated sync and rollout functionality.
- Error Handling & Debugging	Corrected Go map/HOF type mismatches, non-interactive failures, and Argo manifest structure issues.
- Security Hardening	Integrated service accounts with least privilege and validated Argo controller ClusterRoleBindings.
- Documentation & Visualization	Completed full architecture diagram, 50 + screenshots, and comprehensive runbook documentation.

---

## ⚙️ Technical Highlights

- GitOps End-to-End: ArgoCD continuously monitors and syncs GitHub repo manifests to GKE.
- Artifact Registry Integration: CI/CD automatically builds, tags, and pushes Docker images to GCP Artifact Registry.
- Kubernetes Automation: Validated automated pod health, rollouts, and replica management using GKE Autopilot.
- Troubleshooting Depth: Included both successful and intentionally failed pods to demonstrate environment behavior.
- Security Posture: IAM roles scoped per service account with enforced API key rotation and SSH ingress blocking.

---

## 🧱 Architecture

Workflow Steps

1. Code Go app locally (Ubuntu laptop)
2. Build and tag Docker image
3. Push code to GitHub main branch
4. GitHub Actions pipeline triggers image build
5. Push image to GCP Artifact Registry
6. ArgoCD syncs GKE cluster to updated manifests
7. Validate pod health, logs, and resource state via kubectl
8. Troubleshoot, iterate, and redeploy

---
```
## 🧰 Repository Structure
├── Code/                # Go code and Dockerfiles
├── Pipeline/            # GitHub Actions workflows
├── Argo/                # ArgoCD configuration and manifests
├── Terraform/           # IaC manifests for GKE and supporting resources
├── Docs_Screenshots/    # Architecture diagrams, logs, validation screenshots
└── README.md
```

---

## 📘 Documentation

Full runbook and troubleshooting steps:

➡️ [Docs_Screenshots/runbook.md](https://github.com/Justin-Sniesak/GCP-Argo-GHA-Pipeline/blob/main/Docs_Screenshots/runbook.md )
