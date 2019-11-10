function upload_one_file_to_server(file, value_change, percentage_string_change) {
    axios.post("/upload", file, {
        onUploadProgress: function (progressEvent) {
            value_change = progressEvent.loaded;
            console.log(progressEvent)
            // percentage_string_change = progressEvent
            // for (let i in value_list_need_to_change) {
            //     value_list_need_to_change[i] = progressEvent.loaded;
            // }
        }, headers: {
            'Content-Type': 'multipart/form-data'
        }
    }).then(() => {
    }, (err) => {
        alert("upload error.")
    })
}