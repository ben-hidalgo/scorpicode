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
  name  = "mongo"

  ports {
    internal = 27017
    external = 27017
  }
}


resource "docker_image" "rabbitmq" {
  name = "rabbitmq:3.8.8-management"
}

resource "docker_container" "rabbitmq" {
  image = docker_image.rabbitmq.latest
  name  = "rabbitmq"

  ports {
    internal = 5672
    external = 5672
  }
  ports {
    internal = 15672
    external = 15672
  }

  provisioner "local-exec" {
    command = <<EOT
      docker exec rabbitmq apt-get update
      docker exec rabbitmq apt-get install -q -y netcat
      docker exec rabbitmq sh -c "while ! /bin/nc -z localhost 5672; do sleep 1; done"
      docker exec rabbitmq rabbitmqctl add_user scuser scpass
      docker exec rabbitmq rabbitmqctl set_user_tags scuser administrator
      docker exec rabbitmq rabbitmqctl set_user_tags scuser management
      docker exec rabbitmq rabbitmqctl add_vhost sc
      docker exec rabbitmq rabbitmqctl set_permissions -p sc scuser ".*" ".*" ".*"
EOT
  }
}
