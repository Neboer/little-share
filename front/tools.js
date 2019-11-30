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
    for(let i in files){
        if (files.hasOwnProperty(i)){
            if (files[i].size > maxSizeEach){
                return false
            }
            total_size += files[i].size
        }
    }
    return total_size < maxSizeTotal;
}

function generate_file_form(files, serverMaxTotalSize) {
    let current_size = 0;
    for (let file of files){
        let line = "<tr><td>" + file.name + "</td></tr>" + bytes_to_readable_string(file.size) +

    }
}