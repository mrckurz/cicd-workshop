#!/bin/bash

set -e

DEPLOYMENT_NAME="cicd-workshop"
IMAGE_NAME="mrckurz/cicd-workshop-image:latest"

# Prüfe ob Minikube läuft, sonst starten
if ! minikube status | grep -q "host: Running"; then
  echo "🚀 Starte Minikube..."
  minikube start
else
  echo "✅ Minikube läuft bereits."
fi

# Prüfe, ob Deployment existiert, sonst apply
if ! kubectl get deployment "$DEPLOYMENT_NAME" &> /dev/null; then
  echo "📦 Deployment '$DEPLOYMENT_NAME' existiert nicht – wende minikube-deployment.yaml an..."
  kubectl apply -f minikube-deployment.yaml
fi

# Zeige aktuelles Image im Deployment
CURRENT_IMAGE=$(kubectl get deployment $DEPLOYMENT_NAME -o jsonpath='{.spec.template.spec.containers[0].image}' 2>/dev/null || echo "Nicht vorhanden")
echo "📦 Aktuelles Image im Deployment: $CURRENT_IMAGE"
echo "🎯 Ziel-Image: $IMAGE_NAME"

echo "🔄 Setze Image auf $IMAGE_NAME im Deployment $DEPLOYMENT_NAME..."
kubectl set image deployment/$DEPLOYMENT_NAME $DEPLOYMENT_NAME=$IMAGE_NAME --record

echo "⏳ Warte auf Rollout..."
kubectl rollout status deployment/$DEPLOYMENT_NAME

echo "🚀 Starte lokalen Zugriff auf Minikube Service..."
minikube service $DEPLOYMENT_NAME