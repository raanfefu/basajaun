
### Generar Certificados.

En esta sección se generan los certificados auto firmadas que seran utilizados por el operador encargado aprovisionar los Pod necesarios para la autorización.

Procedemos a generar una key para nuestra CA con la firmaremos los certificados necesarios para el operador k8s que deseamos crear. 

##### Crear clave raíz
```sh
$ echo "araguaney" | openssl genrsa -des3 -passout stdin -out rootCA.key 4096
```
##### Crear clave raíz

Procedemos a crear nuestro certificados raiz el cual utilizaremos para firma los certificados que vamos a distribur

```sh
$ echo "araguaney" | openssl req -x509 -new -nodes \
    -key rootCA.key -sha256 -days 2048 -out rootCA.crt \
    -subj "/C=PE/ST=LI/L=LI/O=araguaney/OU=admission_ca/CN=admission_ca/emailAddress=rafael.fernandez.ve@gmailcom" \
    -passin stdin
```


##### Crear clave para admision controller 

```sh
$ echo "araguaney" | openssl genrsa -out admission-controller.authz-opa.svc.key 2048 -passout stdin
```

##### Crear una solicitud de firma para admision controller

```sh
$ openssl req -new -key admission-controller.authz-opa.svc.key \
    -subj "/C=PE/ST=LI/L=LI/O=araguaney/OU=admission_ca/CN=admission-controller.authz-opa.svc/emailAddress=rafael.fernandez.ve@gmailcom" \
    -out admission-controller.authz-opa.svc.csr
```

##### Crear un certificado para nuestro admision controller 

```sh
$ echo "araguaney" | openssl x509 -req -in "admission-controller.authz-opa.svc.csr" \
    -CA rootCA.crt -CAkey rootCA.key -CAcreateserial \
    -out "admission-controller.authz-opa.svc.crt" -days 2048 -sha256 -passin stdin
```