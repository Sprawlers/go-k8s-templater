package main

import (
    appsv1  "k8s.io/api/apps/v1"
    metav1  "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) updatePod(namespace string, i *Image) (*appsv1.Deployment, error) {
    deploymentName := i.Name + "-dev-deployment"
    deploymentsClient := s.clientset.AppsV1().Deployments(namespace)
    deployment, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
    if err != nil {
        return nil, err
    }
    deployment.Spec.Template.Spec.Containers[0].Image = i.Repo + ":" + i.Tag
    result, err := deploymentsClient.Update(deployment)
    if err != nil {
        return nil, err
    }
    return result, nil
}
