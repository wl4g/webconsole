const server_admin_url = "http://10.0.0.136:8888/admin";
const server_ws_url = "ws://10.0.0.136:8888/ws";


function openAdd() {
    const node = $('#add');
    if(node.is(':hidden')){　　//如果node是隐藏的则显示node元素，否则隐藏
        node.show();
        $('#sessions').hide();
    }else{
        node.hide();
    }
}

function openSession() {
    const node = $('#sessions');
    if(node.is(':hidden')){　　//如果node是隐藏的则显示node元素，否则隐藏
        node.show();
        $('#add').hide();
        list();
    }else{
        node.hide();
    }
}

function changeFile(){
    debugger
    let files = $('#sshKeyFile').prop('files');
    if(!files || !files[0]){
        return;
    }
    let f = files[0];
    let reader = new FileReader(); //新建一个FileReader
    reader.readAsText(f, "UTF-8"); //读取文件
    reader.onload = function (evt) { //读取完文件之后会回来这里
        let fileString = evt.target.result;
        $('#sshKey').val(fileString);
    }
}

function changeType(type) {
    if(type === 1){
        $('#passwordDiv').show();
        $('#sshKeyDiv').hide();
    }else if(type === 2){
        $('#sshKeyDiv').show();
        $('#passwordDiv').hide();
    }

}

function save() {
    let address=$('#address').val();
    let username=$("#username").val();
    let password=$('#password').val();
    let sshKey=$('#sshKey').val();
    let alias=$('#alias').val();
    $.post(server_admin_url + "/add",{
        address: address,
        username: username,
        password: password,
        name: alias,
        sshKey: sshKey,
    },function(result){
        $('#add').hide();
        if(result && result.id){
            //TODO
            ws_connect(result.id)
        }
    });
}

function del(id) {
    $.post(server_admin_url + "/del",{
        id: id,
    },function(result){
        list();
    });
}

function list() {
    $.get(server_admin_url + "/list", function(result){
        if(!result || !result.sessions){
            return;
        }
        $('#sessionsul').empty();
        for(let i in result.sessions){
            let session = result.sessions[i];
            let sessionHtml = '<li><a style="cursor:pointer" onclick="ws_connect('+session.Id+')">'+session.Name+'('+session.Username+'@'+session.Address+')</a><button onclick="ws_connect('+session.Id+')">Connect</button><button onclick="del('+session.Id+')">del</button></li>';
            $("#sessionsul").append(sessionHtml);
        }
    });
}


//init
$('input:radio[name=type]')[0].checked = true;
changeType(1);

////===========
