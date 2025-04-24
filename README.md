
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

![Build](https://img.shields.io/github/actions/workflow/status/mrckurz/cicd-workshop/ci.yml?branch=main)
![Docker Pulls](https://img.shields.io/docker/pulls/mrckurz/cicd-workshop-image)
![License](https://img.shields.io/github/license/mrckurz/cicd-workshop)
