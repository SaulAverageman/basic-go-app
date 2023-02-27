pipeline{
    agent {
        label 'docker-build-agent'
    }

    stages{
        stage ("Docker build") {
            steps{
                sh 'docker build -t go-app-delete-asap .'
            }
        }

        stage ("Docker scan"){
            steps{
                sh 'docker scan -f Dockerfile go-app-delete-asap'
            }
        }
    }
}