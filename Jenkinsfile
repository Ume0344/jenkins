/* Requires the Docker Pipeline plugin */
pipeline {
    agent { docker { image 'golang:1.21.0-alpine3.18' } }
    stages {
        stage('build') {
            steps {
                sh 'go --version'
            }
        }
    }
}
