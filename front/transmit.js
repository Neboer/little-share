function upload_files_to_server(file) {
    axios.post("/upload", file, {
        onUploadProgress: function (progressEvent) {
            console.log(progressEvent)
        }, headers: {
            'Content-Type': 'multipart/form-data'
        }
    }).then(() => {
        console.log("good luck")
    }, (err) => {
        console.log("error!")
    })
}