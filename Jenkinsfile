pipeline {
    agent any

    stages {
        stage('Setup Docker') {
            steps {
                // Set up Docker environment
                sh 'docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine'
                sh 'sleep 10'
            }
        }

        stage('Run db migration') {
            steps {
                script {
                    // Run migration with the updated hostname
                    sh 'migrate -path db/migration -verbose -database "postgres://root:admin@localhost:5432/bank?sslmode=disable" up'
                }
            }
        }
    }

    post {
        always {
            // Stop and remove the container after the pipeline finishes
            sh 'docker stop postgres12'
            sh 'docker rm postgres12'
        }
    }
}
