
# 1
sudo docker build -t aymanpatel/gin-web:1.0.1 .

# 2
sudo docker push aymanpatel/gin-web:1.0.1 



For k8 to pull from local registry, 
1. Need to add imageSecrets 
2. Add Docker Hub image under spec.container.image in deployment.yml 