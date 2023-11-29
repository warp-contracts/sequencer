variable "TAG" {
  default = "dev"
}

group "default" {
  targets = ["sequencer"]
}

target "sequencer" {
  args = {
    ENV = "devnet"
  }
  dockerfile = "Dockerfile"
  platforms = ["linux/amd64"]
  tags = ["docker.io/warpredstone/sequencer:${TAG}"]
}
