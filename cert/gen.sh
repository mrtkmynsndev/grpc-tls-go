# delete pem file
rm *.pem 

# Create CA private key and self-signed certificate
# adding -nodes to not encrypt the private key
openssl req -x509 -newkey rsa:4096 -nodes -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=TR/ST=ASIA/L=ISTANBUL/O=DEV/OU=TUTORIAL/CN=*.tutorial.dev/emailAddress=mert@tutorial.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text 

# Create Web Server private key and CSR
# adding -nodes to not encrypt the private key
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=TR/ST=ASIA/L=ISTANBUL/O=DEV/OU=BLOG/CN=*.mertkimyonsenblog.com/emailAddress=info@mertkimyonsenblog.com"

# Sign the Web Server Certificate Request CSR
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.conf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# Verify certificate
echo "Verifying certificate"
openssl verify -CAfile ca-cert.pem server-cert.pem


