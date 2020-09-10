terraform {
  required_providers {
    docker = {
      source = "terraform-providers/docker"
    }
  }
}

provider "docker" {}



resource "docker_image" "mongo" {
  name = "mongo:latest"
}

resource "docker_container" "mongo" {
  image = docker_image.mongo.latest
  name = "mongo"

  ports {
    internal = 27017
    external = 27017
  }
}

resource "docker_image" "rabbitmq" {
  name = "rabbitmq:latest"
}

resource "docker_container" "rabbitmq" {
  image = docker_image.rabbitmq.latest
  name = "rabbitmq"

  ports {
    internal = 5672
    external = 5672
  }
  ports {
    internal = 15672
    external = 15672
  }

  provisioner "local-exec" {
    command = "docker exec rabbitmq rabbitmq-plugins enable rabbitmq_management"
  }  
}
