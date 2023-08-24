/* Requires the Docker Pipeline plugin */
pipeline {
    agent any
    tools {
        go 'go1.20'
    }
    stages {
        stage('Pre-Build') {
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
