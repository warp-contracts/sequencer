variable "ENV" {
  default = "devnet"
}

variable "VERSION" {
  default = "devnet"
}

variable "FROM_VERSION" {
  default = "devnet"
}

target "default" {
  args = {
    ENV="${ENV}"
    VERSION="${VERSION}"
    FROM_VERSION="${FROM_VERSION}"
  }
  dockerfile = "Dockerfile"
  platforms = ["linux/amd64"]
  tags = ["docker.io/warpredstone/sequencer:${VERSION}-${ENV}"]
}
