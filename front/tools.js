function bytes_to_readable_string(size) {
    if (size < 1e4) {
        return size.toString() + "B";
    } else if (size < 1e6) {
        return (size / 1000).toFixed(2) + "KB";
    } else if (size < 1e9) {
        return (size / 1e6).toFixed(2) + "MB";
    } else
        return (size / 1e9).toFixed(3) + "GB";
}

function file_size_check_allow(files, maxSizeEach, maxSizeTotal) {
    let total_size = 0;
    for (let i in files) {
        if (files.hasOwnProperty(i)) {
            if (files[i].size > maxSizeEach) {
                return false
            }
            total_size += files[i].size
        }
    }
    return total_size < maxSizeTotal;
}

function calculate_last_time_seconds(size, spare_space_bytes) {
    let store_for_days = spare_space_bytes / size;
    if (store_for_days > 400) {
        return null;// forever.
    } else {
        return Math.floor(store_for_days * 24 * 3600)
    }
}


function seconds_to_readable(seconds) {
    // TIP: to find current time in milliseconds, use:
    // let  current_time_milliseconds = new Date().getTime();
    if (!seconds) {
        return "永久"
    }

    function numberEnding(number) {
        return (number > 1) ? 's' : '';
    }

    let years = Math.floor(seconds / 31536000);
    if (years) {
        return years + ' year' + numberEnding(years);
    }
    //TODO: Months! Maybe weeks?
    let days = Math.floor((seconds %= 31536000) / 86400);
    if (days) {
        return days + ' day' + numberEnding(days);
    }
    let hours = Math.floor((seconds %= 86400) / 3600);
    if (hours) {
        return hours + ' hour' + numberEnding(hours);
    }
    let minutes = Math.floor((seconds %= 3600) / 60);
    if (minutes) {
        return minutes + ' minute' + numberEnding(minutes);
    }
    return 'less than a minute'; //'just now' //or other string you like;
}

function append_files_to_upload_form(files, serverMaxTotalSize, form) {
    let current_size = 0;
    for (let file of files) {
        let line = "<tr><td>" + file.name + "</td><td>" + bytes_to_readable_string(file.size)
            + "</td>" + "<td>" + seconds_to_readable(calculate_last_time_seconds(file.size, serverMaxTotalSize - current_size)) + "</td>"
            + "<td><div class='progress'><div class='progress-bar' role='progressbar' style='width: 0;'>0%</div></div></td></tr>";
        form.append(line);
        current_size += file.size;
    }
}

function upload_single_file_to_server(file, onUploadProgressCallBack) {
    let formData = new FormData();
    formData.append('file', file);
    return axios.post('/upload', formData, {onUploadProgress: onUploadProgressCallBack})
}

function refresh_upload_files_table(files, server_max_size, files_table) {
    files_table.find("tr").remove();
    append_files_to_upload_form(files, server_max_size, files_table)
}

function refill_server_files_list(server_files_json, server_files_table, deleteCallBack) {
    server_files_table.find("tr").remove();
    for (let single_file_info of server_files_json) {
        let line = "<tr><td>" + single_file_info["FileName"] + "</td><td>" + bytes_to_readable_string(single_file_info["FileSizeBytes"]) + "</td><td>"
            + seconds_to_readable(single_file_info["FileSurplusKeepSeconds"]) + "</td></tr>";
        let download_label = $("<a class='btn btn-outline-primary' style='margin-right: 10px'>下载</a>").prop('href', '/files/' + single_file_info["FileName"]);
        let delete_label = $("<a class='btn btn-outline-danger' style='margin-left: 10px'>删除</a>").prop('href', 'javascript:void(0)');
        delete_label.click(() => {
            axios.delete(URLPrefix + '/files/' + single_file_info["FileName"]).then(() => {
                alert("删除成功");
            }, () => {
                alert("删除失败！");
            }).then(deleteCallBack)
        });
        let prepared_line = $(line).append($("<td style='text-align: center'></td>").append(download_label, delete_label));
        server_files_table.append(prepared_line)
    }
}