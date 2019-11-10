/* 存储时间计算方法：剩余空间/文件大小 小时，当时间超过400天时，永久存储
* */

function seconds_to_readable(seconds) {
    // TIP: to find current time in milliseconds, use:
    // let  current_time_milliseconds = new Date().getTime();
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

function calculate_last_time_seconds(size, spare_space_bytes) {
    let store_for_days = spare_space_bytes / size;
    if (store_for_days > 400) {
        return null;// forever.
    } else {
        return Math.floor(store_for_days * 24 * 3600)
    }
}

function bytes_to_readable_string(size, spare_space_bytes) {
    if (size >= spare_space_bytes) {
        alert("内部存储空间不足，暂时无法传输！");
        return null
    }
    if (size < 1e4) {
        return size.toString() + "B";
    } else if (size < 1e6) {
        return (size / 1000).toFixed(2) + "KB";
    } else if (size < 1e9) {
        return (size / 1e6).toFixed(2) + "MB";
    } else if (size < 1e10) {
        return (size / 1e9).toFixed(3) + "GB";
    } else {
        alert("上传过大文件（超过10G），无法传输！");
        return null
    }
}

function clear_table(table_item) {
    while (table_item.childElementCount > 1) {
        table_item.removeChild(table_item.children[table_item.childElementCount - 1])
    }
}

function create_table_line(file_metadata, file_size_bytes) {
    let line = document.createElement("tr");
    let child_table_row_element;
    for (let i in file_metadata) {
        if (file_metadata.hasOwnProperty(i)) {
            child_table_row_element = document.createElement("td");
            child_table_row_element.innerText = file_metadata[i];
            line.appendChild(child_table_row_element);
        }
    }
    let progress_obj = document.createElement("progress");
    progress_obj.max = file_size_bytes;
    progress_obj.className = "bar";
    let tr_container_hold_progress_obj = document.createElement("td");
    tr_container_hold_progress_obj.className = "lint";
    tr_container_hold_progress_obj.insertBefore(progress_obj, null);
    let progress_text = document.createTextNode("Waiting...");
    tr_container_hold_progress_obj.insertBefore(progress_text, null);
    line.insertBefore(tr_container_hold_progress_obj, null);
    return line;// 在此处直接产生一个progress对象，便于以后的函数操作。
}

function check_upload_files(event) {
    let server_spare_space = 1e7;// TODO: 询问后端
    let table = document.getElementById("files_table");
    clear_table(table);
    let files = this.files;
    for (let index in files) {
        if (files.hasOwnProperty(index)) {
            let file = files[index];
            let file_bytes_string = bytes_to_readable_string(file.size);
            if (!file_bytes_string) {
                continue;
            }
            let last_time_seconds = calculate_last_time_seconds(file.size, server_spare_space);
            let keep_time_string;
            if (last_time_seconds == null) {// 永久保存
                keep_time_string = "永久";
            } else {
                keep_time_string = seconds_to_readable(last_time_seconds);
            }
            let table_line = create_table_line([file.name, file_bytes_string, keep_time_string], file.size);
            table.insertBefore(table_line, null)
        }
    }
}

function upload_single_file(file, index) {// index意思就是，这是第几个正在上传的文件，以便绑定。
    let bar = document.getElementsByClassName("bar")[index];
    let tex = document.getElementsByClassName("lint")[index];
    bar.value = 0;
    tex.innerText = "0%";
    let formData = new FormData();// 传输文件的话，是按照formData格式传输对象，所以在此处构建FormData。
    formData.append("file", file);
    upload_one_file_to_server(formData, bar.value, tex.innerText);// 让下层引用progress对象的value，在上传过程中动态修改。
}

function submit_all_files(file_list) {
    for (let i in file_list) {
        upload_single_file(file_list[i], i)
    }
}

window.onload = function () {
    let upload_file_list = document.getElementById("up_input");
    let submit_button = document.getElementById("sm_button");
    submit_button.onclick = function () {
        submit_all_files(upload_file_list.files);
        submit_button.enabled = false
    };
    upload_file_list.onchange = check_upload_files;
};
