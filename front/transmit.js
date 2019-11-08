function upload_files_to_server(file) {
    axios.post("/", file, {
        onUploadProgress: function (progressEvent) {
            console.log(progressEvent)
        }
    }).then(() => {
        console.log("good luck")
    },(err) => {
        console.log("error!")
    })
}