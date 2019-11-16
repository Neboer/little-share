function update_progressBar_and_text(index, value, max_value) {
    let bar = document.getElementsByClassName("bar")[index];
    let tex = document.getElementsByClassName("lint")[index];
    bar.max = max_value;
    bar.value = value;
    let percent = value * 100 / bar.max;
    tex.innerText = percent.toFixed(1) + "%"
}

function get_max_spare_space() {
    return axios.get("/maxSpace").then((res) => {
        return parseInt(res.data)
    })
}

function upload_one_file_to_server(file, index) {
    return axios.post("/upload", file, {
        onUploadProgress: function (progressEvent) {
            update_progressBar_and_text(index, progressEvent.loaded, progressEvent.total);
        }
    }).then(() => {
    }, (err) => {
        alert("upload error.");
    }).catch((err) => {
        console.debug(err);
    })
}

function get_files_list() {
    return axios.get("/files").then((list) => {
        return list.data
    })
}

