pipeline {
    agent any

    stages {
        stage('Run Postgres ') {
            steps {
                // Check if the container already exists
                script {
                    
                    // Run Postgres Docker image
                    sh 'docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -v /home/vboxuser/postgres-bank-volume/data:/var/lib/postgresql/data -d --rm postgres:12-alpine'
                    // Wait for Postgres to start 
                    sh 'sleep 10'
                    // drop db if exists
                    sh 'docker exec postgres12 dropdb --if-exists bank'
                    // Create a database inside the Postgres container
                    sh 'docker exec postgres12 createdb --username=root --owner=root bank'
                    
                }
            }
        }

        stage('Run db migration') {
            steps {
                script {
                    sh 'migrate -path db/migration -verbose -database "postgres://root:admin@localhost:5432/bank?sslmode=disable" up'
                }
            }
        }
    }

    post {
        always {
            // Stop and remove the container after the pipeline finishes
            sh 'docker exec postgres12 dropdb --if-exists bank'
            sh 'docker stop postgres12'
        }
    }
}
