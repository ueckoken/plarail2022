group "default" {
  targets = [
    "auto-operation",
    "external",
    "internal",
    "json2grpc",
    "multicaster",
    "positioning",
    "webrtc-sender",
    "camera_sender",
    "frontend",
    "receiver-test",
    "logviewer",
  ]
}

variable "PREFIX" {
  default = "plarail2022"
}

variable "TAG" {
  default = "latest"
}

function "GET_TAG" {
  params = [image]
  result = "ghcr.io/ueckoken/${PREFIX}-${image}:${TAG}"
}

target "auto-operation" {
  context = "./backend/auto_operation"
  tags = [
    GET_TAG("auto-operation")
  ]
}

target "external" {
  context = "./backend/external"
  tags = [
    GET_TAG("external")
  ]
}

target "internal" {
  context = "./backend/internal"
  tags = [
    GET_TAG("internal")
  ]
}
target "json2grpc" {
  context = "./backend/json2grpc"
  tags = [
    GET_TAG("json2grpc")
  ]
}

target "multicaster" {
  context = "./backend/multicaster"
  tags = [
    GET_TAG("frontend-python")
  ]
}

target "positioning" {
  context = "./backend/positioning"
  tags = [
    GET_TAG("positioning")
  ]
}

target "webrtc-sender" {
  context = "./frontend/momo_sender"
  tags = [
    GET_TAG("webrtc-sender")
  ]
}

target "camera_sender" {
  context = "./frontend/camera_sender/"
  tags = [
    GET_TAG("camera_sender")
  ]
}

target "receiver-test" {
  context = "./frontend/skyway_receiver"
  tags = [
    GET_TAG("receiver-test")
  ]
}

target "frontend" {
  context = "./frontend/site"
  tags = [
    GET_TAG("frontend")
  ]
}


target "logviewer" {
  context = "./infra/logviewer"
  tags = [
    GET_TAG("logviewer")
  ]
}
