# Filename: main.tf

# Configure GCP project
provider "google" {
  project = "<your-project-id>"
}

resource "google_cloud_run_service" "run_service" {
  name     = "<your-service-name>"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "gcr.io/<your-project-id>/<your-image-name>:<your-tag>"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

}

# Allow unauthenticated users to invoke the service
resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

# Display the service URL
output "service_url" {
  value = google_cloud_run_service.run_service.status[0].url
}