/* Requires the Docker Pipeline plugin */
pipeline {
    agent none
    stages {
        stage('Pre-Build') {
            agent { 
                docker { 
                    image 'golang:1.21.0-alpine3.18' 
                    } 
                }
            steps {
                echo "Installing dependencies to run go code"
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }

        }
        stage('build') {
            steps {
                echo "Compiling and building"
                sh 'go build'
            } 
        }
    }
}
