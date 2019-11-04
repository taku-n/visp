package main

import "fmt"
import "net/http"
import "os/exec"

type DispViewHandler struct{}

func (h *DispViewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, CONTENTS)
}

func main() {
	var handler DispViewHandler = DispViewHandler{}

	go http.ListenAndServe("127.0.0.1:8080", &handler)

	exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
			"--new-window",
			"http://127.0.0.1:8080/").Run()

	exec.Command("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
			"--new-window",
			"http://127.0.0.1:8080/").Run()
}

const CONTENTS = `<head>
<title>visp</title>
</head>

<body>
<h1>visp</h1>
<video autoplay playsinline width="960" height="540"></video>

<script>
const mediaStreamConstraints = {
	video: true
};

const localVideo = document.querySelector("video");

function gotLocalMediaStream(mediaStream) {
	const localStream = mediaStream;
	localVideo.srcObject = mediaStream;
}

function handleLocalMediaStreamError(error) {
	console.log("navigator.getUserMedia error: ", error);
}

navigator.mediaDevices
		.getDisplayMedia(mediaStreamConstraints)
		.then(gotLocalMediaStream)
		.catch(handleLocalMediaStreamError);
</script>
</body>
`
