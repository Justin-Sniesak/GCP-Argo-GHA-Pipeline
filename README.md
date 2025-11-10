GCP-Argo-GHA-Pipeline

End-to-end CI/CD lab automating Go builds through Docker, GitHub Actions, and GCP Artifact Registry → GKE with ArgoCD. (Completed 11/09/2025)

🧩 Platform Engineering Pipeline Lab

Fully automated GitOps CI/CD pipeline integrating Go, Docker, GitHub Actions, GCP Artifact Registry, and ArgoCD on GKE.

🚀 Project Overview

This project extends the Cloud-Native Infrastructure Lab with a focus on CI/CD, GitOps, and Kubernetes automation.
It delivers a complete, reproducible pipeline from Go application development → containerization → CI/CD → deployment → validation using modern DevOps toolchains.

Category	Tools & Services
Languages	Go
Containers	Docker
CI/CD	GitHub Actions, ArgoCD
Cloud	GCP (Artifact Registry, GKE, Pub/Sub, Secrets Manager)
IaC	Terraform
Security	IAM, Secret Manager, API Key Rotation
📦 Key Outcomes
Milestone	Description
CI/CD Pipeline Implementation	Automated Go app builds using Docker and GitHub Actions, publishing artifacts to GCP Artifact Registry.
GKE + ArgoCD Deployment	Deployed ArgoCD on a GKE Autopilot cluster with private IPs and validated full end-to-end pipeline sync.
Error Handling & Debugging	Resolved Go map and HOF type mismatches, non-interactive code failures, and Argo folder structure issues.
Security & IAM	Integrated GCP service accounts with least-privilege access and validated ArgoCD authentication with GKE.
Documentation & Visualization	Completed full architecture diagram and runbook, logging 50+ screenshots across all stages.
🧠 Technical Highlights

ArgoCD GitOps Deployment: Automated multi-app deployments via per-folder manifests.

GCP Artifact Registry Integration: GitHub Actions pipelines automatically push Docker images for ArgoCD sync.

GKE Autopilot Configuration: Private IP, IAM-based access, and cluster-level service account integration.

Troubleshooting Discipline: Validated both successful and intentionally failed workloads for educational reproducibility.

🧰 Repository Structure
├── /Code/                # Go code and Dockerfiles
├── /Pipeline/            # GitHub Actions workflows
├── /Argo/                # ArgoCD configuration and manifests
├── /Terraform/           # IaC manifests for GKE and supporting resources
├── /Docs_Screenshots/    # Architecture diagrams, logs, and validation screenshots
└── README.md


Build Time: 7 hours (including troubleshooting and validation)
Completion Date: November 9, 2025

📘 Full runbook: Docs_Screenshots/runbook.md
