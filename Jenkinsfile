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
                    // Create a database inside the Postgres container
                    sh 'docker exec postgres12 createdb --username=root --owner=root bank'
                    
                }

                post {
                    always {
                        // Stop and remove the container after the pipeline finishes
                        sh 'docker stop postgres12'
                        sh 'docker rm postgres12'
                    }
                }
            }
        }
    }
}
