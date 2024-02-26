output "dev_ip" {
  value = aws_instance.dev_node_a1_docker_golangsvelte_development.public_ip
}
output "owner_of_instance" {
  value = var.developer_name
}