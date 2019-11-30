// 我决定拥抱时代，使用JQuery
$(() => {
    $.get('/maxSpace', (responseTxt) => {
        $("#space").text(bytes_to_readable_string(responseTxt))
    });

    $("#up_input").change(() => {
        $.get('/maxSpace', (responseTxt) => {
            let files = $("#up_input").prop('files');
            if (file_size_check_allow(files, 5e7,parseInt(responseTxt))){
                //　文件大小校验通过。

            }
        })
    })
});