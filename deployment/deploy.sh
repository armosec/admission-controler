#!/usr/bin/env bash

# Generate certificate
echo "Generating certificates"
key_dir="certs"
mkdir "$key_dir"
chmod 0700 "$key_dir"
cat >server.conf <<EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
prompt = no
[req_distinguished_name]
CN = admission-server-armo.default.svc
[ v3_req ]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = admission-server-armo.default.svc
EOF

# Generate the CA cert and private key
openssl req -nodes -new -x509 -keyout certs/ca.key -out certs/ca.crt -subj "/CN=Admission Controller Armo"

# Generate the private key for the webhook server
openssl genrsa -out certs/admission-tls.key 2048

# Generate a Certificate Signing Request (CSR) for the private key, and sign it with the private key of the CA.
openssl req -new -key certs/admission-tls.key -subj "/CN=admission-server-armo.default.svc" -config server.conf | openssl x509 -req -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/admission-tls.crt -extensions v3_req -extfile server.conf

echo "Creating k8s Secret"
minikube kubectl -- create secret tls admission-tls --cert "certs/admission-tls.crt" --key "certs/admission-tls.key"

echo "Creating k8s admission deployment"
minikube kubectl -- create -f deployment.yaml

echo "Creating k8s webhooks for demo"
CA_BUNDLE=$(cat certs/ca.crt | base64 | tr -d '\n')
sed -e 's@${CA_BUNDLE}@'"$CA_BUNDLE"'@g' <"webhooks.yaml" | minikube kubectl -- create -f -