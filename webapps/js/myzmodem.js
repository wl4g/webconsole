function checkwindow() {
    event.returnValue=false;
}

function utf8_to_b64(rawString) {
    return btoa(unescape(encodeURIComponent(rawString)));
}

function b64_to_utf8(encodeString) {
    return decodeURIComponent(escape(atob(encodeString)));
}

function readFile(element_id, res_id) {
    const objFile = document.getElementById(element_id);
    if(objFile.value === '') {
        return
    }
    // 获取文件
    const files = objFile.files;
    if (files[0].size > 16 * 1024) {
        toastr.warning("文件大于 16 KB，请选择正确的密钥");
        objFile.value = "";
        return
    }
    // 新建一个FileReader
    const reader = new FileReader();
    // 读取完文件之后会回来这里
    reader.onload = function(e) {
        // 读取文件内容
        let fileString = e.target.result;
        // 接下来可对文件内容进行处理
        $("#" + res_id).text(fileString);
    };
    reader.onerror = function(e) {
        console.log(e);
        toastr.warning("读取密钥错误");
        objFile.value = "";
    };
    // 读取文件
    reader.readAsText(files[0], "UTF-8");
}

function get_term_size() {
    let init_width = 9;
    let init_height = 17;
    let windows_width = $(window).width();
    let windows_height = $(window).height();
    return {
        cols: Math.floor(windows_width / init_width),
        rows: Math.floor(windows_height / init_height),
    }
}

function uploadFile(zsession) {
    let uploadHtml = "<div>" +
        "<label class='upload-area' style='width:100%;text-align:center;' for='fupload'>" +
        "<input id='fupload' name='fupload' type='file' style='display:none;' multiple='true'>" +
        "<i class='fa fa-cloud-upload fa-3x'></i>" +
        "<br />" +
        "点击选择文件, 请尽量使用 rz -O 方式上传" +
        "</label>" +
        "<br />" +
        "<span style='margin-left:5px !important;' id='fileList'></span>" +
        "</div><div class='clearfix'></div>";

    let upload_dialog = bootbox.dialog({
        message: uploadHtml,
        title: "上传文件",
        buttons: {
            cancel: {
                label: '关闭',
                className: 'btn-default',
                callback: function (res) {
                    try {
                        term.detach();
                    } catch (e) {
                        // console.log(e);
                    }
                    try {
                        term.attach(socket);
                    } catch (e) {
                        // console.log(e);
                    }
                    try {
                        // zsession 每 5s 发送一个 ZACK 包，5s 后会出现提示最后一个包是 ”ZACK“ 无法正常关闭
                        // 这里直接设置 _last_header_name 为 ZRINIT，就可以强制关闭了
                        zsession._last_header_name = "ZRINIT";
                        zsession.close();
                    } catch (e) {
                        console.log(e);
                    }
                }
            },
        },
        closeButton: false,
    });

    function hideModal() {
        upload_dialog.modal('hide');
    }

    let file_el = document.getElementById("fupload");

    return new Promise((res) => {
        file_el.onchange = function (e) {
            let files_obj = file_el.files;
            hideModal();
            let files = [];
            for (let i of files_obj) {
                if (i.size <= 2048 * 1024 * 1024) {
                    files.push(i);
                } else {
                    toastr.warning(`${i.name} 超过 2048 MB, 无法上传`);
                    // console.log(i.name, i.size, '超过 2048 MB, 无法上传');
                }
            }
            if (files.length === 0) {
                try {
                    term.detach();
                } catch (e) {
                    // console.log(e);
                }
                try {
                    term.attach(socket);
                } catch (e) {
                    // console.log(e);
                }
                try {
                    // zsession 每 5s 发送一个 ZACK 包，5s 后会出现提示最后一个包是 ”ZACK“ 无法正常关闭
                    // 这里直接设置 _last_header_name 为 ZRINIT，就可以强制关闭了
                    zsession._last_header_name = "ZRINIT";
                    zsession.close();
                } catch (e) {
                    console.log(e);
                }
                return
            }
            //Zmodem.Browser.send_files(zsession, files, {
            Zmodem.Browser.send_block_files(zsession, files, {
                    on_offer_response(obj, xfer) {
                        if (xfer) {
                            // term.write("\r\n");
                        } else {
                            term.write(obj.name + " was upload skipped\r\n");
                            // socket.send(JSON.stringify({ type: "ignore", data: utf8_to_b64("\r\n" + obj.name + " was upload skipped\r\n") }));
                        }
                    },
                    on_progress(obj, xfer) {
                        updateProgress(xfer);
                    },
                    on_file_complete(obj) {
                        term.write("\r\n");
                        socket.send(JSON.stringify({ type: "ignore", data: utf8_to_b64("\r\n" + obj.name + " was upload success\r\n") }));
                    },
                }
            ).then(zsession.close.bind(zsession), console.error.bind(console)
            ).then(() => {
                res();
            });
        };
    });
}

function saveFile(xfer, buffer) {
    return Zmodem.Browser.save_to_disk(buffer, xfer.get_details().name);
}

async function updateProgress(xfer) {
    let detail = xfer.get_details();
    let name = detail.name;
    let total = detail.size;
    let offset = xfer.get_offset();
    let percent;
    if (total === 0 || total === offset) {
        percent = 100
    } else {
        percent = Math.round(offset / total * 100);
    }
    term.write("\r" + name + ": " + total + " " + offset + " " + percent + "% ");
}

function downloadFile(zsession) {
    zsession.on("offer", function(xfer) {
        function on_form_submit() {
            if (xfer.get_details().size > 2048 * 1024 * 1024) {
                xfer.skip();
                toastr.warning(`${xfer.get_details().name} 超过 2048 MB, 无法下载`);
                return
            }
            let FILE_BUFFER = [];
            xfer.on("input", (payload) => {
                updateProgress(xfer);
                FILE_BUFFER.push( new Uint8Array(payload) );
            });

            xfer.accept().then(
                () => {
                    saveFile(xfer, FILE_BUFFER);
                    term.write("\r\n");
                    socket.send(JSON.stringify({ type: "ignore", data: utf8_to_b64("\r\n" + xfer.get_details().name + " was download success\r\n") }));
                },
                console.error.bind(console)
            );
        }
        on_form_submit();
    });
    let promise = new Promise( (res) => {
        zsession.on("session_end", () => {
            res();
        });
    });
    zsession.start();
    return promise;
}

$("body").attr("onbeforeunload", 'checkwindow()'); //增加刷新关闭提示属性

let zsentry = new Zmodem.Sentry( {
    to_terminal: function(octets) {},  //i.e. send to the terminal
    on_detect: function(detection) {
        debugger
        let zsession = detection.confirm();
        let promise;
        if (zsession.type === "receive") {
            promise = downloadFile(zsession);
        } else {
            promise = uploadFile(zsession);
        }
        promise.catch( console.error.bind(console) ).then( () => {
            //
        });
    },
    on_retract: function() {},
    sender: function(octets) { socket.send(new Uint8Array(octets)) },
});
