# Admission Controller
In Kubernetes, an admission controller is a crucial component of the admission control mechanism, responsible for intercepting and validating API requests before they are persisted and acted upon by the Kubernetes API server. These controllers are webhook plugins that can be configured to enforce custom rules, policies, and constraints on the cluster's resources. Admission controllers play a vital role in enhancing the security, compliance, and governance of Kubernetes clusters. They help prevent unauthorized or invalid operations from being executed, ensuring that only valid and permitted requests are allowed to modify the cluster's state. By utilizing admission controllers, administrators can enforce policies tailored to their specific requirements, promoting a more robust and controlled Kubernetes environment.

# Admission Controllers Types
There are two main types of webhooks: validators and mutators.

`ValidatingAdmissionWebhooks`: Validator webhooks are admission controllers designed to validate and potentially deny or modify incoming API requests. When a request is sent to the API server, it is intercepted by the validator webhook, which performs checks against predefined policies or rules. These checks could include verifying resource configurations, checking permissions, or ensuring compliance with security standards. If the request fails to meet the specified criteria, the validator webhook can reject the request, preventing any invalid or unauthorized modifications to the cluster.

`MutatingAdmissionWebhook`: Mutator webhooks, on the other hand, work slightly differently. Instead of rejecting requests outright, they intercept the API requests and have the ability to modify them before they are persisted to the Kubernetes datastore. This means that mutators can dynamically alter the content of the request, adding or modifying fields as required. Mutator webhooks are often used to enforce default values or apply standard configurations to resources, ensuring consistency and simplifying resource creation for users.

# Prerequisites to develop our own admission controller
Before developing our own admission controller for Kubernetes, there are several prerequisites to consider. The first step involves writing and deploying the webhook server, ensuring it is accessible from the Kubernetes cluster. Once the webhook server is set up, the next task is to configure the webhooks on Kubernetes, enabling communication with the webhook server. When an API request is made to the Kubernetes API server, the webhook receives a POST request with a JSON payload containing an AdmissionReview API object. This AdmissionReview object holds all the relevant details about the incoming request. The crucial role of the webhook is to process this information and generate an AdmissionReview response in JSON format. This response will contain the decision on whether the request should be allowed or denied based on the established policies and rules. By following these steps and implementing the necessary logic in the webhook server, we can develop a custom admission controller that enforces specific validation and mutation rules to enhance the security and governance of our Kubernetes cluster.

We also need to validate `admissionregistration.k8s.io/v1` or `admissionregistration.k8s.io/v1beta1` is enabled.\
We can do so using the following command: `kubectl api-versions | grep -i admissionregistration.k8s.io`

# Usage
This repo contains a template project for an admission controller, you can expand it for your own needs.

You can run it with the following command:
```bash
chmod +x deployment/deploy.sh
./deployment/deploy.sh
```
Please note that I am using minikube inside the `deploy.sh` file, you can change it to use `kubectl`.\
Also, don't use the `deploy.sh` in production.
