#!/bin/bash
echo -e "\n\033[0;36m=== Welcome to Echo API local deployment Manager ===\033[0m\n"
echo -e "Please select the option you want to perform:
1. Build Echo API image.
2. Deploy Echo API locally.
3. Remove Echo API instance.\n"

read -p 'Option: ' uservar


if [ "$uservar" = "1" ]; then
    echo -e "\033[0;32mAttempting to build Echo API image...\033[0m"
    (
        export DOCKER_DEFAULT_PLATFORM=linux/amd64 && docker build -f Dockerfile -t echo-api .
    ) || (
        echo -e "\033[0;31mFailed to build Echo API image!\033[0m\n"
    )
elif [ "$uservar" = "2" ]; then
    echo -e "Deploying Echo API instance..."
    kubectl apply -f resources/local-deployment/kubernetes/api-cm.yaml &&
    kubectl apply -f resources/local-deployment/kubernetes/api.yaml &&
    echo -e "\n\033[0;32mEcho API instance deployed. \033[0m\nPlease wait a few seconds for the service to be ready!"
elif [ "$uservar" = "3" ]; then
    echo -e "Removing Echo API instance..."
    (
        kubectl delete -f resources/local-deployment/kubernetes/api.yaml &&
        kubectl delete -f resources/local-deployment/kubernetes/api-cm.yaml
    ) || (
        echo -e "\033[0;31mFailed to remove Echo API instance!\033[0m\n"
    )
    echo -e "\n\033[0;32mEcho API instance removed. \033[0m\n"
else
    echo -e "\033[0;31mInvalid option!\033[0m\n"
fi