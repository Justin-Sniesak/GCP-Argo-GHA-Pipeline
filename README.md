# GCP-Argo-GHA-Pipeline
End-to-end CI/CD lab automating Go builds through Docker, GitHub Actions, and GCP Artifact Registry → GKE with ArgoCD. (WIP - Estimated completion early/mid December 2025)

# 🧩 Platform Engineering Pipeline Lab  
**End-to-end CI/CD lab automating Go builds through Docker, GitHub Actions, and GCP Artifact Registry → GKE with ArgoCD. (WIP)**

---

## 🚀 Project Overview  
This lab extends the **Cloud-Native Infrastructure Lab** by focusing on CI/CD design, automation, and deployment using modern DevOps toolchains.  
It demonstrates a full, reproducible pipeline for containerized Go applications from build → test → publish → deploy using **GitHub Actions**, **GCP Artifact Registry**, and **ArgoCD**.

| Category | Tools & Services |
| :-- | :-- |
| **Languages** | Go |
| **Containers** | Docker |
| **CI/CD** | GitHub Actions, ArgoCD |
| **Cloud** | GCP (Artifact Registry, GKE, Pub/Sub, Secrets Manager) |
| **IaC** | Terraform |
| **Security** | IAM, Secret Rotation, API Key Lifecycle |

---

## 📦 Current Objectives (3–4 Week Roadmap)

| Milestone | Description |
| :--- | :--- |
| **Enhanced CI/CD Pipeline & Go Course Completion** | Dockerize code, create enhanced GitHub Actions pipelines, and connect builds to GCP Artifact Registry. |
| **GKE Integration** | Deploy 2-node GKE cluster; configure ArgoCD for continuous deployment; validate end-to-end functionality. |
| **Security Enhancements** | Create Pub/Sub topic, store API keys in Secrets Manager, and configure automatic rotation. |
| **Documentation & Diagrams** | Build architectural flow: `Dockerized Code → GHA → GCP Artifact Registry → GKE → ArgoCD`. |
| **Finalization** | Complete Go “Structs, Methods, and Interfaces” module; expand roadmap for advanced automation phase. |

---

## 🧰 Repository Structure
```
├── /Code/ # Go code and Dockerfiles
├── /Pipeline/ # GitHub Actions workflows
├── /Terraform/ # IaC manifests for GKE and supporting resources
├── /Argo/ # ArgoCD configuration and manifests
├── /Docs_Screenshots/ # Diagrams, Ops logs, architecture notes
└── README.md
```
