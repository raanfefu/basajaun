CAKEY=ssl/ca/private/rootCA.key
CA=ssl/ca/certs/rootCA.crt
CACNF=ssl/ca/config/ca.cnf

KEY_K8SAC=ssl/admission-controller/private/admission-controller.authz-system.svc.key
CSR_K8SAC=ssl/admission-controller/certs/admission-controller.authz-system.svc.csr
CRT_K8SAC=ssl/admission-controller/certs/admission-controller.authz-system.svc.crt
CNF_k8SAC=ssl/admission-controller/config/admission-controller.cnf 

genca:
	echo "araguaney" | openssl genrsa -des3 -passout stdin -out ${CAKEY} 4096
	echo "araguaney" | openssl req -x509 -new -nodes  -key ${CAKEY} -sha256 -days 2048 -out ${CA}  -config ${CACNF} -passin stdin
genac:	
	echo "araguaney" | openssl genrsa -out ${KEY_K8SAC} 4096 -passout stdin
	openssl req -new -key ${KEY_K8SAC} -config  ${CNF_k8SAC} -out ${CSR_K8SAC}
	openssl req -noout -text -in ${CSR_K8SAC} | grep -A 1 "Subject Alternative Name"
	echo "araguaney" | openssl x509 -req -in ${CSR_K8SAC}  -CA ${CA} -CAkey ${CAKEY} -CAcreateserial  -out ${CRT_K8SAC} -days 2048 -sha256 -passin stdin -extensions req_ext -extfile ${CNF_k8SAC}
 

 