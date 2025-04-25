
# CI/CD Workshop: GitHub + Docker + Minikube

This project demonstrates a complete CI/CD pipeline using GitHub Actions, Docker, and Minikube. It automates building, testing, scanning, and deploying a containerized application.

## 🚀 Project Overview

- **Repository**: [github.com/mrckurz/cicd-workshop](https://github.com/mrckurz/cicd-workshop)
- **Technologies**: GitHub Actions, Docker, Trivy, SonarCloud, Minikube
- **Pipeline Stages**:
  - Build Docker image
  - Push image to Docker Hub
  - Run vulnerability scan with Trivy
  - Perform code quality analysis with SonarCloud
  - Deploy to local Minikube cluster

---

## 🌿 Branches

This repository follows a structured branching model to progressively build the CI/CD pipeline step by step. Each branch represents a different step in the pipeline implementation.

### Main Branch (`main`)

- **Description**: This is the primary branch where the **complete solution** resides. All final code and configurations for the CI/CD pipeline, Docker setup, security scanning, and deployment are merged into this branch.
- **Usage**: The `main` branch contains the finished version of the pipeline that runs on GitHub Actions, including all the stages like build, push, vulnerability scan, code analysis, and deployment to Minikube.

### Step-by-Step Branches

These branches represent intermediate steps of the CI/CD pipeline implementation and allow for incremental learning.

#### Step 1: `step-1-base-setup`

- **Description**: This branch sets up the initial pipeline structure with a simple `echo` command to verify pipeline execution.
- **Focus**: Introducing the `.github/workflows/go.yml` file and understanding basic pipeline operations.
- **Use Case**: When starting with CI/CD basics and setting up the first workflow.

#### Step 2: `step-2-docker-build`

- **Description**: This branch adds a job to build a Docker image for the project.
- **Focus**: Adding Docker build commands and understanding how to package applications into containers.
- **Use Case**: Building the Docker image and pushing it to Docker Hub.

#### Step 3: `step-3-vulnerability-scan`

- **Description**: This branch adds the Trivy vulnerability scan to ensure security checks for the Docker image.
- **Focus**: Integrating Trivy into the pipeline to scan for known vulnerabilities in the image.
- **Use Case**: Enhancing security by scanning images for vulnerabilities before deployment.

#### Step 4: `step-4-sonarcloud`

- **Description**: This branch adds the SonarCloud integration for code quality analysis.
- **Focus**: Configuring SonarCloud to analyze the source code for quality issues and maintainability.
- **Use Case**: Adding code quality checks to the pipeline to ensure the health of the codebase.

#### Step 5: `step-5-deploy-minikube`

- **Description**: This branch adds the deployment step to Minikube, enabling automatic deployment of the built image to a local Kubernetes cluster.
- **Focus**: Deploying Docker containers to a Kubernetes cluster using Minikube.
- **Use Case**: Testing and deploying the application in a Kubernetes environment locally.

Each step allows incremental learning and builds on top of the previous one. These branches give a clear progression for setting up a complete CI/CD pipeline.

---

## 🧰 Prerequisites

### Windows

- **Git**: [Download](https://git-scm.com/download/win)
- **Docker Desktop**: [Download](https://www.docker.com/products/docker-desktop/)
- **Minikube**: [Installation Guide](https://minikube.sigs.k8s.io/docs/start/)
- **kubectl**: [Installation Guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- **GitHub-hosted Runner (automatic)**: [Installation Guide](https://docs.github.com/en/actions/hosting-your-own-runners/about-self-hosted-runners)

### macOS

- **Homebrew**: [Installation Guide](https://brew.sh/)
- **Git**: `brew install git`
- **Docker Desktop**: [Download](https://www.docker.com/products/docker-desktop/)
- **Minikube**: `brew install minikube`
- **kubectl**: `brew install kubectl`
- **GitHub-hosted Runner (automatic)**: [Installation Guide](https://docs.github.com/en/actions/hosting-your-own-runners/about-self-hosted-runners)

### Linux

- **Git**: `sudo apt install git` or `sudo yum install git`
- **Docker**: [Installation Guide](https://docs.docker.com/engine/install/)
- **Minikube**: [Installation Guide](https://minikube.sigs.k8s.io/docs/start/)
- **kubectl**: [Installation Guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- **GitHub-hosted Runner (automatic)**: [Installation Guide](https://docs.github.com/en/actions/hosting-your-own-runners/about-self-hosted-runners)

---

## 🧪 CI/CD Pipeline Explained

The `.github/workflows/go.yml` defines the following stages:

### 1. **Build**

- **Image**: `docker:latest`
- **Script**:
  - Build Docker image tagged with `latest` and commit SHA.

### 2. **Push**

- **Image**: `docker:latest`
- **Script**:
  - Push both `latest` and commit SHA tags to Docker Hub.

### 3. **Vulnerability Scan**

- **Image**: `aquasec/trivy:latest`
- **Script**:
  - Scan the Docker image for vulnerabilities using Trivy.

### 4. **Code Quality Analysis**

- **Image**: `sonarsource/sonar-scanner-cli:latest`
- **Script**:
  - Run SonarCloud analysis using `sonar-scanner`.
  - Requires `SONAR_TOKEN` environment variable.

### 5. **Deploy**

- **Image**: `docker:latest`
- **Script**:
  - Deploy the Docker image to Minikube using `redeploy.sh`.

---

## 🛠 Manual Commands

### Build Docker Image

```bash
docker build -t mrckurz/cicd-workshop-image:latest -t mrckurz/cicd-workshop-image:$(git rev-parse --short HEAD) .
```

### Push Docker Image

```bash
docker push mrckurz/cicd-workshop-image:latest
docker push mrckurz/cicd-workshop-image:$(git rev-parse --short HEAD)
```

### Run Trivy Scan

```bash
trivy image mrckurz/cicd-workshop-image:latest
```

### Run SonarCloud Analysis

```bash
sonar-scanner -Dsonar.projectKey=your_project_key -Dsonar.organization=your_org -Dsonar.host.url=https://sonarcloud.io -Dsonar.login=$SONAR_TOKEN
```

### Deploy to Minikube

```bash
./redeploy.sh
```

---

## 🧭 Minikube Deployment Guide

### 1. **Start Minikube**

```bash
minikube start
```

### 2. **Apply Deployment Configuration**

```bash
kubectl apply -f minikube-deployment.yaml
```

### 3. **Deploy Application**

```bash
./redeploy.sh
```

### 4. **Access the Application**

```bash
minikube service cicd-workshop
```

---

## 📂 Project Structure

```
cicd-workshop/
├── .github/workflows/ci.yml
├── Dockerfile
├── README.md
├── redeploy.sh
├── minikube-deployment.yaml
├── go.mod
├── go.sum
└── main.go
```

---

## 🏷️ GitHub Badges

![Build](https://img.shields.io/github/actions/workflow/status/mrckurz/cicd-workshop/go.yml?branch=main)
![Docker Pulls](https://img.shields.io/docker/pulls/mrckurz/cicd-workshop-image)
