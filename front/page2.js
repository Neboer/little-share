// 我决定拥抱时代，使用JQuery
function update_server_info() {
    $.get('/maxSpace', (responseTxt) => {
        $("#space").text(bytes_to_readable_string(responseTxt))
    });

    $.getJSON('/files', (data) => {
        refill_server_files_list(data, $("#files_list_table"), () => {
            update_server_info();
        })
    });
}

$(() => {
    update_server_info();
    $("#up_input").change(() => {
        $.get('/maxSpace', (responseTxt) => {
            let server_max_size = parseInt(responseTxt);
            let files = $("#up_input").prop('files');
            if (file_size_check_allow(files, 5e7, server_max_size)) {
                //　文件大小校验通过。
                refresh_upload_files_table(files, server_max_size, $("#files_table"))
            } else {
                alert("文件过大，无法传输！")
            }
        })
    });

    $("#sm_button").click(() => {
        let files = $('#up_input').prop('files');
        for (let index in files) {
            if (files.hasOwnProperty(index)) {
                upload_single_file_to_server(files[index], (progressEvent) => {
                    //　箭头函数!!!!箭头函数中的this不能够提供我们想要的功能!
                    let percentage = (progressEvent.loaded / progressEvent.total * 100).toFixed(1) + '%';
			// 找到需要更新的进度条。使用“同行的filename”来辨别。
                    let pbar_need_update = $("div .progress-bar").filter(function () {
                        return $(this).parents("tr").children(":first").text() === files[index].name
                    });
                    pbar_need_update.css("width", percentage);
                    pbar_need_update.css("color", "white");
                    pbar_need_update.text(percentage);
                }).then(update_server_info, () => {
                    alert("上传异常！");
                })
            }
        }
    })
});
