# CI/CD Workshop
This simple go program is used to demonstrate how easy a CI/CD pipeline can be setup with GitHub Actions. The subsequent steps are:

1. add the go-code
2. add the GitHub Action Configuration File (/.github/workflows/go.yml). There, the tests shall be executed upon every commit to the main branch
3. Configure Docker and add a Dockerfile
4. create the docker image and push it to Dockerhub
5. add a static SourceCode quality scan (SonarCloud) for monitoring the code quality
6. add a Vulnerability Scanner (Trivy)

## Some hints for basic functionalty:
### Go
* go run main.go
* go test -v

### Docker
* docker image build -f Dockerfile -t mrckurz/cicd-workshop-test-image:latest ./ 
* docker images
* docker container run -p 9090:8888 mrckurz/cicd-workshop-test-image:latest
* docker ps
* docker stop [containerId]
* docker image pull mrckurz/cicd-workshop-test-image:[latest|sha]

### Trivy
* trivy image [YOUR_IMAGE_NAME]
* trivy config .

### Git
* git config --global alias.s status
* git config --global alias.glog "log --pretty=oneline --graph -5 --decorate --abbrev-commit"
* git config --global alias.gl 'log --oneline --graph --all'


