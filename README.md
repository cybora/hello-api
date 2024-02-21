## Setup
 
### Install Go          
`sudo make setup`
 
### Upgrade Go          
`sudo make install-go`


gcloud iam service-accounts add-iam-policy-binding galvanic-augury-415006@appspot.gserviceaccount.com --member "serviceAccount:hello-api@galvanic-augury-415006.iam.gserviceaccount.com" --role roles/iam.serviceAccountUser