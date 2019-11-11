function update_progressBar_and_text(index, value) {
    let bar = document.getElementsByClassName("bar")[index];
    let tex = document.getElementsByClassName("lint")[index];
    bar.value = value;
    let percent = value * 100 / bar.max;
    tex.innerText = percent.toFixed(1) + "%"
}

function upload_one_file_to_server(file, index) {
    axios.post("/upload", file, {
        onUploadProgress: function (progressEvent) {
            update_progressBar_and_text(index, progressEvent.loaded);
            console.log(progressEvent)
        }, headers: {
            'Content-Type': 'multipart/form-data'
        }
    }).then(() => {
    }, (err) => {
        alert("upload error.")
    })
}